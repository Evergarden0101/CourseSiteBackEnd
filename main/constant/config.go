package constant

const (
	MAIL_USER = "2541601705@qq.com"
	MAIL_PASS = "lqbaygvkeuejebee"
	MAIL_HOST = "smtp.qq.com"
	MAIL_PORT = "465"
)

//用户类型
const (
	ADMIN = "admin" //管理员
	STUDENT = "student" //学生
	TEACHER = "teacher" //教师
)

//学生课程关系类型
const (
	STU = 1 //学生
	ASS = 2 //助教
)
//申请类型
const(
	TEACHER_JOIN = "TEACHER_JOIN"
	COURSE_JOIN = "COURSE_JOIN"
)
//审核结果
const(
	NONE = "none"
	AGREE = "AGREE"
	DISAGREE = "DISAGREE"
)