package dao

import (
	"awesomeProject/main/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

//增
//只插入一条
func AddOneSCRelation(r *domain.StudentCourseRelation){
	collection := dataBase.Collection("StudentCourseRelation")
	collection.InsertOne(context.TODO(),r)
}


//删
//func DeleteById(id string) bool{
//	collection := dataBase.Collection("topic")
//	d := bson.D{{
//		"id",id,
//	}}
//	println(id)
//	deleteResult,err := collection.DeleteOne(context.TODO(),d)
//	if(err!=nil){
//		log.Fatal(err)
//		return false
//	}
//	fmt.Printf("deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
//	return true
//}
//
//func DeleteByTitle(title string){
//	collection := dataBase.Collection("topic")
//	d := bson.D{{
//		"title",title,
//	}}
//	deleteResult,err := collection.DeleteOne(context.TODO(),d)
//	if(err!=nil){
//		log.Fatal(err)
//	}
//	fmt.Printf("deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
//}
//
//func DeleteByCourseId(courseId string){
//	collection := dataBase.Collection("topic")
//	d := bson.D{{
//		"courseid",courseId,
//	}}
//	deleteResult,err := collection.DeleteOne(context.TODO(),d)
//	if(err!=nil){
//		log.Fatal(err)
//	}
//	fmt.Printf("deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
//}
//
////改
////改title
//func ChangeTitleById(id string,newTitle string) (*domain.Topic){
//	collection := dataBase.Collection("topic")
//	var topic domain.Topic
//	filter := bson.D{{"id",id}}
//	update := bson.D{{"$set",bson.D{
//		{"title",newTitle},
//	}}}
//	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
//	if err !=nil{
//		log.Fatal(err)
//	}
//	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
//	collection.FindOne(context.TODO(),bson.D{{"title",newTitle}}).Decode(&topic)
//	return &topic
//}
//
////改Detail
//func ChangeDetailById(id string, newdetail string) (*domain.Topic){ //3
//	collection := dataBase.Collection("topic")
//	var topic domain.Topic
//	filter := bson.D{{"id",id}}
//	update := bson.D{{"$set",bson.D{
//		{"detail",newdetail},  //2
//	}}}
//	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
//	if err !=nil{
//		log.Fatal(err)
//	}
//	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
//	collection.FindOne(context.TODO(),bson.D{{"detail",newdetail}}).Decode(&topic)  //2
//	return &topic
//}
////改rule
//func ChangeruleById(id string,newrule string) (*domain.Topic){
//	collection := dataBase.Collection("topic")
//	var topic domain.Topic
//	filter := bson.D{{"id",id}}
//	update := bson.D{{"$set",bson.D{
//		{"rule",newrule},  //2
//	}}}
//	updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
//	if err !=nil{
//		log.Fatal(err)
//	}
//	fmt.Printf("matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)
//	collection.FindOne(context.TODO(),bson.D{{"rule",newrule}}).Decode(&topic)  //2
//	return &topic
//}
//
//
//
//查
func GetSCRById(cid string,sid string) (*domain.StudentCourseRelation){
	collection := dataBase.Collection("collection")
	var scr domain.StudentCourseRelation
	d := bson.M{
		"studentid":sid,
		"courseid":cid,
	}
	collection.FindOne(context.TODO(),d).Decode(&scr)
	return &scr
}

