package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)
import "awesomeProject/main/domain"

func GetUserById(id string) (*domain.User)  {
	collection := dataBase.Collection("user")
	var user domain.User

	d := bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&user)
	return &user
}

func GetUserByEmail(email string) (*domain.User)  {
	collection := dataBase.Collection("user")
	var user domain.User

	d := bson.D{{
		"email",email,
	}}

	collection.FindOne(context.TODO(),d).Decode(&user)
	return &user
}

func GetUserByType(name string) []*domain.User {
	collection := dataBase.Collection("user")
	var userList []*domain.User
    d := bson.D{{
    	"userType",name,
	}}
    cur,err := collection.Find(context.Background(),d)
    cur.All(context.Background(),&userList)
    log.Println(err)
    return userList
}

func CheckId(id string) bool{
	collection := dataBase.Collection("user")
	var user domain.User

	d := bson.D{{
		"id",id,
	}}

	collection.FindOne(context.TODO(),d).Decode(&user)
	if(user == (domain.User{})) {
		return true
	}
	return false
}

func CheckEmail(email string) bool{
	collection := dataBase.Collection("user")
	var user domain.User

	d := bson.D{{
		"email",email,
	}}

	collection.FindOne(context.TODO(),d).Decode(&user)
	if(user == (domain.User{})) {
		return true
	}
	return false
}

func InsertUser(user *domain.User) {
	collection := dataBase.Collection("user")
	collection.InsertOne(context.TODO(),user)
}