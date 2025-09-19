package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/AishwaryGathe/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

//connect with database mongodb

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected Sucessfully!!!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready!!")

}

// mongodb helpers - files

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 Movie in db with id:", inserted.InsertedID)
}


// update 1 record

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched":true}}

 result, err :=	collection.UpdateOne(context.Background(), filter, update)
 if err != nil {
	log.Fatal(err)
 }

 fmt.Println("Modified count:", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	deletecount, err := collection.DeleteOne(context.Background(),filter)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Moive got delete with delete count:", deletecount)

// delete all records from mongodb
}

func deleteAllMoive() int64 {
	
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}},nil)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Number of movie delete:",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}
