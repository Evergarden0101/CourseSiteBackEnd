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

func AddApply(c *gin.Context){
	var apply domain.Apply
	if !util.BindData(c,apply){
		return
	}

	apply.Time = time.Now()
	apply.Status = constant.NONE
	apply.Id = dao.GetIncrementId("apply")
	apply.UserId = util.GetUser(c)

	dao.InsertApply(&apply)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "申请成功",
		"data": "",
	})
}

func DeleteApply(c *gin.Context){
	type jsonData struct {
		applyId string
	}

	var json jsonData
	if !util.BindData(c,json){
		return
	}

	apply := dao.GetApplyById(json.applyId)
	if (apply.UserId == util.GetUser(c)) {
		dao.DeleteApplyById(apply.Id)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "撤销成功",
			"data": "",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.ERROR,
			"msg":  "没有权限",
			"data": "",
		})
	}

}

func GetApplyByTeacher(c *gin.Context){

	list := dao.GetApplysByType(constant.TEACHER_JOIN)
    sortApply(list)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "获取申请列表成功",
		"data": list,
	})
}

func GetApply(c *gin.Context){
	type jsonData struct {
		applyId string
	}
	var json jsonData
	if !util.BindData(c,json){
		return
	}
	if !util.AdminAuth(c){
		return
	}
	apply := dao.GetApplyById(json.applyId)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "获取申请信息成功",
		"data": apply,
	})

}

func DealApply(c *gin.Context){

	type jsonData struct {
		applyId string
		dealReuslt int
	}

	var json jsonData
	if !util.BindData(c,json){
		return
	}
	if !util.AdminAuth(c){
		return
	}

	apply := dao.GetApplyById(json.applyId)
	if json.dealReuslt == 1{
		apply.Status = constant.AGREE
	}else{
		apply.Status = constant.DISAGREE
	}

	dao.UpdateApply(apply)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "审核成功",
		"data": "",
	})

}

func sortApply(list []*domain.Apply){
	for i:=0;i<len(list);i++{
		for j:=1;j<len(list);j++{
			if list[j-1].Time.Before(list[j].Time){
				apply := list[j-1]
				list[j-1] = list[j]
				list[j] = apply
			}
		}
	}
}