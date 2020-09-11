package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)
import "awesomeProject/main/domain"

func GetPostById(id string) (*domain.Post)  {
	collection := dataBase.Collection("post")
	var post domain.Post

	d := bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&post)
	return &post
}


func GetPostByTitle(title string) []*domain.Post {
	collection := dataBase.Collection("post")
	var postList []*domain.Post
	d := bson.D{{
		"title",title,
	}}
	cur,err := collection.Find(context.Background(),d)
	cur.All(context.Background(),&postList)
	log.Println(err)
	return postList
}
func GetPostByUserId(userId string) []*domain.Post {
	collection := dataBase.Collection("post")
	var postList []*domain.Post
	d := bson.D{{
		"userid",userId,
	}}
	cur,err := collection.Find(context.Background(),d)
	cur.All(context.Background(),&postList)
	log.Println(err)
	return postList
}

func GetPostByCourseId(courseId string) []*domain.Post {
	collection := dataBase.Collection("post")
	var postList []*domain.Post
	d := bson.D{{
		"courseid",courseId,
	}}
	cur,err := collection.Find(context.Background(),d)
	cur.All(context.Background(),&postList)
	log.Println(err)
	return postList
}

func DropPostById (id string) bool {
	collection := dataBase.Collection("post")
	d := bson.D{{
		"id", id,
	}}
	_, err := collection.DeleteOne(context.TODO(), d)
	if (err != nil) {
		return false
	}
	return true
}
func UpdatePostDetailById(id string,detail string)bool  {
	collection := dataBase.Collection("post")
	filter := bson.D{{
		"id",id,
	}}
	update := bson.D{
		{"$set",bson.D{
			{"detail",detail},
		}},
	}
	_,err:=collection.UpdateOne(context.TODO(), filter, update)
	if err!=nil{
		return false
	}
	return true
}

func UpdatePostTitleById(id string,title string)bool  {
	collection := dataBase.Collection("post")
	filter := bson.D{{
		"id",id,
	}}
	update := bson.D{
		{"$set",bson.D{
			{"title",title},
		}},
	}
	_,err:=collection.UpdateOne(context.TODO(), filter, update)
	if err!=nil{
		return false
	}
	return true
}

func UpdatePostIsTopById(id string)bool  {
	collection := dataBase.Collection("post")
	filter := bson.D{{
		"id",id,
	}}
	istop:=!(GetPostById(id).IsTop)
	update := bson.D{
		{"$set",bson.D{
			{"istop",istop},
		}},
	}
	_,err:=collection.UpdateOne(context.TODO(), filter, update)
	if err!=nil{
		return false
	}
	return true
}

func UpdatePostIsEliteById(id string)bool  {
	collection := dataBase.Collection("post")
	filter := bson.D{{
		"id",id,
	}}
	iselite:=!(GetPostById(id).IsElite)
	update := bson.D{
		{"$set",bson.D{
			{"iselite",iselite},
		}},
	}
	_,err:=collection.UpdateOne(context.TODO(), filter, update)
	if err!=nil{
		return false
	}
	return true
}
func InsertPost(post *domain.Post) {
	collection := dataBase.Collection("post")
	collection.InsertOne(context.TODO(),post)
}