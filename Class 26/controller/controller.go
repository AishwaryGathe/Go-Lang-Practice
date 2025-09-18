package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/AishwaryGathe/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:admin123@cluster0.o2i0gao.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

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

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(),movie)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 Movie in db with id:", inserted.InsertedID)
}
