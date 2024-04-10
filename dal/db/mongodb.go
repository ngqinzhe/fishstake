package db

import (
	"context"
	"fmt"

	"github.com/ngqinzhe/fishstake/dal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient interface {
	WriteQuery(ctx context.Context, query *model.Query) error
	FindQueries(ctx context.Context, options ...*options.FindOptions) ([]model.Query, error)
	Close(ctx context.Context)
}

type mongoDB struct {
	client *mongo.Client
}

const (
	driverStr        = "mongodb+srv://ngqinzhe2:qBL9Q0JbqYyu9nhA@cluster0.wrjlydx.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	dbName           = "ip_queries"
	dbCollectionName = "queries"
)

func Init() MongoDBClient {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(driverStr))
	if err != nil {
		panic(err)
	}
	return &mongoDB{
		client: client,
	}
}

func (c *mongoDB) FindQueries(ctx context.Context, options ...*options.FindOptions) ([]model.Query, error) {
	collection := c.client.Database(dbName).Collection(dbCollectionName)
	// find without filters
	cursor, err := collection.Find(ctx, bson.D{{}}, options...)
	if err != nil {
		// TOOD: log
		fmt.Println("find error", err)
		return nil, err
	}
	queries := []model.Query{}

	for cursor.Next(ctx) {
		var query model.Query
		if err := cursor.Decode(&query); err != nil {
			fmt.Println("decode error", err)
			return nil, err
		}
		queries = append(queries, query)
	}

	return queries, nil
}

func (c *mongoDB) WriteQuery(ctx context.Context, query *model.Query) error {
	collection := c.client.Database(dbName).Collection(dbCollectionName)
	_, err := collection.InsertOne(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (c *mongoDB) Close(ctx context.Context) {
	if err := c.client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
