package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Person struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"Caption,omitempty" bson:"Caption,omitempty"`
	ImageURL  string             `json:"ImageURL,omitempty" bson:"ImageURL,omitempty"`
	Timestamp string             `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	UserID    string             `json:"userid,omitempty" bson:"userid,omitempty"`
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(request.URL.Path, "/users/"))
	var person Person
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection := client.Database("Instagram").Collection("Users")
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	checkid := strings.TrimPrefix(request.URL.Path, "/users")
	if checkid != "/" {
		GetPersonEndpoint(response, request)
	} else {
		response.Header().Set("content-type", "application/json")
		var person Person
		_ = json.NewDecoder(request.Body).Decode(&person)
		h := sha256.New()
		person.Password = (base64.StdEncoding.EncodeToString(h.Sum([]byte(person.Password))))
		collection := client.Database("Instagram").Collection("Users")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection.InsertOne(ctx, person)
		json.NewEncoder(response).Encode(map[string]string{"success": "Upload successful"})
	}
}

func GetPostEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(request.URL.Path, "/posts/"))
	var post Post
	collection := client.Database("Instagram").Collection("Posts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Post{ID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))

		return
	}
	json.NewEncoder(response).Encode(post)
}

func CreatePostEndpoint(response http.ResponseWriter, request *http.Request) {
	checkid := strings.TrimPrefix(request.URL.Path, "/posts/")
	if checkid != "" {
		GetPostEndpoint(response, request)
	} else {
		response.Header().Set("content-type", "application/json")
		var post Post
		_ = json.NewDecoder(request.Body).Decode(&post)
		post.Timestamp = time.Now().String()
		collection := client.Database("Instagram").Collection("Posts")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection.InsertOne(ctx, post)
		json.NewEncoder(response).Encode(map[string]string{"success": "Upload successful"})
	}
}

func GetAllPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	id := strings.TrimPrefix(request.URL.Path, "/posts/users/")
	collection := client.Database("Instagram").Collection("Posts")
	page, ok := request.URL.Query()["page"]

	if !ok || len(page[0]) < 1 {
		log.Println("Url Param 'page' is missing")
		return
	}
	p, _ := strconv.ParseInt(page[0], 10, 64)
	p = p - 1
	p = p * 2
	filter := bson.D{{"userid", id}}
	opts := options.Find()
	opts.SetSort(bson.M{"_id": 1})
	opts.SetSkip(p)
	opts.SetLimit(2)
	findCursor, findErr := collection.Find(context.TODO(), filter, opts)
	if findErr != nil {
		panic(findErr)
	}
	var findResults []bson.M
	if findErr = findCursor.All(context.TODO(), &findResults); findErr != nil {
		panic(findErr)
	}
	json.NewEncoder(response).Encode(findResults)
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	http.HandleFunc("/users/", CreatePersonEndpoint)
	http.HandleFunc("/posts/users/", GetAllPosts)
	http.HandleFunc("/posts/", CreatePostEndpoint)
	http.ListenAndServe(":12345", nil)
}
