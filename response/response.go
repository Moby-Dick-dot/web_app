package response

import (
	"net/http"
	g "web_app/global"

	"github.com/gin-gonic/gin"
)

// 响应的数据结构
type ResponseData struct {
	Code    g.MyCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 响应返回错误
func ResponseError(ctx *gin.Context, c g.MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

// 返回错误，并注明错误信息
func ResponseErrorWithMsg(ctx *gin.Context, code g.MyCode, errMsg string) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

// 响应返回成功信息
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    g.CodeSuccess,
		Message: g.CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
