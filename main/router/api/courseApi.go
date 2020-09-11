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
	println(course)
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

//返回课程列表
//分三步，一是从SCR中获取该学生的所有课程id
//然后通过课程id获取所有课程的结构体
//最后要排一下序
func GetCoursesStruct(c *gin.Context){
	//只需要前端发送sid(学生id)
	type jsonData struct{
		Sid string `json:"sid"`
	}

	var json jsonData
	err := c.BindJSON(&json)
	if(err!=nil){
		println(err)
	}
	list := dao.GetCourseListByStudentId(json.Sid)
	sortCourse(list)

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"返回课程成功",
		"data":list,
	})
}

//获取所有圈子
func GetAllCourse(c *gin.Context){
	list := dao.GetAllCourse()
	sortCourse(list)

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"返回课程成功",
		"data":list,
	})
}

//设置课程介绍
func SetDetail(c *gin.Context){
	type jsonData struct{
		Cid string `json:"cid"`
		Detail string `json:"detail"`
	}
	var json jsonData
	err := c.BindJSON(&json)
	if(err!=nil){
		println(err)
	}
	if(dao.SetDetailByCourseId(json.Cid,json.Detail)){
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"修改课程介绍成功",
			"data":"",
		})
	} else{
		c.JSON(http.StatusOK,gin.H{
			"code":constant.ERROR,
			"msg":"修改课程介绍失败",
			"data":"",
		})
	}
}

//设置课程规则
func SetRule(c *gin.Context){
	type jsonData struct{
		Cid string `json:"cid"`
		Rule string `json:"rule"`
	}
	var json jsonData
	err := c.BindJSON(&json)
	if(err!=nil){
		println(err)
	}
	if(dao.SetRuleByCourseId(json.Cid,json.Rule)){
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"修改课程规则成功",
			"data":"",
		})
	} else{
		c.JSON(http.StatusOK,gin.H{
			"code":constant.ERROR,
			"msg":"修改课程规则失败",
			"data":"",
		})
	}
}

func sortCourse(list []*domain.Course){
	for i:=0;i<len(list);i++{
		for j:=1;j<len(list);j++{
			if list[j-1].Time.Before(list[j].Time){
				course := list[j-1]
				list[j-1] = list[j]
				list[j] = course
			}
		}
	}
}
