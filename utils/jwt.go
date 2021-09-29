package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	_ "github.com/wujunyi792/ginFrame/logger"
	"time"
)

type JWTClaims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

//声明过期时间
const TokenExpireDuration = time.Hour * 12

var MySecret = []byte("")

// GenToken 生成JWT
func GenToken() (string, error) {
	// 创建一个我们自己的声明
	c := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JWTClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
