package router

import (
	"awesomeProject/main/constant"
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
	http.ServeContent(c.Writer, c.Request, "test.mp4", time.Now().In(constant.CstZone), video)
}
func Run(){

	r := gin.Default()
	r.Use(NewrelicMiddleware("GoTest", "9eacdfcf41c66bfc64e9c533127f9159b0feNRAL"))

	//助教接口
	r.POST("/api/addAssistant",util.JWTAuth(),api.AddAssistant)
	r.POST("/api/deleteAssistant",util.JWTAuth(),api.DeleteAssistant)
	r.POST("/api/addMoreStudent",util.JWTAuth(),api.AddMore)

	//申请接口
	r.POST("/api/applyTeacher",util.JWTAuth(),api.ApplyTeacher)
	r.POST("/api/applyCourse",util.JWTAuth(),api.ApplyCourse)
	r.POST("/api/getTeacherApply",util.JWTAuth(),api.GetApplyByAdmin)
	r.POST("/api/getCourseApply",util.JWTAuth(),api.GetApplyByCourse)
	r.POST("/api/dealApply",util.JWTAuth(),api.DealApply)

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
	r.POST("/api/addStudent",util.JWTAuth(),api.AddStudent)
	r.POST("/api/deleteStudent",util.JWTAuth(),api.DeleteStudent)
	r.POST("/api/getStudentCourse",util.JWTAuth(),api.GetStudentCourses)
	r.POST("/api/getallCourse",util.JWTAuth(),api.GetAllCourse)
	r.POST("/api/setdetail",util.JWTAuth(),api.SetDetail)
	r.POST("/api/setrule",util.JWTAuth(),api.SetRule)
	r.POST("/api/deletecourse",util.JWTAuth(),api.DeleteCourseById)
	r.POST("/api/getTeacherCourse",util.JWTAuth(),api.GetTeacherCourse)
	r.POST("/api/getCircles",util.JWTAuth(),api.GetCircles)
	r.POST("/api/isInCourse",util.JWTAuth(),api.IsInCourse)
	r.POST("/api/getAllRelations",util.JWTAuth(),api.GetAllRelation)

	//视频接口
	r.POST("/api/getvideos",util.JWTAuth(),api.GetVideos)
	r.POST("/api/deletevideo",util.JWTAuth(),api.DeleteVideo)
	r.GET("/api/getvideostream",api.GetVideoStream)
	r.POST("/api/fileupload",util.JWTAuth(),api.FileUpload)

	//私信接口
	r.POST("/api/sendmessage",util.JWTAuth(),api.SendMessage)
	r.POST("/api/receivemessage",util.JWTAuth(),api.FindMessageByUser)
	r.POST("/api/readmessage",util.JWTAuth(),api.ReadMessage)
	r.POST("/api/getmessagenum",util.JWTAuth(),api.GetSumUnreadMessage)

	r.POST("/api/imageupload",util.JWTAuth(),api.ImageUpload)
	r.GET("/api/getimage",api.GetFile)

	r.GET("/log/readLog",util.ReadLog)


	r.Run() // listen and serve on 0.0.0.0:8080

}