package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
func GetVideosByCourseId(courseid string)[]*domain.Video{
	collection := dataBase.Collection("video")
	findOptions := options.Find()
	var results []*domain.Video

	cur, err := collection.Find(context.TODO(), bson.D{{"courseid",courseid}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Video
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return results
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

