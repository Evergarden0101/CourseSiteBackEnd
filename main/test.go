package main

import (
	"awesomeProject/main/router"
	_ "github.com/gin-gonic/gin"
	_ "github.com/newrelic/go-agent"
)
import "awesomeProject/main/dao"

type Student struct {
	Name string
	Age int
}

func main() {


	dao.InitDB()
	//fmt.Println(util.GetIncrementId("user"))
    router.Run()

}

