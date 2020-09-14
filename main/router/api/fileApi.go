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
	var video domain.Video

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

func ImageUpload(c *gin.Context){

	err :=c.Request.ParseMultipartForm(100000)
	if err !=nil{
		http.Error(c.Writer,err.Error(),http.StatusInternalServerError)
		return
	}

	m:=c.Request.MultipartForm.File["image"]
	//fmt.Println(m[0])
	var file domain.File
	file.Time = time.Now().In(constant.CstZone)
	file.Id = dao.GetIncrementId("file")
	file.Name = m[0].Filename
	file.Url = util.GetUser(c)+"/"+m[0].Filename
    dao.InsertFile(&file)
	util.Write(m[0],util.GetUser(c))

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"上传成功",
		"data":file,
	})
	return

}

func GetFile(c *gin.Context){
	var file domain.File
	file.Id = c.Query("id")
    file = *dao.GetFileById(file.Id)
	filetream := util.Read(file.Url)
	defer filetream.Close()

	http.ServeContent(c.Writer, c.Request, file.Name, time.Now(), filetream)
}