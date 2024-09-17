package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	g "web_app/global"
	"web_app/model"
	"web_app/request"
	"web_app/response"
	"web_app/utils"
)

// LoginHandler 登录
func LoginHandler(ctx *gin.Context, req *request.User) {
	if err := mysql.Login(req); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		response.ResponseError(ctx, g.CodeInvalidPassword)
		return
	}
	// 生成token并返回
	aToken, rToken, _ := utils.GenToken(req.UserID)
	response.ResponseSuccess(ctx, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       req.UserID,
		"username":     req.UserName,
	})
}

// SignUpHandler 注册用户
func SignUpHandler(ctx *gin.Context, req *request.RegisterReq) {
	err := mysql.Register(&model.User{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		response.ResponseError(ctx, g.CodeServerBusy)
		return
	}
	response.ResponseSuccess(ctx, nil)
}
