package util

import (
	"awesomeProject/main/dao"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)
type increment struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}
func GetIncrementId(name string)(string){

	collection := dao.GetDataBase().Collection("incrementId")
	var incre increment
	d := bson.D{{
		"name",name,
	}}

	val := collection.FindOne(context.TODO(),d)
	val.Decode(&incre)

	filter := bson.D{{"name", name}}
	update := bson.D{
		{"$set", bson.D{
			{"number", incre.Number +1},
		}},
	}
	collection.UpdateOne(context.TODO(), filter, update)
	return strconv.Itoa(incre.Number)

}