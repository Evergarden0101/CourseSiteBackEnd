package domain

import "github.com/dgrijalva/jwt-go"

type LoginResult struct {
	Token string `json:"token"`
	User
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Id string `json:"id"`
	UserName string  `json:"userName"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	UserType string  `json:"userType"`
	jwt.StandardClaims
}