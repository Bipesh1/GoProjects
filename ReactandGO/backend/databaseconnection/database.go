package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	Mongouri:= os.Getenv("MONGO_URI")
	if err!=nil{
		log.Fatal("Cannot load .env file",err)
	}
	client,err:= mongo.Connect(context.Background(), options.Client().ApplyURI(Mongouri))
	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client= DBinstance()
