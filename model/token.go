package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "efdegv"

type JWTCustomClaims struct {
	jwt.StandardClaims
	//追加自己需要的信息
	Uid  string `json:"uid"`  //phone
	UDID string `json:"udid"` //设备类型
}

/*
1. 生成token
SecretKey是一个const常量
*/
func CreateToken(SecretKey []byte, issuer string, Uid string, Udid string) (tokenString string, err error) {
	claims := &JWTCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 360).Unix()),
			Issuer:    issuer,
		},
		Uid,
		Udid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Println(token)
	tokenString, err = token.SignedString(SecretKey)
	//fmt.Println(err)
	return
}

//解析token
func ParseToken(tokenStr string, SecretKey []byte) (jwt.Claims, error) {
	var token *jwt.Token
	token, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	} else {
		if token.Valid {
			claims := token.Claims
			return claims, nil
		}
	}
	return nil, nil
}
