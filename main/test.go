package main

import (
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
    //router.Run()

	dao.UpdatePostDetailById("1","6666666")
    dao.UpdatePostIsEliteById("1",true)
    dao.UpdatePostTitleById("1","amazing")
}

