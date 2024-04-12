package book

import (
	"context"
	"fmt"
	"go-web/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName         = "go-web"
	CollectionName = "books"
)

type BookRepositry struct {
	client *db.MongoClient
}

func NewBookRepository(uri string) *BookRepositry {
	client := db.NewMongoClient(uri)
	return &BookRepositry{
		client: client,
	}
}

func (r *BookRepositry) Add(book Book) (string, error) {
	r.client.Connect()
	defer r.client.Disconnect()
	nativeClient := r.client.GetNativeClient()

	result, err := getCollection(nativeClient).InsertOne(context.TODO(), book)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%v", result.InsertedID)
	return id, nil
}

func (r *BookRepositry) GetAll() ([]Book, error) {
	r.client.Connect()
	defer r.client.Disconnect()
	nativeClient := r.client.GetNativeClient()

	cursor, err := getCollection(nativeClient).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var books []Book
	if err := cursor.All(context.TODO(), &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepositry) Get(isbn string) (*Book, error) {
	r.client.Connect()
	defer r.client.Disconnect()
	nativeClient := r.client.GetNativeClient()

	filter := bson.D{{"isbn", isbn}}
	result := getCollection(nativeClient).FindOne(context.TODO(), filter)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var book Book
	if err := result.Decode(&book); err != nil {
		return nil, err
	}

	return &book, nil
}

func getCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(DbName).Collection(CollectionName)
}
