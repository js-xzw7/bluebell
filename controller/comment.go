package controller

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddCommentHandler(c *gin.Context) {
	var comment models.CommentInfo

	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, response.CodeInvalidParams)
		zap.L().Error(err.Error())
		return
	}

	err := mysql.AddCommont(comment)

	if err != nil {
		response.Error(c, response.CodeServerBusy)
		zap.L().Error(err.Error())
		return
	}

	response.Success(c, nil)
}

func GetCommentListByIdsHandler(c *gin.Context) {
	var commentIds models.CommentIds

	if err := c.ShouldBindJSON(&commentIds); err != nil {
		response.Error(c, response.CodeInvalidParams)
		zap.L().Error(err.Error())
		return
	}

	comments, err := mysql.GetCommentListByIds(commentIds.Ids)

	if err != nil {
		response.Error(c, response.CodeServerBusy)
		zap.L().Error(err.Error())
		return
	}

	response.Success(c, comments)
}

func GetCommentListHandler(c *gin.Context) {
	var comment models.CommentInfo

	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, response.CodeInvalidParams)
		zap.L().Error(err.Error())
		return
	}

	comments, err := mysql.GetCommentList(comment)

	if err != nil {
		response.Error(c, response.CodeServerBusy)
		zap.L().Error(err.Error())
		return
	}

	response.Success(c, comments)
}
