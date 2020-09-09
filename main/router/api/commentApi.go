package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddComment(c *gin.Context){
	var comment domain.Comment
	err := c.BindJSON(&comment)
	if err !=nil{
		fmt.Println(err)
	}
	comment.Time = time.Now()
	dao.InsertComment(&comment)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "评论成功",
		"data": "",
	})
}

func GetComments(c *gin.Context){
	type jsonData struct {
		Id string `json:"id"`
	}
	var json jsonData
	c.BindJSON(&json)
	list := dao.GetCommentsByPostId(json.Id)
	sort(list)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "获取评论成功",
		"data": list,
	})
}

func DeleteComment(c *gin.Context){
	type jsonData struct {
		Id string `json:"id"`
		UserId string `json:"userId"`
	}

	var json jsonData
	c.BindJSON(&json)
	comment := dao.GetComment(json.Id)
	if comment.UserId == json.UserId {
		dao.DeleteComment(comment.Id)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "删除评论成功",
			"data": "",
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.ERROR,
			"msg":  "没有权限",
			"data": "",
		})
	}
}

func sort(list []*domain.Comment){
	for i:=0;i<len(list);i++{
		for j:=1;j<len(list);j++{
			if list[j-1].Time.Before(list[j].Time){
				comment := list[j-1]
				list[j-1] = list[j]
				list[j] = comment
			}
		}
	}
}