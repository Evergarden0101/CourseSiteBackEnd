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
	//data1:=domain.Topic{
	//	"123456",
	//	"chinese",
	//	"123456",
	//	"nothing",
	//	"nothing",
	//}
	//data2:=domain.Topic{
	//	"123457",
	//	"chinglish",
	//	"123457",
	//	"nothing",
	//	"nothing",
	//}
	//data3:=domain.Topic{
	//	"123458",
	//	"math",
	//	"123458",
	//	"nothing",
	//	"nothing",
	//}


	dao.InitDB()
    //router.Run()
    //1.
    //dao.AddOneData(&data1)
	//dao.AddOneData(&data2)
	////2.

	//data4:=domain.Topic{
	//	"123459",
	//	"music",
	//	"123459",
	//	"nothing",
	//	"nothing",
	//}
	//
	//datas := []interface{}{data3,data4}
	//dao.AddManyData(datas)
	//3.
	//print(dao.DeleteById("123456"))
	////4.
	//dao.DeleteByCourseId("123457")
	////5.
	//dao.DeleteByTitle("math")
	//
	////恢复以上3个数据
	//datass:=[] interface{}{data1,data2,data3}
	//dao.AddManyData(datass)

	////6.
	//topic:=dao.ChangeTitleById("123456","yuwen")
	//print(topic)

	////7.
	//topic:=dao.ChangeDetailById("123456","something")
	//print(topic)

	////8.
	//topic:=dao.ChangeruleById("123456","something")
	//print(topic)

	//9
	//fmt.Println(dao.GetTopicById("123456"))

	//10
	//fmt.Println(dao.GetTopicByTitle("yuwen"))

	//11
	//fmt.Println(dao.GetTopicByCourseId("123459"))
}

