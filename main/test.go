package main

import (
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/newrelic/go-agent"
	"strconv"
)
import "awesomeProject/main/dao"

type Student struct {
	Name string
	Age int
}

func main() {
	dao.InitDB()
	//course := domain.Course{}
	//course.Id=strconv.Itoa(util.GetIncrementId("course"))
   //// dao.InserCourse(&course)
	//fmt.Println(dao.GetCourseById(course.Id))
	//fmt.Println(dao.DeleteCourseById("2"))

	video :=domain.Video{}
	video.Id=strconv.Itoa(util.GetIncrementId("video"))
	dao.InserVideo(&video)
	fmt.Println(dao.GetVideoById(video.Id))

	fmt.Println(dao.DeleteVideoById("2"))

	//router.Run()

}

