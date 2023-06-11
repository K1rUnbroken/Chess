package service

import (
	"chess/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var myJwtSecret = []byte("kyr666")

// GetToken 生成JWT
func GetToken(c *model.MyClaims) (string, error) {
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	aTokenSigned, err := aToken.SignedString(myJwtSecret)
	if err != nil {
		return "", err
	}

	// 返回
	return aTokenSigned, nil
}

// ParseToken 解析JWT
func ParseToken(tokenStr string) (*model.MyClaims, error) {
	// 解析tokenStr字符串，得到JWT
	JWT, err := jwt.ParseWithClaims(tokenStr, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return myJwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 获取JWT中的自定义MyClaims部分，同时校验JWT是否有效
	myClaims, aOK := JWT.Claims.(*model.MyClaims)
	if !(aOK && JWT.Valid) {
		return nil, errors.New("invalid JWT")
	}

	return myClaims, nil
}
