package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanay13/go-task/controllers"
	"gopkg.in/mgo.v2"
)

// entry point function

func main(){

	r := httprouter.New() 

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser )
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)


}

// *mgo.Session is the return type of the function

func getSession() *mgo.Session{

	s,err := mgo.Dial("mongodb://localhost:27017")

	if err!=nil{
		panic(err)
	}
	return s;

}