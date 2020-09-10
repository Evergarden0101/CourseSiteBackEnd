package router

import (
	"awesomeProject/main/router/api"
	"awesomeProject/main/util"
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
	video, err := os.Open("data1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()
	http.ServeContent(c.Writer, c.Request, "test.mp4", time.Now(), video)
}
//func Options(w *gin.Context) {
//
//	w.Header("Content-Type", "application/json")
//	w.Header("Access-Control-Allow-Origin", "*")
//	w.Header("Access-Control-Allow-Credentials", "false")
//	w.Header("Access-Control-Allow-Headers", "Authorization,Content-Type")
//	w.Header("Access-Control-Allow-Methods","GET,POST,DELETE")
//}
func Run(){

	r := gin.Default()
	r.Use(NewrelicMiddleware("GoTest", "9eacdfcf41c66bfc64e9c533127f9159b0feNRAL"))

	r.POST("/api/getcomments",api.GetComments)
    r.POST("/api/addcomment",util.JWTAuth(),api.AddComment)
	r.POST("/api/deletecomment",api.DeleteComment)

	r.POST("/api/register",api.Register)
	r.POST("/api/login",api.Login)
	r.POST("/api/modify",api.ModifyInfo)
	r.POST("/api/findpasswd",api.FindPassword)

	r.GET("/test",ServeHTTP)

	r.POST("/api/createcourse",api.CreateCourse)
	r.POST("/api/includestudents",api.IncludeStudents)


>>>>>>> main/router/router.go
	r.Run() // listen and serve on 0.0.0.0:8080

}