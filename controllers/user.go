package controllers

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanay13/go-task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct{
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController{
	return &UserController{s}
}



func (uc UserController) GetUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){

	fmt.Println("asd")
	id:= p.ByName("id")

	if !primitive.IsValidObjectID(id){
		
		w.WriteHeader(http.StatusNotFound)
	}

	oid,err:=primitive.ObjectIDFromHex(id)

	if err!=nil{
		fmt.Printf(err.Error())
	}

	u:= models.User{}


	if err := uc.session.Database("mongo-golang").Collection("users").FindOne(context.TODO(),bson.M{"_id":oid}).Decode(&u); err!=nil{
		fmt.Printf(err.Error())
		w.WriteHeader(404)
		return;
	}

	// Converting Go objects into JSON is called marshalling 
	uj,err := json.Marshal(u)

	if err!=nil{
		fmt.Println(err)
	}

	n:= u.Name



	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"%s\n",n)

	fmt.Fprintf(w,"%s\n",uj)


}

func (uc UserController) CreateUser(w http.ResponseWriter,r *http.Request,_ httprouter.Params){

	u:=models.User{}

	
	// to decode the body
	
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = primitive.NewObjectID();

	// Encrypting the password

	p:= md5.Sum([]byte(u.Password))
	u.Password = string(p[:])

	// create DB with the given name if its not there
	uc.session.Database("mongo-golang").Collection("users").InsertOne(context.TODO(),u)

	uj,err := json.Marshal(u)

	if err!=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"%s\n",uj)


}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	id:= p.ByName("id")

	if !primitive.IsValidObjectID(id){
		w.WriteHeader(404)
		return
	}

	oid,err:= primitive.ObjectIDFromHex(id)

	if err!=nil{
		fmt.Println("Error")
	}

	if _,err:=uc.session.Database("mongo-golang").Collection("users").DeleteOne(context.TODO(),bson.M{"_id": oid}); err!=nil{
		w.WriteHeader(404)
		return;
	}


	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"Deleted user",oid,"\n")


}