package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong!pong!")
	})

	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.POST("/post", controller.CreatePostHandler)
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.GET("/posts", controller.GetPostListHandler)
	v1.POST("/vote", controller.PostVoteHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
