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
	video.Time=time.Now().In(constant.CstZone)
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

	var file domain.File
	file.Time = time.Now().In(constant.CstZone)
	file.Id = dao.GetIncrementId("file")
	file.Name = m[0].Filename
	file.Url = util.GetUser(c)+"/"+m[0].Filename
	file.UserId = util.GetUser(c)

    dao.InsertFile(&file)
	util.Write(m[0],util.GetUser(c))
    file.Url = constant.IMAGE_PATH+file.Id
	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"上传成功",
		"data":file,
	})
	return

}

func GetFile(c *gin.Context){
	//user := dao.GetUserById(util.GetUser(c))

	var file domain.File
	file.Id = c.Query("id")
    file = *dao.GetFileById(file.Id)

	//if user.UserType == constant.ADMIN ||user.Id == file.UserId {
	filetream := util.Read(file.Url)
	defer filetream.Close()
	http.ServeContent(c.Writer, c.Request, file.Name, time.Now().In(constant.CstZone), filetream)
	//}else{
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": constant.DENIED,
	//		"msg":  "没有权限",
	//		"data": "",
	//	})
	//}
}