package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	g "web_app/global"
)

type ResponseData struct {
	Code    g.MyCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, c g.MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code g.MyCode, errMsg string) {
	rd := &ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    g.CodeSuccess,
		Message: g.CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
