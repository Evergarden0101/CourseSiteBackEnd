package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetVideos(c *gin.Context)  {
	type jsonData struct {
		Courseid string `json:"courseid"`
	}
	var json jsonData
	if !util.BindData(c,&json){
		return
	}
	list :=dao.GetVideosByCourseId(json.Courseid)
	sortVideo(list)
	c.JSON(http.StatusOK,gin.H{
		"code": constant.SUCCESS,
		"msg":  "获取视频列表成功",
		"data": list,
	})

}
func DeleteVideo(c *gin.Context){
	type jsonData struct {
		Id string `json:"id"`
		UserId string `json:"userId"`
	}

	var json jsonData
	if !util.BindData(c,&json){
		return
	}
	video := dao.GetVideoById(json.Id)
	if video.UserId ==util.GetUser(c) {
		dao.DeleteVideoById(video.Id)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "删除视频成功",
			"data": "",
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.ERROR,
			"msg":  "没有权限",
			"data": "",
		})
	}
}

func sortVideo(list []*domain.Video){
	for i:=0;i<len(list);i++{
		for j:=1;j<len(list);j++{
			if list[j-1].Time.Before(list[j].Time){
				comment := list[j-1]
				list[j-1] = list[j]
				list[j] = comment
			}
		}
	}

}
func GetVideoStream(c *gin.Context)  {
	type jsonData struct {
		Id string `json:"id"`
	}
	var json jsonData
	json.Id = c.Query("id")
	video := dao.GetVideoById(json.Id)

	log.Println(json.Id)

	videostream := util.Read(video.Path)
	defer videostream.Close()

	if(len(dao.GetSVRelation(util.GetUser(c),json.Id))==0){
		var SVR domain.StudentVideoRelation
		SVR.StudentId=util.GetUser(c)
		SVR.VideoId=json.Id
		SVR.Id=dao.GetIncrementId("studentvideorelation")
		SVR.WatchTime=0.0
		SVR.LastWatch=time.Now()
		dao.InsertStudentVideoRelation(&SVR)
	}

	http.ServeContent(c.Writer, c.Request, json.Id+".mp4", time.Now().In(constant.CstZone), videostream)

}