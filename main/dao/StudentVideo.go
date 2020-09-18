package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertStudentVideoRelation(svr *domain.StudentVideoRelation) {
	collection := dataBase.Collection("studentvideorelation")
	collection.InsertOne(context.TODO(),svr)
}

func GetSVRelation(userId string,videoId string) []*domain.StudentVideoRelation {
	collection := dataBase.Collection("studentvideorelation")
	var SVRList []*domain.StudentVideoRelation
	d := bson.M{
		"studentid":userId,
		"videoid":videoId,
	}
	cur,err := collection.Find(context.Background(),d)
	cur.All(context.Background(),&SVRList)
	log.Println(err)
	return SVRList
}

func GetSVRelationById(Id string) *domain.StudentVideoRelation {
	collection := dataBase.Collection("studentvideorelation")
	var SVR *domain.StudentVideoRelation
	d := bson.D{{
		"id",Id,
	}}
	cur,err := collection.Find(context.Background(),d)
	cur.All(context.Background(),&SVR)
	log.Println(err)
	return SVR
}

func UpdateSVRDuration(userId string,videoId string,watchTime float32)bool  {
	collection := dataBase.Collection("studentvideorelation")
	filter := bson.M{
		"studentid":userId,
		"videoid":videoId,
	}
	update := bson.D{
		{"$set",bson.D{
			{"watchtime",watchTime},
		}},
	}
	_,err:=collection.UpdateOne(context.TODO(), filter, update)
	if err!=nil{
		return false
	}
	return true
}