package mongo101

import (
	"context"
	"fmt"
	"grpc_mongo/connection"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type server struct{}

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

type Person struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func Basic() {
	// init
	uri := "mongodb://mongo:mongo@localhost:27017"
	db_name := "grpc"
	ctx := context.TODO()
	connection.MongoConnection(uri, db_name, 5*time.Second)

	// data
	t1 := Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      "wew",
		Completed: true,
	}

	p1 := Person{
		Name:  "Syamsul",
		Email: "msyamsula1995@gmail.com",
	}

	p2 := Person{
		Name:  "Fajar",
		Email: "fajarip@gmail.com",
	}
	// get collection
	personCollection := connection.MongoDB.Collection("person")

	// insert one
	_, insertErr := personCollection.InsertOne(ctx, t1)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	// insert many
	_, insertManyErr := personCollection.InsertMany(ctx, []interface{}{
		p1,
		p2,
	})

	if insertManyErr != nil {
		log.Fatal(insertManyErr)
	}

	// read all from mongo
	cursor, findErr := personCollection.Find(ctx, bson.M{})
	if findErr != nil {
		log.Fatal(findErr)
	}
	var person []bson.M
	if findErr = cursor.All(ctx, &person); findErr != nil {
		log.Fatal(findErr)
	}

	// read one by one
	cursor, findErr = personCollection.Find(ctx, bson.M{})
	if findErr != nil {
		log.Fatal(findErr)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person bson.M
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", person)
	}

	// read with query/filter
	cursor, findErr = personCollection.Find(ctx, bson.M{"completed": true})
	if findErr != nil {
		log.Fatal(findErr)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person bson.M
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", person)
	}

	// update many with filter
	updateFilter := bson.M{"name": "Fajar"}
	updateValue := bson.M{"$set": bson.M{"email": "f@gmail.com"}}
	_, updateErr := personCollection.UpdateMany(
		ctx,
		updateFilter,
		updateValue,
	)
	if updateErr != nil {
		log.Fatal(updateErr)
	}

	// re-read update
	cursor, findErr = personCollection.Find(ctx, bson.M{"name": "Fajar"})
	if findErr != nil {
		log.Fatal(findErr)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person bson.M
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", person)
	}

	// delete
	// deleteFilter := bson.M{"name": "Syamsul"}
	deleteFilter := bson.M{} // no filter, delete all
	deleteResult, deleteErr := personCollection.DeleteMany(ctx, deleteFilter)
	if deleteErr != nil {
		log.Fatal(deleteErr)
	}

	fmt.Printf("%v", deleteResult.DeletedCount)
}
