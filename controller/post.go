package controller

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddPostHandler(c *gin.Context) {
	var post models.PostReq
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, response.CodeInvalidParams)
		zap.L().Error(err.Error())
		return
	}

	err := mysql.AddPost(post)

	if err != nil {
		response.Error(c, response.CodeServerBusy)
		zap.L().Error(err.Error())
		return
	}

	response.Success(c, post)
}

func GetPostByIdHandler(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		response.Error(c, response.CodeInvalidParams)
		zap.L().Error(err.Error())
		return
	}
	post, err := mysql.GetPostById(id)

	if err != nil {
		response.Error(c, response.CodeServerBusy)
		zap.L().Error(err.Error())
		return
	}

	response.Success(c, post)
}

func GetPostListByIdsHandler(c *gin.Context) {
	var postIds models.PostIds

	if err := c.ShouldBindJSON(&postIds); err != nil {
		response.Error(c, response.CodeInvalidParams)
		zap.L().Error(err.Error())
		return
	}

	posts, err := mysql.GetPostListByIds(postIds.Ids)
	if err != nil {
		response.Error(c, response.CodeServerBusy)
		zap.L().Error(err.Error())
		return
	}

	response.Success(c, posts)
}

func GetPostListHandler(c *gin.Context) {
	var post models.PostReq

	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, response.CodeInvalidParams)
		zap.L().Error(err.Error())
		return
	}

	posts, err := mysql.GetPostList(post)
	if err != nil {
		response.Error(c, response.CodeServerBusy)
		zap.L().Error(err.Error())
		return
	}

	response.Success(c, posts)
}
