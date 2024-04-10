package handler

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/fishstake/dal/db"
	"github.com/ngqinzhe/fishstake/dal/model"
)

type LookupHandler struct {
	DBClient db.MongoDBClient
}

func NewLookupHandler(db db.MongoDBClient) *LookupHandler {
	return &LookupHandler{
		DBClient: db,
	}
}

const (
	queryNameDomain = "domain"
)

func (l *LookupHandler) Handle(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryMap := c.Request.URL.Query()
		domain, ok := queryMap[queryNameDomain]
		if !ok {
			c.JSON(http.StatusBadRequest, model.HttpError{
				Message: "no domain query provided",
			})
			return
		}
		if len(domain) != 1 {
			c.JSON(http.StatusBadRequest, model.HttpError{
				Message: "expected 1 domain in request",
			})
			return
		}
		ips, err := findIPv4Addresses(domain[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, model.HttpError{
				Message: "invalid domain",
			})
			return
		}

		query := &model.Query{
			ClientIp:  c.ClientIP(),
			CreatedAt: time.Now().Unix(),
			Domain:    domain[0],
		}

		for _, ip := range ips {
			query.Addresses = append(query.Addresses, model.Address{Ip: ip})
		}

		// write to db
		if err := l.DBClient.WriteQuery(ctx, query); err != nil {
			// TODO: log
		}
		c.JSON(http.StatusOK, query)
	}
}

func findIPv4Addresses(domain string) ([]string, error) {
	var res []string
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			res = append(res, ipv4.String())
		}
	}
	return res, nil
}
