package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Error(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: msgFlags[c],
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, rd)
}

func ErrorWithMsg(ctx *gin.Context, code MyCode, errMsg string) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, rd)
}

func Success(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}

	ctx.JSON(http.StatusOK, rd)
}
