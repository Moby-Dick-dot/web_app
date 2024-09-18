package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	g "web_app/global"
	"web_app/response"
	"web_app/utils"
)

// JWTAuthMiddleware JWT鉴权中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseErrorWithMsg(c, g.CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.ResponseErrorWithMsg(c, g.CodeInvalidToken, "Token格式不对")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			fmt.Println(err)
			response.ResponseError(c, g.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(g.ContextUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("userID")来获取当前请求的用户信息

	} // func
}
