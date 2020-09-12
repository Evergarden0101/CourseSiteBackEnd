package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAssistant(c *gin.Context){

	var relation domain.StudentCourseRelation
	c.BindJSON(&relation)

	relation.Id = dao.GetIncrementId("studentcourserelation")
	relation.Type = constant.ASS

	dao.AddOneSCRelation(&relation)
	c.JSON(http.StatusOK,gin.H{
		"code": constant.ERROR,
		"msg": "添加助教成功",
		"data": "",
	})
}

func GetAssistants(c *gin.Context){



}