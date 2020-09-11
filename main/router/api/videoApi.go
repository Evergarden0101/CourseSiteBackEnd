package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func GetVideos(c *gin.Context)  {
	type jsonData struct {
		Courseid string `json:"courseid"`
	}
	var json jsonData
	c.BindJSON(&json)
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
	c.BindJSON(&json)
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
		UserId string `json:"userId"`
	}

	var json jsonData
	c.BindJSON(&json)
	video := dao.GetVideoById(json.Id)
	videostream,err:=os.Open(video.Path)
	defer videostream.Close()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "播放视频成功",
		"data": "",
		"videostream":*videostream,
	})
	//http.ServeContent(c.Writer, c.Request, "test.mp4", time.Now(), videostream)
}
func ServeHTTP(c *gin.Context) {
	video, err := os.Open("3.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()
	defer fmt.Println("sss")
	http.ServeContent(c.Writer, c.Request, "test.mp4", time.Now(), video)
}