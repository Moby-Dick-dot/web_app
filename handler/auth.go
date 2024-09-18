package handler

import (
	"web_app/dao/mysql"
	g "web_app/global"
	"web_app/model"
	"web_app/request"
	"web_app/response"
	"web_app/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// LoginHandler 登录
func LoginHandler(ctx *gin.Context, req *request.User) {
	// 调用数据库函数进行登录
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
	// 调用数据库函数进行注册
	err := mysql.Register(&model.User{
		UserName: req.UserName,
		Password: req.Password,
	})
	// 登录失败返回，登录失败信息
	if err != nil {
		response.ResponseError(ctx, g.CodeServerBusy)
		return
	}
	// 登录成功，返回登录成功的信息
	response.ResponseSuccess(ctx, nil)
}
