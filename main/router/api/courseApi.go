package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateCourse(c *gin.Context){
	var course domain.Course
	error := c.BindJSON(&course)
	if error!=nil {
		log.Println(error)
	}
	course.Id=dao.GetIncrementId("course")
	course.Time=time.Now()
	findresult := dao.GetCourseByName(course.Name)
	if(findresult){
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
		if(dao.GetCourseById(json.Cid)){
			sids:=strings.Split(json.Sid,",")
			len := len(sids)
			for i:=0;i<len;i++ {
				if(!dao.GetSCRById(json.Cid,sids[i])){
					//构造相应的SCR结构体
					scr.Id=dao.GetIncrementId("studentcourserelation")
					scr.Type=1
					scr.CourseId=json.Cid
					scr.StudentId=sids[i]
					dao.AddOneSCRelation(&scr)
				}
			}
			c.JSON(http.StatusOK,gin.H{
				"code":constant.SUCCESS,
				"msg":"加入学生成功",
				"data":"",
			})
		} else{
			c.JSON(http.StatusOK,gin.H{
				"code":constant.ERROR,
				"msg":"查找不到该课程，您输入的课程id为"+json.Cid,
				"data":"",
			})
		}
	}
}

//把学生踢出课程，也就是删除相应的scr数据
func DeleteStudent(c *gin.Context){
	type jsonData struct{
		Cid string `json:"cid"`
		Sid string `json:"sid"`
	}
	var json jsonData
	err:=c.BindJSON(&json)
	if(err!=nil){
		println(err)
	} else{
		if(dao.GetCourseById(json.Cid)){
			sids:=strings.Split(json.Sid,",")
			len := len(sids)
			for i:=0;i<len;i++ {
				if(dao.GetSCRById(json.Cid,sids[i])){
					if(!dao.DeleteSCR(json.Cid,sids[i])){
						c.JSON(http.StatusOK,gin.H{
							"code":constant.ERROR,
							"msg":"删除编号为:"+sids[i]+" 的学生时发生了不明错误",
							"data":"",
						})
					}
				}else{
					c.JSON(http.StatusOK,gin.H{
						"code":constant.ERROR,
						"msg":"找不到编号为:"+sids[i]+" 的学生",
						"data":"",
					})
				}
			}
			c.JSON(http.StatusOK,gin.H{
				"code":constant.SUCCESS,
				"msg":"删除学生完成",
				"data":"",
			})
		} else{
			c.JSON(http.StatusOK,gin.H{
				"code":constant.ERROR,
				"msg":"查找不到该课程，您输入的课程id为"+json.Cid,
				"data":"",
			})
		}
	}
}

