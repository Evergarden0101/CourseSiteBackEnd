package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetVideoById(id string) (*domain.Video) {
	collection := dataBase.Collection("video")
	var video domain.Video

	d :=bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&video)
	return &video
}

func GetVideoByCourseId(courseid string) (*domain.Video) {
	collection := dataBase.Collection("video")
	var video domain.Video

	d :=bson.D{{
		"courseId",courseid,
	}}

	collection.FindOne(context.TODO(),d).Decode(&video)
	return &video
}

//func InserCourse(course *domain.Course) {
//	collection := dataBase.Collection("course")
//	collection.InsertOne(context.TODO(),course)
//}

func InserVideo(video *domain.Video) {
	collection := dataBase.Collection("video")
	collection.InsertOne(context.TODO(),video)
}


//func DeleteCourseById(id string) bool  {
//	collction :=dataBase.Collection("course")
//	_, err :=collction.DeleteOne(context.TODO(),bson.D{{"id",id}})
//
//	if  err!=nil{
//		return false
//	} else {
//		return true
//	}
//}
func DeleteVideoById(id string) bool  {
	collction :=dataBase.Collection("video")
	_, err :=collction.DeleteOne(context.TODO(),bson.D{{"id",id}})

	if  err!=nil{
		return false
	} else {
		return true
	}
}

