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

func SendMessage(c *gin.Context) {
	var msg domain.Message   //toid detail
	if !util.BindData(c,&msg){
		return
	}
	msg.Id=dao.GetIncrementId("message")
	msg.Time=time.Now()
	msg.FromId=util.GetUser(c)
	msg.Read=false

	dao.InsertMessage(&msg)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "成功发送",
		"data": "",
	})
}


func FindMessageByUser(c *gin.Context) {
	//type jsonData struct {
	//	toId string `json:"toid"`
	//}
	//var touserid jsonData
	//if !util.BindData(c,&touserid){
	//	return
	//}
	touserid:=util.GetUser(c)
	list:=dao.GetMessageByToUserId(touserid)
	sortMessage(list)
	if len(list)>0{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功返回",
			"data": list,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "暂无私信",
			"data": "",
		})
	}
}

func ReadMessage(c *gin.Context) {
	type jsonData struct {
		Id string `json:"id"`
	}
	var id jsonData
	if !util.BindData(c,&id){
		return
	}
	if dao.ModifyReadById(id.Id){
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "已阅成功",
			"data": unreadMessageNum(dao.GetMessageByToUserId(dao.GetMessageById(id.Id).ToId)),
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "已阅失败",
			"data": "",
		})
	}
}

func sortMessage(list []*domain.Message)  {
	for i:=0;i<len(list);i++{
		for j:=1;j<len(list);j++{
			if list[j-1].Time.Before(list[j].Time){
				msg := list[j-1]
				list[j-1] = list[j]
				list[j] = msg
			}
		}
	}
}

func unreadMessageNum(list []*domain.Message)int{
	var ans int
	for i:=0;i<len(list);i++{
		if list[i].Read==false{
			ans++
		}
	}
	return ans
}
