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

func CreatePost(c *gin.Context) {

	var post domain.Post
	if !util.BindData(c,&post){
		return
	}
	if !util.TeacherCourseAuth(c,post.CourseId)&&!util.StudentCourseAuth(c,post.CourseId){
		return
	}

	post.Id=dao.GetIncrementId("post")
	post.Time=time.Now()
	post.UserId=util.GetUser(c)

	dao.InsertPost(&post)
	postlist:=dao.GetPostByCourseId(post.CourseId)
	sortPost(postlist)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "成功发布",
		"data": postlist,
	})
}

func DeletePost(c *gin.Context) {
	type PostId struct {
		Id string `json:"id"`
	}
	var postId PostId
	var userId string =util.GetUser(c)

	if !util.BindData(c,&postId){
		return
	}

	var post *domain.Post
	post=dao.GetPostById(postId.Id)
	var teacherId string =dao.GetCourse(post.CourseId).TeacherId
	var ownerId string=post.UserId
	//是否为发帖人或为负责教师删除
	if((post.UserId!=userId)&&(dao.GetCourse(post.CourseId).TeacherId!=userId)){
			c.JSON(http.StatusOK, gin.H{
				"code": constant.DENIED,
				"msg":  "无删除权限",
				"data": post,
			})
	}else{
		if (dao.DropPostById(postId.Id)){
			postlist:=dao.GetPostByCourseId(post.CourseId)
			sortPost(postlist)
			if(userId==teacherId){
				var msg domain.Message
				msg.Id=dao.GetIncrementId("message")
				msg.FromId=userId
				msg.ToId=ownerId
				msg.Detail="请注意发帖规范"
				msg.Time=time.Now()
				dao.InsertMessage(&msg)
			}
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "成功删除",
				"data": postlist,
			})
		} else{
			c.JSON(http.StatusOK, gin.H{
				"code": constant.ERROR,
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
	if !util.BindData(c,&userid){
		return
	}

	list:=dao.GetPostByUserId(userid.Id)
	sortPost(list)
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
	if !util.BindData(c,&courseId){
		return
	}

	list:=dao.GetPostByCourseId(courseId.Id)
	sortPost(list)
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
	if !util.BindData(c,&title){
		return
	}

	list:=dao.GetPostByTitle(title.Title)
	sortPost(list)
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
	if !util.BindData(c,&userid){
		return
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
	userid=util.GetUser(c)
	if !util.BindData(c,&postid){
		return
	}
	if(dao.GetCourse(dao.GetPostById(postid.Id).CourseId).TeacherId!=userid){
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
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
				"code": constant.ERROR,
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
	userid=util.GetUser(c)
	if !util.BindData(c,&postid){
		return
	}
	if(dao.GetCourse(dao.GetPostById(postid.Id).CourseId).TeacherId!=userid){
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
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
				"code": constant.ERROR,
				"msg":  "操作失败",
				"data": "",
			})
		}
	}

}

func sortPost(list []*domain.Post)  {
	for i:=0;i<len(list);i++{
		for j:=1;j<len(list);j++{
			if list[j-1].Time.Before(list[j].Time){
				post := list[j-1]
				list[j-1] = list[j]
				list[j] = post
			}
		}
	}
}

