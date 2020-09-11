package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func FileUpload(c *gin.Context) {

	//c.Writer
	//
	//c.Request

	err :=c.Request.ParseMultipartForm(100000)
	if err !=nil{
		http.Error(c.Writer,err.Error(),http.StatusInternalServerError)
		return
	}

	//file,header,err:=c.Request.FormFile("video")
	//
	//
	//filename :=header.Filename
	m:=c.Request.MultipartForm.File["video"]

	fmt.Println(c.Request.MultipartForm.File)
	file,err:=m[0].Open()
	defer file.Close()
//	filename:=c.Request.MultipartForm.Value["filename"]

	fmt.Println(file,err,m[0].Filename)
	f:= func(c rune) bool{
		if(c=='.'){
			return true
		}else {
			return false
		}
	}
	result:=strings.FieldsFunc(m[0].Filename,f)
	fmt.Println(result)
	//var i  =0
	//for i,_ :=range result{
	//
	//	if (result[i] == "mp4"){
	//		fmt.Println("mp4")
	//		i=1
	//	}
	//}
	//if(i==0 ){
	//	fmt.Println("not mp4")
	//}

	out,err :=os.Create("./upload"+m[0].Filename)
	defer out.Close()

	_,err=io.Copy(out,file)

//	c.String(http.StatusCreated,"upload sucessful")

	var video domain.Video
	c.BindJSON(&video)
	video.Name=m[0].Filename
	video.Time=time.Now()
	video.UserId=util.GetUser(c)
	video.Id=dao.GetIncrementId("video")
	video.Path="./upload"+m[0].Filename
	video.CourseId=c.PostForm("courseid")
	dao.InserVideo(&video)
	fmt.Println(video.CourseId)
	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"申请成功",
		"data":"",
		"id":video.Id,

	})
	return

}