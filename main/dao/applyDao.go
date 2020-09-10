package dao

import (
	"awesomeProject/main/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetApplysByType(applyType string)[]*domain.Apply{
	collection := dataBase.Collection("apply")
	findOptions := options.Find()
	var results []*domain.Apply

	cur, err := collection.Find(context.TODO(), bson.D{{"type",applyType}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Apply
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

func GetApplyById(id string)*domain.Apply{
	collection := dataBase.Collection("apply")
	var apply domain.Apply

	d := bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&apply)
	return &apply
}

func UpdateApply(apply *domain.Apply){
	collection := dataBase.Collection("apply")
	filter := bson.D{{"id", apply.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", apply.Status},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateResult)


}

func InsertApply(apply *domain.Apply){
	collection := dataBase.Collection("apply")
	_,err := collection.InsertOne(context.TODO(),apply)
	if err !=nil {
		fmt.Println(err)
	}
}

