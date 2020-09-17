package main

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/newrelic/go-agent"
)
import "awesomeProject/main/dao"

type Student struct {
	Name string
	Age int
}
func PrepareData(){

	var user *domain.User
	user = &domain.User{}
	user.Id ="17373273"
	user.UserName ="佐藤璃果"
	user.UserType=constant.STUDENT
	user.Password = util.Encode("123456a")
	user.Email="2541601705@qq.com"
	user.Phone="1234"
	dao.InsertUser(user)
	user = &domain.User{"123456","测试教师",util.Encode("123456a"),"1234","1234@qq.com",constant.TEACHER}
	dao.InsertUser(user)

}
func main() {


	dao.InitDB()
	var list []string
	list = append(list,"01:23")
	list = append(list,"03:33")
	fmt.Println(util.SumTime(list))
	//router.Run()
}

