package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertMessage(comment *domain.Message){
	comment.TimeString = comment.Time.Format("2006-01-02 15:04:05")
	collection := dataBase.Collection("message")
	collection.InsertOne(context.TODO(),comment)
}

func GetMessageByToUserId(toId string) []*domain.Message {
	collection := dataBase.Collection("message")
	var messageList []*domain.Message
	d := bson.D{{
		"toid",toId,
	}}
	cur,err := collection.Find(context.Background(),d)
	cur.All(context.Background(),&messageList)
	log.Println(err)
	return messageList
}

func GetMessageById(id string) (*domain.Message)  {
	collection := dataBase.Collection("message")
	var message domain.Message

	d := bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&message)
	return &message
}

//func GetMessageByFromUserId(fromId string) []*domain.Message {
//	collection := dataBase.Collection("message")
//	var messageList []*domain.Message
//	d := bson.D{{
//		"fromid",fromId,
//	}}
//	cur,err := collection.Find(context.Background(),d)
//	cur.All(context.Background(),&messageList)
//	log.Println(err)
//	return messageList
//}

func ModifyReadById(Id string) bool{
	collection := dataBase.Collection("message")
	filter := bson.D{{
		"id",Id,
	}}
	update := bson.D{
		{"$set",bson.D{
			{"read",true},
		}},
	}
	_,err:=collection.UpdateOne(context.TODO(), filter, update)
	if err!=nil||(GetMessageById(Id).FromId==""){
		return false
	}
	return true
}