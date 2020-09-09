package router

import (
	"awesomeProject/main/router/api"
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
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

func Run(){

	r := gin.Default()
	r.Use(NewrelicMiddleware("GoTest", "9eacdfcf41c66bfc64e9c533127f9159b0feNRAL"))

	r.POST("/api/register",api.Register)
	r.POST("/api/login",api.Login)
	r.POST("/api/modify",api.ModifyInfo)
	r.POST("/api/findpasswd",api.FindPassword)
	//r.GET("/api/getUser",util.JWTAuth(),util.GetUser)
	//r.GET("/ping",func(c *gin.Context) {
	//	str := make([]string,1)
	//	str[0] = "2541601705@qq.com"
	//	util.SendMail(str,"1234","1234")
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run() // listen and serve on 0.0.0.0:8080

}