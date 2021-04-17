package connection

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database
var MongoClient *mongo.Client

func MongoConnection(uri string, db string, timeout time.Duration) {

	// connect to mongo with option
	clientOption := options.Client().ApplyURI(uri)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	MongoDB = client.Database(db)
	MongoClient = client

}
