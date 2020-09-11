package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context) {

	var post domain.Post
	error := c.BindJSON(&post)

	if error != nil {
		log.Println(error)
	}

	post.Id=dao.GetIncrementId("post")
	post.Time=time.Now()

	dao.InsertPost(&post)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "成功发布",
		"data": "",
	})
}

func DeletePost(c *gin.Context) {
	type PostId struct {
		Id string `json:"id"`
	}
	var postId PostId
	//var userId string =util.GetUser(c)
	error := c.BindJSON(&postId)
	if error != nil {
		log.Println(error)
	}
	var userId string="62"
	var post *domain.Post
	post=dao.GetPostById(postId.Id)
	//是否为发帖人或为负责教师删除
	if((post.UserId!=userId)&&(dao.GetCourse(post.CourseId).TeacherId!=userId)){
			c.JSON(http.StatusOK, gin.H{
				"code": constant.DENIED,
				"msg":  "无删除权限",
				"data": "",
			})
	}else{
		if (dao.DropPostById(postId.Id)){
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "成功删除",
				"data": "",
			})
		} else{
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "删除失败",
				"data": "",
			})
		}
	}
}

func FindPostByUser(c *gin.Context) {
	type PostId struct {
		Id string `json:"id"`
	}
	var postid PostId
	error := c.BindJSON(&postid)

	if error != nil {
		log.Println(error)
	}

	if (dao.DropPostById(postid.Id)){
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功删除",
			"data": "",
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "删除失败",
			"data": "",
		})
	}
}

