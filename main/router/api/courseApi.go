package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func CreateCourse(c *gin.Context){

	if !util.TeacherAuth(c){
		return
	}

	var course domain.Course
	if !util.BindData(c,&course){
		return
	}

	course.Id=dao.GetIncrementId("course")
	course.Time=time.Now().In(constant.CstZone)
	course.TeacherId = util.GetUser(c)

	findresult := dao.GetCourseByName(course.Name)
	if(findresult){
		c.JSON(http.StatusOK,gin.H{
			"code": constant.ERROR,
			"msg": "该课程名已存在，创建失败",
			"data": "",
		})
	}else{
		dao.InsertCourse(&course)
		GetCircles(c)
	}
}

func DeleteCourseById(c *gin.Context){
	var course domain.Course
	if !util.BindData(c,&course){
		return
	}
	if !util.TeacherCourseAuth(c,course.Id){
		return
	}

	if(dao.DeleteCourseById(course.Id)){
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"删除课程成功",
			"data":"",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":constant.ERROR,
			"msg":"删除课程失败",
			"data":"",
		})
	}

}

func AddMore(c *gin.Context){
	err :=c.Request.ParseMultipartForm(100000)
	if err !=nil{
		http.Error(c.Writer,err.Error(),http.StatusInternalServerError)
		return
	}
	m:=c.Request.MultipartForm.File["file"]
	courseId := c.PostForm("courseid")
	util.Write(m[0],util.GetUser(c))
	var scrList []*domain.StudentCourseRelation
    userList := util.AnalyzeExcel(util.GetUser(c)+"/"+m[0].Filename)
    for _,v := range userList{
    	if (dao.CheckId(v.Id))||dao.GetSCRById(courseId,v.Id){
    		continue
		}
		scr := domain.StudentCourseRelation{}
		scr.CourseId = courseId
		scr.StudentId = v.Id
		scr.StudentName = v.UserName
		scr.Id = dao.GetIncrementId("studentcourserelation")
		scr.Type = constant.STU
		dao.AddOneSCRelation(&scr)
		scrList = append(scrList,&scr)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"批量导入成功",
		"data":scrList,
	})
}

//把学生批量拉入课程中
//需要传入的是课程id和学生id数组,其中学生id数组是以逗号隔开的
func AddStudent(c *gin.Context){
	type jsonData struct{
		Cid string `json:"courseid"`
		Sid string `json:"studentid"`
	}
	var json jsonData
	if !util.BindData(c,&json){
		return
	}
	if !util.TeacherCourseAuth(c,json.Cid){
		return
	}

	var scr domain.StudentCourseRelation

	if(dao.GetCourseById(json.Cid)){
		sids:=strings.Split(json.Sid,",")
		len := len(sids)
		for i:=0;i<len;i++ {
			if(!dao.GetSCRById(json.Cid,sids[i])){
					//构造相应的SCR结构体
				scr.Id=dao.GetIncrementId("studentcourserelation")
				scr.Type=1
				scr.CourseId=json.Cid
				scr.StudentId=sids[i]
				dao.AddOneSCRelation(&scr)
			}
		}
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"加入学生成功",
			"data":"",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.ERROR,
			"msg":  "查找不到该课程，您输入的课程id为" + json.Cid,
			"data": "",
		})
	}


}

//把学生踢出课程，也就是删除相应的scr数据
func DeleteStudent(c *gin.Context) {
	type jsonData struct {
		Cid string `json:"courseid"`
		Sid string `json:"studentid"`
	}
	var json jsonData
	if !util.BindData(c, &json) {
		return
	}
	if !util.TeacherCourseAuth(c,json.Cid){
		return
	}

	if (dao.GetCourseById(json.Cid)) {
		if (dao.DeleteSCR(json.Cid, json.Sid)) {
			c.JSON(http.StatusOK, gin.H{
				"code": constant.SUCCESS,
				"msg":  "已成功把该学生移出课程",
				"data": dao.GetSCRListByCid(json.Cid),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": constant.ERROR,
				"msg":  "把学生移除课程失败，该学生id为" + json.Sid,
				"data": "",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": constant.ERROR,
			"msg":  "查找不到该课程，您输入的课程id为" + json.Cid,
			"data": "",
		})
	}
}

//返回课程列表
//分三步，一是从SCR中获取该学生的所有课程id
//然后通过课程id获取所有课程的结构体
//最后要排一下序
func GetStudentCourses(c *gin.Context){
	//只需要前端发送sid(学生id)
	type jsonData struct{
		Sid string `json:"studentid"`
	}

	var json jsonData
	if !util.BindData(c, &json) {
		return
	}

	list := dao.GetCourseListByStudentId(json.Sid)
	sortCourse(list)

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"返回学生课程成功",
		"data":list,
	})
}

func GetCircles(c *gin.Context){

	userId := util.GetUser(c)
	user := dao.GetUserById(userId)
	allList := dao.GetAllCourse()
	sortCourse(allList)
	type result struct {
		AllList []*domain.Course `json:"allList"`
		TeacherList []*domain.Course `json:"teacherList"`
	}
	var res result
	res = result{}
	if user.UserType == constant.STUDENT{
		teacherList := dao.GetCourseListByStudentId(userId)
		sortCourse(teacherList)
		res.AllList = allList
		res.TeacherList =teacherList
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"课程列表成功",
			"data": res,
		})


	}else{
		teacherList := dao.GetTeacherCourse(userId)
		sortCourse(teacherList)
		res.AllList = allList
		res.TeacherList =teacherList
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"课程列表成功",
			"data": res,
		})
	}
}

//获取老师的所有课程
func GetTeacherCourse(c *gin.Context){
	type jsonData struct{
		Teacherid string `json:"teacherid"`
	}
	var json jsonData
	if !util.BindData(c, &json) {
		return
	}

	list := dao.GetTeacherCourse(json.Teacherid)
	sortCourse(list)

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"返回老师课程成功",
		"data":list,
	})
}

//获取所有圈子
func GetAllCourse(c *gin.Context){
	list := dao.GetAllCourse()
	sortCourse(list)

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"返回课程成功",
		"data":list,
	})
}

//设置课程介绍
func SetDetail(c *gin.Context){
	type jsonData struct{
		Cid string `json:"cid"`
		Detail string `json:"detail"`
	}
	var json jsonData
	if !util.BindData(c, &json) {
		return
	}
	if !util.TeacherCourseAuth(c,json.Cid){
		return
	}

	if(dao.SetDetailByCourseId(json.Cid,json.Detail)){
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"修改课程介绍成功",
			"data":"",
		})
	} else{
		c.JSON(http.StatusOK,gin.H{
			"code":constant.ERROR,
			"msg":"修改课程介绍失败",
			"data":"",
		})
	}
}

//设置课程规则
func SetRule(c *gin.Context){
	type jsonData struct{
		Cid string `json:"cid"`
		Rule string `json:"rule"`
	}
	var json jsonData
	if !util.BindData(c, &json) {
		return
	}
	if !util.TeacherCourseAuth(c,json.Cid){
		return
	}

	if(dao.SetRuleByCourseId(json.Cid,json.Rule)){
		c.JSON(http.StatusOK,gin.H{
			"code":constant.SUCCESS,
			"msg":"修改课程规则成功",
			"data":dao.GetCourse(json.Cid).Rule,
		})
	} else{
		c.JSON(http.StatusOK,gin.H{
			"code":constant.ERROR,
			"msg":"修改课程规则失败",
			"data":"",
		})
	}
}

func IsInCourse(c *gin.Context){
	type jsonData struct {
		Cid string `json:"courseid"`
	}
	var json jsonData
	if !util.BindData(c, &json){
		return
	}

	var result int
	result = 0
	if(dao.GetSCR(json.Cid,util.GetUser(c)).Type == constant.ASS || dao.GetCourse(json.Cid).TeacherId == util.GetUser(c)){
		result = 3
	}else if(dao.GetApplyByCourseId(json.Cid,util.GetUser(c)).Status == constant.NONE) {
		result = 1
	}else if(dao.GetSCR(json.Cid,util.GetUser(c)).Type == constant.STU){
		result = 2
	}

	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"",
		"data":result,
	})
}

func GetAllRelation(c *gin.Context){

	type jsonData struct{
		Cid string `json:"cid"`
	}
	var json jsonData
	if !util.BindData(c, &json) {
		return
	}
	if !util.TeacherCourseAuth(c,json.Cid){
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":constant.SUCCESS,
		"msg":"",
		"data":dao.GetSCRListByCid(json.Cid),
	})
}

func sortCourse(list []*domain.Course){
	for i:=0;i<len(list);i++{
		for j:=1;j<len(list);j++{
			if list[j-1].Time.Before(list[j].Time){
				course := list[j-1]
				list[j-1] = list[j]
				list[j] = course
			}
		}
	}
}
