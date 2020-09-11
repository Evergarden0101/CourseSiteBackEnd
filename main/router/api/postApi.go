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
	type jsonData struct {
		Id string `json:"id"`
	}
	var userid jsonData
	error := c.BindJSON(&userid)
	if error != nil {
		log.Println(error)
	}
	list:=dao.GetPostByUserId(userid.Id)
	if len(list)>0{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功返回",
			"data": list,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "无对应帖子",
			"data": "",
		})
	}

	}
func FindPostByCourse(c *gin.Context) {
	type jsonData struct {
		Id string `json:"id"`
	}
	var courseId jsonData
	error := c.BindJSON(&courseId)
	if error != nil {
		log.Println(error)
	}
	list:=dao.GetPostByCourseId(courseId.Id)
	if len(list)>0{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功返回",
			"data": list,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "无对应帖子",
			"data": "",
		})
	}

}

func FindPostByTitle(c *gin.Context) {
	type jsonData struct {
		Title string `json:"title"`
	}
	var title jsonData
	error := c.BindJSON(&title)
	if error != nil {
		log.Println(error)
	}
	list:=dao.GetPostByTitle(title.Title)
	if len(list)>0{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功返回",
			"data": list,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "无对应帖子",
			"data": "",
		})
	}

}

func FindPostById(c *gin.Context) {
	type jsonData struct {
		Id string `json:"id"`
	}
	var userid jsonData
	error := c.BindJSON(&userid)
	if error != nil {
		log.Println(error)
	}
	ans:=dao.GetPostById(userid.Id)
	if ans!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功返回",
			"data": ans,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "无对应帖子",
			"data": "",
		})
	}
}

func ChangePostIstop(c *gin.Context) {
	type jsonData struct {
		Id string `json:"id"`
	}
	var postid jsonData
	var userid string
	//userid=util.GetUser(c)
	userid="68"
	error := c.BindJSON(&postid)
	if error != nil {
		log.Println(error)
	}
	if(dao.GetCourseById(dao.GetPostById(postid.Id).CourseId).TearchId!=userid){
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "无操作权限",
			"data": "",
		})
	}else{
		ans:=dao.UpdatePostIsTopById(postid.Id)
		if ans{
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "操作成功",
				"data": dao.GetPostById(postid.Id),
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "操作失败",
				"data": "",
			})
		}
	}

}

func ChangePostIselite(c *gin.Context) {
	type jsonData struct {
		Id string `json:"id"`
	}
	var postid jsonData
	var userid string
	//userid=util.GetUser(c)
	userid="68"
	error := c.BindJSON(&postid)
	if error != nil {
		log.Println(error)
	}
	if(dao.GetCourseById(dao.GetPostById(postid.Id).CourseId).TearchId!=userid){
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "无操作权限",
			"data": "",
		})
	}else{
		ans:=dao.UpdatePostIsEliteById(postid.Id)
		if ans{
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "操作成功",
				"data": dao.GetPostById(postid.Id),
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "操作失败",
				"data": "",
			})
		}
	}

}

