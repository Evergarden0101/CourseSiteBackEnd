package util

import (
	"awesomeProject/main/constant"
	"awesomeProject/main/dao"
	"awesomeProject/main/domain"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

func GetUser(c *gin.Context)(string){
	claim,_ := c.Get("claims")
	user := claim.(*domain.CustomClaims)
	return user.Id
}

func TeacherAuth(c *gin.Context)bool{
	userId := GetUser(c)
	user := dao.GetUserById(userId)
	if user.UserType != constant.TEACHER{
		log.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": "",
		})
		return false
	}
	return true
}

func AdminAuth(c *gin.Context)bool{
	userId := GetUser(c)
	user := dao.GetUserById(userId)
	if user.UserType != constant.ADMIN{
		log.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": "",
		})
		return false
	}
	return true
}

func TeacherCourseAuth(c *gin.Context,courseId string)bool{
	userId := GetUser(c)
	user := dao.GetUserById(userId)
	course := dao.GetCourse(courseId)
	scr := dao.GetSCR(courseId,userId)
	if user.Id != course.TeacherId && scr.Type != constant.ASS{
		log.Println(user)
		log.Println(course)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": "",
		})
		return false
	}
	return true
}

func StudentCourseAuth(c *gin.Context,courseId string)bool{
	userId := GetUser(c)
	if !dao.GetSCRById(courseId,userId){
		log.Println(userId)
		log.Println(courseId)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.DENIED,
			"msg":  "没有权限",
			"data": "",
		})
		return false
	}
	return true
}

func BindData(c *gin.Context,obj interface{})bool{
	err :=c.ShouldBindJSON(obj)
	if err !=nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
		})
		return false
	}
	log.Println(obj)
	return true
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims domain.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*domain.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*domain.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*domain.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}