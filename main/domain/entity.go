package domain

import "time"

/**
数据库实体类
名称对应均按照json的名称，与变量名无关
 */
type User struct {
	Id string `json:"id"`
	UserName string  `json:"userName"`
	Password string  `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	UserType string  `json:"userType"`
}

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TearchId string `json:"teacherId"`
	Detail string `json:"detail"`
	Time time.Time `json:"time"`
}

type StudentCourseRelation struct {
	Id string `json:"id"`
	StudentId string `json:"studentId"`
	CourseId string `json:"courseId"`
	Type int `json:"type"`
}

type Video struct {
	Id string `json:"id"`
	CourseId string `json:"courseId"`
	Name string `json:"name"`
	Detail string `json:"detail"`
	Path string `json:"path"`
	Time time.Time `json:"time"`
}


type Apply struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Title string `json:"title"`
	Message string `json:"message"`
	Type int `json:"type"`
	Status string `json:"status"`
	Time time.Time `json:"time"`
}

type Topic struct {
	Id string `json:"id"`
	Title string `json:"title"`
	CourseId string `json:"courseId"`
	Detail string `json:"detail"`
	Rule string `json:"rule"`
}

type StudentTopicRelation struct {
	Id string `json:"id"`
	StudentId string `json:"studentId"`
	TopicId string `json:"topicId"`
	Type int `json:"type"`
	Score int `json:"score"`
}

type Post struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	IsTop bool `json:"isTop"`
	IsElite bool `json:"isElite"`
	Time time.Time `json:"time"`
}

type Message struct {
	Id string `json:"id"`
	FromId string `json:"fromId"`
	ToId string `json:"toId"`
	Detail string `json:"detail"`
	Time time.Time `json:"time"`
}

type Like struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	PostId string `json:"postId"`
	Time time.Time `json:"time"`
}

type Comment struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	PostId string `json:"postId"`
	Detail string `json:"detail"`
	Time time.Time `json:"time"`
}