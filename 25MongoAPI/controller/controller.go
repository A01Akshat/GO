package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/A01Akshat/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://akshat29december:Akshat@cluster0.otdx8vw.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbname = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection // very important

// connect mongo

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString) 

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECTION SUCCESS")

	collection = client.Database(dbname).Collection(collectionName)
	fmt.Println("COLLECTION INSTANCE READY")

}

// MONGODB Helpers

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie) // context.Background() is used to create a context for the operation, which is required by the MongoDB driver.bf 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INSERTED 1 MOVIE with id", inserted.InsertedID)

}

// update 1 record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // converts to id which mongo understands


    // bson.M is a type that represents a BSON document in MongoDB, which is similar to a JSON object. It is used to create filters and updates for MongoDB operations.
	// bson.D is another type that represents a BSON document, but it maintains the order of the fields. It is often used when the order of fields matters, such as in aggregation pipelines.


	filter := bson.M{"_id": id}                        // This line creates a filter used to locate the specific document in the MongoDB collection.The filter here is matching the document where _id equals the given id.
	update := bson.M{"$set": bson.M{"watched": true}} // $set is a MongoDB update operator that sets the value of a field. update tells mongodb how to update the values
	// 	update := bson.M{                            // FOR MULTIPLE THINGS UPDATE
	//     "$set": bson.M{
	//         "watched":   true,
	//         "movieName": "The Dark Knight",
	//     },
	// }
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UPDATED COUNT:", res.ModifiedCount)
}

// delete 1 record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // converts to id which mongo understands

	filter := bson.M{"_id": id} // This line creates a filter used to locate the specific document in the MongoDB collection.The filter here is matching the document where _id equals the given id.

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MOVIE GOT DELETED WITH COUNT:", res)
}

func deleteAllMovies() int64 {
	filter := bson.D{{}} 
	res, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NO. OF MOVIES DELETED:", res.DeletedCount)
	return res.DeletedCount // WILL TELL HOW MANY MOVIES GOT DELETED OVERALL
}

// get all movies

func getAllMovies() []primitive.M { // return type is a slice of primitive.M, which is a map type used to represent BSON documents in MongoDB.
	cur, err := collection.Find(context.Background(), bson.D{{}}) // cur=cursor
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M 

	for cur.Next(context.Background()) { // why context.Background() ? because we need to provide a context for the operation, which is required by the MongoDB driver.
		var movie bson.M
		err := cur.Decode(&movie) 
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// Actual Controller

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode") 

	allmovies := getAllMovies()

	json.NewEncoder(w).Encode(allmovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
