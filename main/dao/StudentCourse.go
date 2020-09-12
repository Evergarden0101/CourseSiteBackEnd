package dao

import (
	"awesomeProject/main/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//增
//只插入一条
func AddOneSCRelation(r *domain.StudentCourseRelation){
	collection := dataBase.Collection("studentcourserelation")
	collection.InsertOne(context.TODO(),r)
}

//删
func DeleteSCR(cid string ,sid string) bool{
	collection := dataBase.Collection("studentcourserelation")
	d := bson.M {
		"courseid":cid,
		"studentid":sid,
	}
	deleteResult,err := collection.DeleteOne(context.TODO(),d)
	if(err!=nil){
		log.Fatal(err)
		return false
	}
	fmt.Printf("deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
	return true
}

//查
func GetSCRById(cid string,sid string) bool{
	collection := dataBase.Collection("studentcourserelation")
	var scr domain.StudentCourseRelation
	d := bson.M{
		"studentid":sid,
		"courseid":cid,
	}
	collection.FindOne(context.TODO(),d).Decode(&scr)
	if(scr==(domain.StudentCourseRelation{})){
		return false
	}
	return true
}

//获取特定学生的SCR列表
func GetSCRListBySid(sid string) []*domain.StudentCourseRelation{
	collection := dataBase.Collection("studentcourserelation")
	findOptions := options.Find()
	var results []*domain.StudentCourseRelation

	cur,err := collection.Find(context.TODO(), bson.D{{"studentid",sid}},findOptions)

	if err!=nil{
		log.Fatal(err)
	}

	for cur.Next(context.TODO()){
		var elem domain.StudentCourseRelation
		err := cur.Decode(&elem)
		if err!=nil{
			log.Fatal(err)
		}

		results = append(results,&elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return results
}


