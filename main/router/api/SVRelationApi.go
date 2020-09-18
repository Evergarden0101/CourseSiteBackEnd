package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetWatchTime(c *gin.Context){
	type json struct {
		VideoId string `json:"videoid"`
		WatchTime float32 `json:"watchtime"`
	}
	var js json
	if !util.BindData(c,&js){
		return
	}

	if dao.GetSVRelation(util.GetUser(c),js.VideoId)!=nil{
		dao.UpdateSVRDuration(util.GetUser(c),js.VideoId,js.WatchTime)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功设置",
			"data": dao.GetSVRelation(util.GetUser(c),js.VideoId),
		})
	}else{
		return
	}
}

func GetWatchTime(c *gin.Context){
	type json struct {
		VideoId string `json:"videoid"`
	}
	var js json
	if !util.BindData(c,&js){
		return
	}
	if dao.GetSVRelation(util.GetUser(c),js.VideoId)!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "成功获取",
			"data": dao.GetSVRelation(util.GetUser(c),js.VideoId),
		})
	}else{
		return
	}
}

