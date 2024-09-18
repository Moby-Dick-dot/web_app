package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	g "web_app/global"
	"web_app/response"
)

func CommunityHandler(c *gin.Context) {
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("mysql.GetCommunityList() failed", zap.Error(err))
		response.ResponseError(c, g.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	communityID := c.Param("id")
	communityList, err := mysql.GetCommunityByID(communityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
		response.ResponseErrorWithMsg(c, g.CodeSuccess, err.Error())
		return
	}
	response.ResponseSuccess(c, communityList)
}
