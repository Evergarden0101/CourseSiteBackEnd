package main

import (
	"awesomeProject/main/router/api"
	"awesomeProject/main/util"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/newrelic/go-agent"
	newrelic "github.com/newrelic/go-agent"
)
import "awesomeProject/main/dao"

type Student struct {
	Name string
	Age int
}

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


func main() {
	dao.InitDB()

	r := gin.Default()
	r.Use(NewrelicMiddleware("GoTest", "9eacdfcf41c66bfc64e9c533127f9159b0feNRAL"))

	r.POST("/api/register",api.Register)
	r.POST("/api/login",api.Login)
	r.GET("/api/getUser",util.JWTAuth(),api.GetUser)
	r.GET("/ping",func(c *gin.Context) {
		str := make([]string,1)
		str[0] = "2541601705@qq.com"
		util.SendMail(str,"1234","1234")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080


}

