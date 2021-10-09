package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Post struct{
	Id	primitive.ObjectID		`json:"id" bson:"_id"`
	Caption string				`json:"caption" bson: "caption"`
	Image_url  string			`json:"imageurl" bson:"imageurl"`
	Timestamp   string			`json:"timestamp" bson:"timestamp"`
	Username 	string			`json:"username" bson:"username"`
}


