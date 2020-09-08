package dao

import (
	"awesomeProject/main/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

import _ "awesomeProject/main/domain"

//增
	//只插入一条
func AddOneData(s *domain.Topic){
	collection := dataBase.Collection("topic")
	collection.InsertOne(context.TODO(),s)
}
	//批量插入
func AddManyData(datas []interface{}){
	collection := dataBase.Collection("topic")
	insertManyResult,err := collection.InsertMany(context.TODO(),datas)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("inserted multiple documents:",insertManyResult.InsertedIDs)
}

//删
func DeleteById(id string) bool{
	collection := dataBase.Collection("topic")
	d := bson.D{{
		"id",id,
	}}
	println(id)
	deleteResult,err := collection.DeleteOne(context.TODO(),d)
	if(err!=nil){
		log.Fatal(err)
		return false
	}
	fmt.Printf("deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
	return true
}

func DeleteByTitle(title string){
	collection := dataBase.Collection("topic")
	d := bson.D{{
		"title",title,
	}}
	deleteResult,err := collection.DeleteOne(context.TODO(),d)
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
}

func DeleteByCourseId(courseId string){
	collection := dataBase.Collection("topic")
	d := bson.D{{
		"courseid",courseId,
	}}
	deleteResult,err := collection.DeleteOne(context.TODO(),d)
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
}

//改
	//改title
func ChangeTitleById(id string,newTitle string) (*domain.Topic){
	collection := dataBase.Collection("topic")
	var topic domain.Topic
	filter := bson.D{{"id",id}}
	update := bson.D{{"$set",bson.D{
		{"title",newTitle},
	}}}
	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
	collection.FindOne(context.TODO(),bson.D{{"title",newTitle}}).Decode(&topic)
	return &topic
}

	//改Detail
func ChangeDetailById(id string, newdetail string) (*domain.Topic){ //3
	collection := dataBase.Collection("topic")
	var topic domain.Topic
	filter := bson.D{{"id",id}}
	update := bson.D{{"$set",bson.D{
		{"detail",newdetail},  //2
	}}}
	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
	collection.FindOne(context.TODO(),bson.D{{"detail",newdetail}}).Decode(&topic)  //2
	return &topic
}
	//改rule
func ChangeruleById(id string,newrule string) (*domain.Topic){
	collection := dataBase.Collection("topic")
	var topic domain.Topic
	filter := bson.D{{"id",id}}
	update := bson.D{{"$set",bson.D{
		{"rule",newrule},  //2
	}}}
	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
	collection.FindOne(context.TODO(),bson.D{{"rule",newrule}}).Decode(&topic)  //2
	return &topic
}



//查
func GetTopicById(id string) (*domain.Topic){
	collection := dataBase.Collection("topic")
	var topic domain.Topic
	d := bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&topic)
	return &topic
}

func GetTopicByTitle(title string) (*domain.Topic){
	collection := dataBase.Collection("topic")
	var topic domain.Topic

	d := bson.D{{
		"title",title,
	}}

	collection.FindOne(context.TODO(),d).Decode(&topic)
	return &topic
}

func GetTopicByCourseId(courseId string) (*domain.Topic){
	collection := dataBase.Collection("topic")
	var topic domain.Topic

	d := bson.D{{
		"courseid",courseId,
	}}

	collection.FindOne(context.TODO(),d).Decode(&topic)
	return &topic
}