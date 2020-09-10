package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InsertComment(comment *domain.Comment){
	collection := dataBase.Collection("comment")
	collection.InsertOne(context.TODO(),comment)
}

func GetComment(commentId string)(*domain.Comment){
	collection := dataBase.Collection("comment")
	var comment domain.Comment
	collection.FindOne(context.TODO(),bson.D{{"id",commentId}}).Decode(&comment)
	return &comment
}

func GetCommentsByPostId(postId string)[]*domain.Comment{
	collection := dataBase.Collection("comment")
	findOptions := options.Find()
	var results []*domain.Comment

	cur, err := collection.Find(context.TODO(), bson.D{{"postid",postId}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Comment
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

func DeleteComment(commentId string){
	collection := dataBase.Collection("comment")
	collection.DeleteOne(context.TODO(),bson.D{{"id",commentId}})


}
