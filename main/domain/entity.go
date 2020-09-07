package domain
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

type Apply struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Title string `json:"title"`
	Message string `json:"message"`
	Type int `json:"type"`
	Status string `json:"status"`
}

type Topic struct {
	Id string `json:"id"`
	Title string `json:"title"`
	CourseId string `json:"courseId"`
	Detail string `json:"detail"`
}