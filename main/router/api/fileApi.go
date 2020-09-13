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

func FileUpload(c *gin.Context) {

	err :=c.Request.ParseMultipartForm(100000)
	if err !=nil{
		http.Error(c.Writer,err.Error(),http.StatusInternalServerError)
		return
	}

	m:=c.Request.MultipartForm.File["video"]

	//fmt.Println(c.Request.MultipartForm.File)
	//file,err:=m[0].Open()
	//defer file.Close()
	//fmt.Println(file,err,m[0].Filename)

	//out,err :=os.Create("./upload"+m[0].Filename)
	//defer out.Close()
	//
	//_,err=io.Copy(out,file)

	var video domain.Video
	c.BindJSON(&video)
	video.Name=m[0].Filename
	video.Time=time.Now()
	video.UserId=util.GetUser(c)
	video.Id=dao.GetIncrementId("video")
	video.CourseId=c.PostForm("courseid")
	video.Path=video.CourseId+"/"+m[0].Filename
	dao.InserVideo(&video)
    util.Write(m[0],video.CourseId)

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"上传成功",
		"data":"",
		"id":video.Id,

	})
	return
}