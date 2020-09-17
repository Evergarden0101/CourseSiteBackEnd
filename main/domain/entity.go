package domain

import "time"

/**
数据库实体类
名称对应均按照json的名称，与变量名无关
 */
type User struct {
	Id string `json:"id"`
	UserName string  `json:"username"`
	Password string  `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	UserType string  `json:"usertype"`
}

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TeacherId string `json:"teacherid"`
	Detail string `json:"detail"`
	Rule string `json:"rule"`
	Time time.Time `json:"time"`
}

type StudentCourseRelation struct {
	Id string `json:"id"`
	StudentId string `json:"studentid"`
	StudentName string `json:"studentname"`
	CourseId string `json:"courseid"`
	Type int `json:"type"`
}

type StudentVedioRelation struct {
	Id string `json:"id"`
	StudentId string `json:"studentid"`
	StudentName string `json:"studentname"`
	CourseId string `json:"courseid"`
	Type int `json:"type"`
}

type Video struct {
	Id string `json:"id"`
	CourseId string `json:"courseid"`
	UserId string `json:"userid"`
	Name string `json:"name"`
	Detail string `json:"detail"`
	Path string `json:"path"`
	Time time.Time `json:"time"`
}


type Apply struct {
	Id string `json:"id"`
	UserId string `json:"userid"`
	UserName string `json:"username"`
	CourseId string `json:"courseid"`
	ImageId string `json:"imageid"`
	Type string `json:"type"`
	Status string `json:"status"`
	Time time.Time `json:"time"`
}

type Topic struct {
	Id string `json:"id"`
	Title string `json:"title"`
	CourseId string `json:"courseid"`
	Detail string `json:"detail"`
	Rule string `json:"rule"`
	TeacherId string `json:"teacherid"`
}

type StudentTopicRelation struct {
	Id string `json:"id"`
	StudentId string `json:"studentid"`
	TopicId string `json:"topicId"`
	Type int `json:"type"`
	Score int `json:"score"`
}

type Post struct {
	Id string `json:"id"`
	UserId string `json:"userid"`
	CourseId string `json:"courseid"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	IsTop bool `json:"istop"`
	IsElite bool `json:"iselite"`
	Time time.Time `json:"time"`
}

type Message struct {
	Id string `json:"id"`
	FromId string `json:"fromid"`
	ToId string `json:"toid"`
	Topic string `json:"topic"`
	Detail string `json:"detail"`
	Read bool `json:"read"`
	Time time.Time `json:"time"`
	TimeString string `json:"timestring"`
}

type Like struct {
	Id string `json:"id"`
	UserId string `json:"userid"`
	PostId string `json:"postid"`
	Time time.Time `json:"time"`
}

type Comment struct {
	Id string `json:"id"`
	UserId string `json:"userid"`
	PostId string `json:"postid"`
	Detail string `json:"detail"`
	Time time.Time `json:"time"`
}

type File struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Time time.Time `json:"time"`
	Name string `json:"name"`
	UserId string `json:"userid"`
}