package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	uri            string
	connectTimeout int
	client         *mongo.Client
}

func NewMongoClient(uri string) *MongoClient {
	return &MongoClient{
		uri: uri,
	}
}

func (c *MongoClient) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.connectTimeout)*time.Second)
	defer cancel()
	nativeClient, err := mongo.Connect(ctx, options.Client().ApplyURI(c.uri))
	if err != nil {
		return err
	}
	c.client = nativeClient
	return nil
}

func (c *MongoClient) Disconnect() {
	if c.client != nil {
		c.client.Disconnect(context.Background())
		c.client = nil
	}
}

func (c *MongoClient) GetNativeClient() *mongo.Client {
	return c.client
}
