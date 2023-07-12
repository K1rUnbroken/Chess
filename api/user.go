package api

import (
	"chess/dao"
	"chess/model"
	"chess/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" && password == "" {
		username = c.Query("username")
		password = c.Query("password")
	}

	if username == "" && password == "" {
		v1, _ := c.Get("username")
		v2, _ := c.Get("password")
		username = fmt.Sprintf("%s", v1)
		password = fmt.Sprintf("%s", v2)
	}

	// 校验账号信息
	value, ok := dao.UserDB[username]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户不存在",
		})
		return
	}
	if value != password {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "密码错误",
		})
		return
	}

	// 颁发token
	claims := model.MyClaims{Username: username}
	jwt, err := service.GetToken(&claims)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"jwt":  jwt,
	})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" && password == "" {
		username = c.Query("username")
		password = c.Query("password")
	}

	if username == "" && password == "" {
		v1, _ := c.Get("username")
		v2, _ := c.Get("password")
		username = fmt.Sprintf("%s", v1)
		password = fmt.Sprintf("%s", v2)
	}

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户名或密码不能为空",
		})
		return
	}

	// 用户名已经存在
	if _, ok := dao.UserDB[username]; ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户名以存在",
		})
		return
	}

	// 插入数据库
	dao.UserDB[username] = password

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}
