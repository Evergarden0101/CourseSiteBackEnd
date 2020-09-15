package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddComment(c *gin.Context){
	var comment domain.Comment
	if !util.BindData(c,&comment){
		return
	}

	post := dao.GetPostById(comment.Id)
    if !util.TeacherCourseAuth(c,post.CourseId) && !util.StudentCourseAuth(c,post.CourseId){
    	return
	}

	comment.Id = dao.GetIncrementId("comment")
	comment.Time = time.Now().In(constant.CstZone)
	comment.UserId = util.GetUser(c)

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
	if !util.BindData(c,&json){
		return
	}
	post := dao.GetPostById(json.Id)
	if !util.TeacherCourseAuth(c,post.CourseId) && !util.StudentCourseAuth(c,post.CourseId){
		return
	}


	list := dao.GetCommentsByPostId(json.Id)
	userId := util.GetUser(c)
	type result struct{
		Id string `json:"id"`
		Username string `json:"username"`
		Detail string `json:"detail"`
		Isself bool `json:"isself"`
		Time string `json:"time"`
	}
	sortComment(list)

	var resultList []*result
	for _,v :=range list{
		var res result
		res.Time = v.Time.Format("2006-01-02 15:04:05")
		res.Id = v.Id
		res.Detail = v.Detail
		user := dao.GetUserById(v.UserId)
		res.Username = user.UserName
		if userId == user.Id{
			res.Isself = true
		}else{
			res.Isself = false
		}
		resultList = append(resultList,&res)

    }
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "获取评论成功",
		"data": resultList,
	})
}

func DeleteComment(c *gin.Context){
	type jsonData struct {
		Id string `json:"id"`
	}

	var json jsonData
	if !util.BindData(c,&json){
		return
	}

	userId := util.GetUser(c)
	comment := dao.GetComment(json.Id)

	if comment.UserId == userId {
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

func sortComment(list []*domain.Comment){
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