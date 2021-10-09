package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanay13/go-task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostController struct{
	session *mongo.Client
}

func NewPostController(s *mongo.Client) *PostController{
	return &PostController{s}
}



func (pc PostController) GetPost(w http.ResponseWriter,r *http.Request,p httprouter.Params){

	fmt.Println("asd")
	id:= p.ByName("id")

	if !primitive.IsValidObjectID(id){
		
		w.WriteHeader(http.StatusNotFound)
	}

	oid,err:=primitive.ObjectIDFromHex(id)

	if err!=nil{
		fmt.Printf(err.Error())
	}

	u:= models.Post{}


	if err := pc.session.Database("mongo-golang").Collection("Posts").FindOne(context.TODO(),bson.M{"_id":oid}).Decode(&u); err!=nil{
		fmt.Printf(err.Error())
		w.WriteHeader(404)
		return;
	}

	// Converting Go objects into JSON is called marshalling 
	uj,err := json.Marshal(u)

	if err!=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"%s\n",uj)


}


func (pc PostController) FindAllPost(w http.ResponseWriter,r *http.Request,p httprouter.Params){

	id:= p.ByName("id")

	if !primitive.IsValidObjectID(id){
		
		w.WriteHeader(http.StatusNotFound)
	}

	oid,err:=primitive.ObjectIDFromHex(id)

	if err!=nil{
		fmt.Printf(err.Error())
	}

	u:= models.User{}


	if err := pc.session.Database("mongo-golang").Collection("users").FindOne(context.TODO(),bson.M{"_id":oid}).Decode(&u); err!=nil{
		fmt.Printf(err.Error())
		w.WriteHeader(404)
		return;
	}


	n:= u.Name


	 posts,err := pc.session.Database("mongo-golang").Collection("Posts").Find(context.TODO(),bson.M{"username":n}); 
	 
	 var allPost []bson.M

	 if err:= posts.All(context.TODO(),&allPost); err!=nil{
		fmt.Printf(err.Error())
		w.WriteHeader(404)
		return;
	 }
		

	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(http.StatusOK)

	// fmt.Println(allPost)

	fmt.Fprintf(w,"%s\n",allPost)


}

func (pc PostController) CreatePost(w http.ResponseWriter,r *http.Request,_ httprouter.Params){

	p:=models.Post{}

	// to decode the body
	json.NewDecoder(r.Body).Decode(&p)

	p.Id = primitive.NewObjectID();

	// create DB with the given name if its not there
	pc.session.Database("mongo-golang").Collection("Posts").InsertOne(context.TODO(),p)

	pj,err := json.Marshal(p)

	if err!=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"%s\n",pj)


}






func (pc PostController) DeletePost(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	id:= p.ByName("id")

	if !primitive.IsValidObjectID(id){
		w.WriteHeader(404)
		return
	}

	oid,err:= primitive.ObjectIDFromHex(id)

	if err!=nil{
		fmt.Println("Error")
	}

	if _,err:=pc.session.Database("mongo-golang").Collection("Posts").DeleteOne(context.TODO(),bson.M{"_id": oid}); err!=nil{
		w.WriteHeader(404)
		return;
	}


	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"Deleted Post",oid,"\n")


}