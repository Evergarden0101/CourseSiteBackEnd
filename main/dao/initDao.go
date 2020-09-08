package dao

import (
	context "context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

