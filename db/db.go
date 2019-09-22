package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goBoilerPlate/config"
	"time"
)
var client *mongo.Client
var database string
func Init() {
	database = config.GetConfig().Database.Database
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().Database.Uri))
	if err != nil {
		fmt.Println("Error Connectino to database", err)
	}
}

func getCollction(collectionName string) *mongo.Collection{
	return client.Database(database).Collection(collectionName)
}