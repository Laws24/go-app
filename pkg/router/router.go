package router

import (
	"go-app/pkg/handler"

	"github.com/BoyYangZai/go-server-lib/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	// Route group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", handler.Login)
		v1.POST("/submit", handler.Submit)
		v1.POST("/read", handler.Read)
	}

	// Route group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", handler.Login)
		v2.POST("/submit", handler.Submit)
		v2.POST("/read", handler.Read)
	}

	user := router.Group("/user")
	{
		user.POST("/verify-code", handler.VerifyCode)
		user.POST("/registry", handler.Registry)
		user.POST("/login", handler.Login)
		user.GET("/view-info", jwt.AuthMiddleware(), handler.ViewInfo)
		user.POST("/modify-info", handler.ModifyInfo)
		user.POST("/changeAvatar", handler.ChangeAvatar)
	}

	auth_test := router.Group("/auth-test")
	{
		auth_test.Use(jwt.AuthMiddleware())
		auth_test.GET("/", handler.Submit)
	}

	router.Run(":8080")

	return router
}
