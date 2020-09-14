package router

import (
	"awesomeProject/main/router/api"
	"awesomeProject/main/util"
	"fmt"
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"log"
	"net/http"
	"os"
	"time"
)
func NewrelicMiddleware(appName string, key string) gin.HandlerFunc {

	if appName == "" || key == "" {
		return func(c *gin.Context) {}
	}

	config := newrelic.NewConfig(appName, key)
	app, err := newrelic.NewApplication(config)

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path, c.Writer, c.Request)
		defer txn.End()
		c.Next()
	}
}
func ServeHTTP(c *gin.Context) {
	video, err := os.Open("3.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()
	defer fmt.Println("sss")
	//io.Copy(c.Writer,video)
	http.ServeContent(c.Writer, c.Request, "test.mp4", time.Now(), video)
}
func Run(){

	r := gin.Default()
	r.Use(NewrelicMiddleware("GoTest", "9eacdfcf41c66bfc64e9c533127f9159b0feNRAL"))

	//申请接口
	r.POST("/api/addapply",util.JWTAuth(),api.AddApply)
	r.POST("/api/getapply",util.JWTAuth(),api.GetApply)
	r.POST("/api/getapplybyteacher",util.JWTAuth(),api.GetApplyByTeacher)
	r.POST("/api/dealapply",util.JWTAuth(),api.DealApply)
	r.POST("/api/deleteapply",util.JWTAuth(),api.DeleteApply)

	//评论接口
	r.POST("/api/getcomments",util.JWTAuth(),api.GetComments)
    r.POST("/api/addcomment",util.JWTAuth(),api.AddComment)
	r.POST("/api/deletecomment",util.JWTAuth(),api.DeleteComment)

	//用户管理接口
	r.POST("/api/register",api.Register)
	r.POST("/api/login",api.Login)
	r.POST("/api/modify",api.ModifyInfo)
	r.POST("/api/findpasswd",api.FindPassword)

	//贴子接口
	r.POST("/api/createpost",util.JWTAuth(),api.CreatePost)
	r.POST("/api/deletepost",util.JWTAuth(),api.DeletePost)
	r.POST("/api/findpostbyuser",api.FindPostByUser)
	r.POST("/api/findpostbycourse",api.FindPostByCourse)
	r.POST("/api/findpostbyid",api.FindPostById)
	r.POST("/api/findpostbytitle",api.FindPostByTitle)
	r.POST("/api/changepostistop",util.JWTAuth(),api.ChangePostIstop)
	r.POST("/api/changepostiselite",util.JWTAuth(),api.ChangePostIselite)

	//课程接口
	r.POST("/api/createcourse",util.JWTAuth(),api.CreateCourse)
	r.POST("/api/includestudents",util.JWTAuth(),api.IncludeStudents)
	r.POST("/api/deletestudents",util.JWTAuth(),api.DeleteStudent)
	r.POST("/api/getStudentCourse",util.JWTAuth(),api.GetStudentCourses)
	r.POST("/api/getallCourse",util.JWTAuth(),api.GetAllCourse)
	r.POST("/api/setdetail",util.JWTAuth(),api.SetDetail)
	r.POST("/api/setrule",util.JWTAuth(),api.SetRule)
	r.POST("/api/deletecourse",util.JWTAuth(),api.DeleteCourseById)
	r.POST("/api/getTeacherCourse",util.JWTAuth(),api.GetTeacherCourse)
	r.POST("/api/getCircles",util.JWTAuth(),api.GetCircles)

	//视频接口
	r.POST("/api/getvideos",util.JWTAuth(),api.GetVideos)
	r.POST("/api/deletevideo",util.JWTAuth(),api.DeleteVideo)
	r.GET("/api/getvideostream",api.GetVideoStream)
	r.POST("/api/fileupload",util.JWTAuth(),api.FileUpload)

	r.POST("/api/imageupload",util.JWTAuth(),api.ImageUpload)
	r.GET("/api/getimage",api.GetFile)

	r.GET("/test",util.Image)


	r.Run() // listen and serve on 0.0.0.0:8080

}