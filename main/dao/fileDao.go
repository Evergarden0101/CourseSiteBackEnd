package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertFile(file *domain.File){
	collection := dataBase.Collection("file")
	collection.InsertOne(context.TODO(),file)
}

func GetFileById(id string)*domain.File{
	collection := dataBase.Collection("file")
	var file domain.File

	d := bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&file)
	return &file
}