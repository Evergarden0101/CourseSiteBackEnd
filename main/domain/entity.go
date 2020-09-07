package domain
/**
数据库实体类
名称对应均按照json的名称，与变量名无关
 */
type User struct {
	Id string `json:"id"`
	UserName string  `json:"userName"`
	Password string  `json:"password"`
	Email string `json:"email"`
	UserType string  `json:"userType"`
}

type Service struct {
	Id string `json:"id"`
	Detail string `json:"detail"`
}

type Relation struct {
	UserId string `json:"userId"`
	ServiceId string `json:"serviceId"`
	Type int `json:"type"`
}

type Apply struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	ServiceId string `json:"serviceId"`
	Type int `json:"type"`
	Status string `json:"status"`
}