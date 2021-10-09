package main

import (
	"context"
	"fmt"
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/tanay13/go-task/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// entry point function

func main(){

	r := httprouter.New() 

	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser )
	r.DELETE("/user/:id", uc.DeleteUser)
	r.GET("/post/:id",pc.CreatePost)
	r.POST("/post", pc.CreatePost )
	r.GET("/posts/user/:id",pc.FindAllPost)

	http.ListenAndServe("localhost:8080", r)


}

// *mgo.Session is the return type of the function

func getSession() *mongo.Client{

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client

}