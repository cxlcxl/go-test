package my_jwt

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserId   int64  `json:"user_id"`
	Name     string `json:"user_name"`
	Email    string `json:"email"`
	RealName string `json:"real_name"`
	Avatar   string `json:"avatar"`
	jwt.StandardClaims
}
