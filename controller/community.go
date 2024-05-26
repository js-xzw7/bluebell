package controller

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/response"

	"github.com/gin-gonic/gin"
)

func AddCommunityHandler(c *gin.Context) {
	var cfo models.CreateCommunityRequest
	if err := c.ShouldBindJSON(&cfo); err != nil {
		response.ErrorWithMsg(c, response.CodeInvalidParams, err.Error())
		return
	}

	err := mysql.AddCommunity(&cfo)

	if err != nil {
		response.ErrorWithMsg(c, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, nil)
}

func GetCommunityListHandler(c *gin.Context) {
	var community models.Community
	if err := c.ShouldBindJSON(&community); err != nil {
		response.ErrorWithMsg(c, response.CodeInvalidParams, err.Error())
		return
	}
	communityList, err := mysql.GetCommunityList(community)
	if err != nil {
		response.ErrorWithMsg(c, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, communityList)
}
