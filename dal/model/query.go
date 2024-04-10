package model

type Query struct {
	Addresses []Address `json:"addresses" bson:"addresses"`
	ClientIp  string    `json:"client_ip" bson:"client_ip"`
	CreatedAt int64     `json:"created_at" bson:"created_at"`
	Domain    string    `json:"domain" bson:"domain"`
}

type Address struct {
	Ip string `json:"ip" bson:"ip"`
}
