package util

import (
	"awesomeProject/main/dao"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetIncrementId(name string)(int){

	collection := dao.GetDataBase().Collection(name)
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	id := 0
	for cur.Next(context.TODO()) {
		id++
	}
	cur.Close(context.TODO())
	fmt.Println(id)
	return id

}