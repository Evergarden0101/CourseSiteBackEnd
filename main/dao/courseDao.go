package dao

import (
	"awesomeProject/main/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)


func InsertCourse(course *domain.Course) {
	collection := dataBase.Collection("course")
	_,err := collection.InsertOne(context.TODO(),course)
	fmt.Println(err)
}


func GetCourseById(id string) bool {
	collection := dataBase.Collection("course")
	var course domain.Course

	d :=bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&course)
	if(course==(domain.Course{})){
		return false
	}
	return true
}

func GetCourse(id string) *domain.Course {
	collection := dataBase.Collection("course")
	var course domain.Course

	d :=bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&course)
	return &course
}

func GetCourseByName(name string) bool {
	collection := dataBase.Collection("course")
	var course domain.Course

	d :=bson.D{{
		"name",name,
	}}

	collection.FindOne(context.TODO(),d).Decode(&course)
	if(course==(domain.Course{})){
		return false
	}
	return true
}

//获取一个学生的所有课程
//先获取SCR结构体，然后通过相应的课程id获取相应的课程结构体
func GetCourseListByStudentId(sid string) []*domain.Course{
	collection := dataBase.Collection("course")
	var result []*domain.Course

	SCRlist := GetSCRListBySid(sid)
	len:= len(SCRlist)
	var elem domain.Course
	for i:=0;i<len;i++{
		collection.FindOne(context.TODO(),bson.D{{"courseid",SCRlist[i].CourseId}}).Decode(&elem)
		result = append(result,&elem)
	}
	return result
}

func GetTeacherCourse(teachid string) ([]*domain.Course) {
	collection := dataBase.Collection("course")
	findOptions := options.Find()
	var results []*domain.Course

	cur, err := collection.Find(context.TODO(), bson.D{{"teacherid",teachid}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Course
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



func DeleteCourseById(id string) bool  {
	collction :=dataBase.Collection("course")
	_, err :=collction.DeleteOne(context.TODO(),bson.D{{"id",id}})

	if  err!=nil{
		return false
	} else {
		return true
	}
}

//获取所有课程
func GetAllCourse() []*domain.Course{
	collection := dataBase.Collection("course")
	findOptions := options.Find()
	var results []*domain.Course

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Course
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

func SetDetailByCourseId(id string, newdetail string) bool{
	collection := dataBase.Collection("course")
	filter := bson.D{{"id",id}}
	update := bson.D{{"$set",bson.D{
		{"rule",newdetail},  //2
	}}}
	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
	if err !=nil{
		log.Fatal(err)
		return false
	}
	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
	return true
}

func SetRuleByCourseId(id string, newrule string) bool{
	collection := dataBase.Collection("course")
	filter := bson.D{{"id",id}}
	update := bson.D{{"$set",bson.D{
		{"rule",newrule},  //2
	}}}
	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
	if err !=nil{
		log.Fatal(err)
		return false
	}
	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
	return true
}


