package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)


func InserCourse(course *domain.Course) {
	collection := dataBase.Collection("course")
	collection.InsertOne(context.TODO(),course)
}


func GetCourseById(id string) (*domain.Course) {
	collection := dataBase.Collection("course")
	var course domain.Course

	d :=bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&course)
	return &course
}

func GetCourseByName(name string) (*domain.Course) {
	collection := dataBase.Collection("course")
	var course domain.Course

	d :=bson.D{{
		"name",name,
	}}

	collection.FindOne(context.TODO(),d).Decode(&course)
	return &course
}

func GetCourseByTeachId(teachid string) (*domain.Course) {
	collection := dataBase.Collection("course")
	var course domain.Course

	d :=bson.D{{
		"teachid",teachid,
	}}

	collection.FindOne(context.TODO(),d).Decode(&course)
	return &course
}
func DeleteCourseById(id string) bool  {
	collction :=dataBase.Collection("course")
	_, err :=collction.DeleteOne(context.TODO(),bson.D{{"id",id}})

	if  err!=nil{
		return false
	} else {
		return true
	}
}
