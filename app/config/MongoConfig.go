package config

import (
	"context"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var DB *mongo.Client

var Database *mongo.Database

func InitDb() {
	uri, _ := revel.Config.String("mongo.uri")
	revel.AppLog.Info("MONGO URI: ", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		revel.AppLog.Error("Error initializing mongodb", err)
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	client.Connect(ctx)

	DB = client
	result, _ := revel.Config.String("mongo.database.name")
	Database = DB.Database(result)
}