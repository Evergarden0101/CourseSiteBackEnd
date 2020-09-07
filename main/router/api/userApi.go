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
		Email string `json:"email"`
		Password string  `json:"password"`
	}

	var json jsonData
	error := c.BindJSON(&json)
	fmt.Println(json)

	if error != nil {
		log.Println(error)
	}

	user := dao.GetUserByEmail(json.Email)
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

func generateToken(c *gin.Context, user domain.User) {
	j := &util.JWT{
		[]byte("newtrekWang"),
	}
	claims := domain.CustomClaims{
		user.Id,
		user.UserName,
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