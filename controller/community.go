package controller

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/response"
	"strconv"

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
	var community models.CommunityInfo
	if err := c.ShouldBind(&community); err != nil {
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

func GetCommunityByIdHandler(c *gin.Context) {
	communityId := c.Query("id")

	if communityId == "" {
		response.ErrorWithMsg(c, response.CodeInvalidParams, "id is empty")
		return
	}
	id, err := strconv.ParseUint(communityId, 10, 64)
	if err != nil {
		response.ErrorWithMsg(c, response.CodeInvalidParams, "id is not a number")
		return
	}

	community, err := mysql.GetCommunityById(id)
	if err != nil {
		response.ErrorWithMsg(c, response.CodeServerBusy, err.Error())
		return
	}

	response.Success(c, community)
}
