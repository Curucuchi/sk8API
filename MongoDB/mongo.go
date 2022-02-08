package dbb

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Trick struct {
	TrickID   string `bson:"_id"`
	TrickName string `bson:"trick_name"`
}

var Tricks []Trick

const DSN = "mongodb://admin:SUPERSECRETPASSWORD@10.0.0.164:27017"

func Connect() {
	ctx := context.Background()

	client, err := mongo.NewClient(options.Client().ApplyURI(DSN))
	if err != nil {
		log.Fatal("There was an issue creating the client: ", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("There was an issue connecting to mongoDB cluster:", err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("There was an issue pinging the DB: ", err)
	}

}

func GetTricks() ([]byte, error) {
	ctx := context.Background()

	client, err := mongo.NewClient(options.Client().ApplyURI(DSN))
	if err != nil {
		log.Fatal("There was an issue creating the client: ", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("There was an issue connecting to mongoDB cluster:", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("There was an issue pinging the DB: ", err)
	}

	defer client.Disconnect(ctx)

	sk8DB := client.Database("sk8opia")
	tricksCollection := sk8DB.Collection("tricks")

	cursor, err := tricksCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("There was an issue querying the tricks collection: ", err)
	}

	defer cursor.Close(ctx)

	//var tricks []bson.M
	if err = cursor.All(ctx, &Tricks); err != nil {
		log.Fatal("There was an issue scanning tricks: ", err)
	}

	fmt.Println(Tricks)

	return json.MarshalIndent(Tricks, "", "  ")
}

func CreateTricks(userInput string) {
	ctx := context.Background()
	userTrick := Trick{TrickID: "",
		TrickName: userInput}

	client, err := mongo.NewClient(options.Client().ApplyURI(DSN))
	if err != nil {
		log.Fatal("There was an issue creating the client: ", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("There was an issue connecting to mongoDB cluster:", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("There was an issue pinging the DB: ", err)
	}

	defer client.Disconnect(ctx)

	sk8DB := client.Database("sk8opia")
	tricksCollection := sk8DB.Collection("tricks")

	cursor, err := tricksCollection.InsertOne(ctx, userTrick)
	if err != nil {
		log.Fatal("There was an issue inserting Trick:", err)
	}
	fmt.Println(cursor)
}
