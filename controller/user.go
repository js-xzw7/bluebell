package controller

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/response"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1.获取请求参数
	//2.校验参数有效性
	var refo models.RegisterForm
	if err := c.ShouldBindJSON(&refo); err != nil {
		response.ErrorWithMsg(c, response.CodeInvalidParams, err.Error())
		return
	}

	//3.注册用户
	err := mysql.Register(&models.User{
		UserName: refo.UserName,
		Password: refo.Password,
	})

	if errors.Is(err, mysql.ErrorUserExit) {
		response.Error(c, response.CodeUserNotExist)
		return
	}

	if err != nil {
		zap.L().Error("mysql.register(&u) failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, nil)
}

func LoginHandler(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		return
	}

	if err := mysql.Login(&u); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		response.Error(c, response.CodeInvalidPassword)
		return
	}

	//生成token
	aToken, rToken, _ := jwt.GenToken(u.UserId)
	response.Success(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       u.UserId,
		"username":     u.UserName,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refreshToken")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		response.ErrorWithMsg(c, response.CodeInvalidToken, "缺少请求头auth token")
		c.Abort()
		return
	}

	//按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.ErrorWithMsg(c, response.CodeInvalidToken, "token 格式不对")
		c.Abort()
		return
	}

	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)

	if err != nil {
		response.ErrorWithMsg(c, response.CodeInvalidToken, "token 解析失败")
		c.Abort()
		return
	}

	response.Success(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})

}
