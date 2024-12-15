package models

import (
	"context"
	"fmt"
	"todo/databaseconnection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `json:"title"`
	Completed bool               `json:"completed"`
	TodoId    string             `json:"todoid"`
}

var db *mongo.Database

func init() {
	fmt.Println("init function")
	var client *mongo.Client = database.Client
	db = client.Database("todo")
}

func Gettodos() ([]Todo, error) {
	fmt.Println("hello")
	t := Todo{}
	todos := []Todo{}
	cursor, err := db.Collection("todo").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&t); err != nil {
			return nil, err
		}
		t.TodoId = t.ID.Hex()
		todos = append(todos, t)
	}
	return todos, nil
}
func Createtodos(t Todo) error {
	t.ID= primitive.NewObjectID()
	_, err := db.Collection("todo").InsertOne(context.Background(), t)
	if err != nil {
		return err
	}
	return nil
}
