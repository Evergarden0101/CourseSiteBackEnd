package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateCourse(c *gin.Context){
	var course domain.Course
	error := c.BindJSON(&course)
	if error!=nil {
		log.Println(error)
	}
	findresult := dao.GetCourseByName(course.Name)
	if(findresult!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code": constant.ERROR,
			"msg": "该课程名已存在，创建失败",
			"data": "",
		})
	}else{
		dao.InsertCourse(&course)
		c.JSON(http.StatusOK,gin.H{
			"code": constant.SUCCESS,
			"msg":  "创建课程成功",
			"data": "",
		})
	}
}

//把学生批量拉入课程中
//需要传入的是课程id和学生id数组,其中学生id数组是以逗号隔开的
func IncludeStudents(c *gin.Context){
	type jsonData struct{
		Cid string `json:"cid"`
		Sid string `json:"sid"`
	}
	var json jsonData
	var scr domain.StudentCourseRelation
	err:=c.BindJSON(&json)
	if(err!=nil){
		println(err)
	} else{
		sids:=strings.Split(json.Sid,",")
		len := len(sids)
		for i:=0;i<len;i++ {
			if(dao.GetSCRById(json.Cid,sids[i])==nil){
				//构造相应的SCR结构体
				scr.Id=strconv.Itoa(i)
				scr.Type=1
				scr.CourseId=json.Cid
				scr.StudentId=json.Sid
				dao.AddOneSCRelation(&scr)
			}
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"加入学生成功",
		"data":"",
	})
}

