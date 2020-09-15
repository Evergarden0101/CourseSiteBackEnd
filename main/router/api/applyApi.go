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

func ApplyTeacher(c *gin.Context){
	var apply domain.Apply
	if !util.BindData(c,&apply){
		return
	}

	apply.Time = time.Now()
	apply.Status = constant.NONE
	apply.Id = dao.GetIncrementId("apply")
	apply.UserId = util.GetUser(c)
	apply.Type = constant.TEACHER_JOIN

	dao.InsertApply(&apply)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "申请成功",
		"data": "",
	})
}

func ApplyCourse(c *gin.Context){
	var apply domain.Apply
	if !util.BindData(c,&apply){
		return
	}

	apply.Time = time.Now()
	apply.Status = constant.NONE
	apply.Id = dao.GetIncrementId("apply")
	apply.UserId = util.GetUser(c)
	apply.Type = constant.COURSE_JOIN

	dao.InsertApply(&apply)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "申请成功",
		"data": "",
	})
}

func DeleteApply(c *gin.Context){
	type jsonData struct {
		ApplyId string `json:"applyid"`
	}

	var json jsonData
	if !util.BindData(c,&json){
		return
	}

	apply := dao.GetApplyById(json.ApplyId)
	if (apply.UserId == util.GetUser(c)) {
		dao.DeleteApplyById(apply.Id)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "撤销成功",
			"data": "",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": "",
		})
	}

}
func GetApplyByAdmin(c *gin.Context){

	if !util.AdminAuth(c){
		return
	}

	list := dao.GetApplysByType(constant.TEACHER_JOIN)

	sortApply(list)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "获取申请列表成功",
		"data": list,
	})
}

func GetApplyByCourse(c *gin.Context){
	type jsonData struct {
		CourseId string `json:"courseid"`
	}

	var json jsonData
	if !util.BindData(c,&json){
		return
	}
	if !util.TeacherCourseAuth(c,json.CourseId){
		return
	}

	list := dao.GetApplysByType(constant.COURSE_JOIN)
	var results []*domain.Apply
	for _,v := range list{
		if v.CourseId == json.CourseId{
			results = append(results,v)
		}
	}
    sortApply(results)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  "获取申请列表成功",
		"data": results,
	})
}

func GetApply(c *gin.Context){
	type jsonData struct {
		ApplyId string `json:"applyid"`
	}
	var json jsonData
	if !util.BindData(c,&json){
		return
	}
	apply := dao.GetApplyById(json.ApplyId)
	if util.AdminAuth(c) && apply.Type == constant.TEACHER_JOIN{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "获取申请信息成功",
			"data": apply,
		})
	}else if util.TeacherCourseAuth(c,apply.CourseId) && apply.Type == constant.COURSE_JOIN{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  "获取申请信息成功",
			"data": apply,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": "",
		})
	}
}

func DealApply(c *gin.Context){

	type jsonData struct {
		applyId string
		dealReuslt int
		courseId string
	}

	var json jsonData
	if !util.BindData(c,&json){
		return
	}
	apply := dao.GetApplyById(json.applyId)
	if util.AdminAuth(c) && apply.Type == constant.TEACHER_JOIN{
		if json.dealReuslt == 1{
			dao.UpdateUserType(apply.UserId)
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
		return
	}else if util.TeacherCourseAuth(c,apply.CourseId) && apply.Type == constant.COURSE_JOIN{
		if json.dealReuslt == 1{

			var scr domain.StudentCourseRelation
			scr.CourseId = apply.CourseId
			scr.StudentId = apply.UserId
			scr.Id = dao.GetIncrementId("studentcourserelation")
			scr.Type = constant.STU
			dao.AddOneSCRelation(&scr)

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
		return
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": "",
		})
	}
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