package routes

import (
	"net/http"
	"web_app/handler"
	"web_app/logger"
	"web_app/middleware"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 测试路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	v1 := r.Group("/api/v1")
	v1.POST("/login", GenHandlerFunc(handler.LoginHandler, ApiParamsErrCallback))
	v1.POST("/signup", GenHandlerFunc(handler.SignUpHandler, ApiParamsErrCallback))
	// 后续路由都需要鉴权
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", handler.CommunityHandler)
		v1.GET("/community/:id", handler.CommunityDetailHandler)
	}
	return r
}
