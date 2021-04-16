package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func mongo_connection(timeout time.Duration) *mongo.Client {
	uri := "mongodb+srv://mongo:mongo@localhost:27017/grpc"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("successfully connect and ping")

	return client

}