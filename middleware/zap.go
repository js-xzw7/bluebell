package middleware

import (
	"bytes"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Zap() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()

		//处理请求
		c.Next()

		//计算响应时间
		end := time.Now()
		latency := end.Sub(startTime)

		//获取请求信息
		path := c.Request.URL.Path
		clientIp := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		//拼接日志输出格式
		var buffer bytes.Buffer
		buffer.WriteString("[")
		buffer.WriteString(clientIp)
		buffer.WriteString("] ")
		buffer.WriteString(method)
		buffer.WriteString(" ")
		buffer.WriteString(path)
		buffer.WriteString(" ")
		buffer.WriteString(strconv.Itoa(statusCode))
		buffer.WriteString(" ")
		buffer.WriteString(latency.String())
		msg := buffer.String()

		//使用zap记录请求信息
		// zap.L().Info(fmt.Sprintf("[%s] %s %s %d %s", clientIp, method, path, statusCode, latency))
		zap.L().Info(msg)
	}
}
