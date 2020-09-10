package dao

import (
	context "context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)
var dataBase *mongo.Database
func InitDB(){
	clientOptions := options.Client().ApplyURI("mongodb://science:123456a@39.105.206.72:27017/science")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	dataBase = client.Database("science")
}

func GetDataBase()(*mongo.Database){
	return dataBase
}

type increment struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}
func GetIncrementId(name string)(string){

	collection := GetDataBase().Collection("incrementid")
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