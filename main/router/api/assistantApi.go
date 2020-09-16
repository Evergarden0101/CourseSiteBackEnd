package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddAssistant(c *gin.Context){

	var relation domain.StudentCourseRelation
	if !util.BindData(c,&relation){
		return
	}
	if !util.TeacherCourseAuth(c,relation.CourseId){
		return
	}
	dao.DeleteSCR(relation.CourseId,relation.StudentId)
	relation.Id = dao.GetIncrementId("studentcourserelation")
	relation.Type = constant.ASS

	dao.AddOneSCRelation(&relation)
	log.Println(dao.GetSCRListByCid(relation.CourseId))
	c.JSON(http.StatusOK,gin.H{
		"code": constant.SUCCESS,
		"msg": "添加助教成功",
		"data": dao.GetSCRListByCid(relation.CourseId),
	})
}

func GetAssistants(c *gin.Context){

	var course domain.Course
	if !util.BindData(c,&course){
		return
	}
	if !util.TeacherCourseAuth(c,course.Id){
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": constant.ERROR,
		"msg": "获取助教成功",
		"data": dao.GetASSListByCid(course.Id),
	})
}

func DeleteAssistant(c *gin.Context){

	var relation domain.StudentCourseRelation
	if !util.BindData(c,&relation){
		return
	}
	course := dao.GetCourse(relation.CourseId)
	if course.TeacherId == util.GetUser(c){
		dao.DeleteSCR(relation.CourseId,relation.StudentId)
		c.JSON(http.StatusOK,gin.H{
			"code": constant.SUCCESS,
			"msg": "删除助教成功",
			"data": dao.GetSCRListByCid(relation.CourseId),
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": dao.GetSCRListByCid(relation.CourseId),
		})
	}
}