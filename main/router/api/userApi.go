package api

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"awesomeProject/main/util"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)


//注册
func Register(c *gin.Context) {

	var user domain.User
    error := c.BindJSON(&user)

    if error != nil {
		log.Println(error)
	}

	if (dao.CheckId(user.Id) && dao.CheckEmail(user.Email)) {
		dao.InsertUser(&user)
		c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  constant.REGISTER_SUCCESS,
		"data": "",
		})
    }
}

//登录
func Login(c *gin.Context){

	type jsonData struct{
		Id string `json:"id"`
		Password string  `json:"password"`
	}

	var json jsonData
	error := c.BindJSON(&json)
	fmt.Println(json)

	if error != nil {
		log.Println(error)
	}

	user := dao.GetUserById(json.Id)
	if(user.Password == json.Password) {
		generateToken(c,*user)
	}
}

func GetUser(c *gin.Context){
	claim,_ := c.Get("claims")
	claim = claim.(*domain.CustomClaims)
	c.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  constant.REGISTER_SUCCESS,
		"data": claim,
	})
}
//修改邮箱，密码，手机号等个人信息
func ModifyInfo(c *gin.Context){

	var user domain.User
	error := c.BindJSON(&user)

	if error != nil {
		log.Println(error)
	}
	oldUser := dao.GetUserById(user.Id)
	if oldUser != nil{
		oldUser.Email = user.Email
		oldUser.Phone = user.Phone
		if oldUser.Password != user.Password{
			oldUser.Password = util.Encode(user.Password)
		}
		dao.UpdateUser(oldUser)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  constant.REGISTER_SUCCESS,
			"data": "修改个人信息成功",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  constant.REGISTER_SUCCESS,
			"data": "修改失败",
		})
	}


}


//找回密码
func FindPassword(c *gin.Context){

	type jsonData struct{
		Id string `json:"id"`
		Email string  `json:"email"`
	}

	var json jsonData
	error := c.BindJSON(&json)
	fmt.Println(json)

	if error != nil {
		log.Println(error)
	}
	user := dao.GetUserById(json.Id)
	if user !=nil && user.Email == json.Email{
		str := make([]string,1)
		str[0] = user.Email
		util.SendMail(str,"重置密码邮件","重制后的密码为:123456a")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.SUCCESS,
			"msg":  constant.REGISTER_SUCCESS,
			"data": "密码已发送到邮箱",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": constant.ERROR,
			"msg":  constant.ERROR,
			"data": "学号和邮箱不匹配",
		})
	}

}

func generateToken(c *gin.Context, user domain.User) {
	j := &util.JWT{
		[]byte("newtrekWang"),
	}
	claims := domain.CustomClaims{
		user.Id,
		user.UserName,
		user.Phone,
		user.Password,
		user.UserType,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	log.Println(claims)

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := domain.LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}