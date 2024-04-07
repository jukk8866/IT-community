package router

import (
	"blue/controller"
	"blue/logger"
	"blue/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	r := gin.New()

	r.Use(middleware.Info, logger.GinLogger(), logger.GinRecovery(true))

	//r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"index.html": "index.html",
		})
	})

	v1 := r.Group("/api/v1")

	//注册
	v1.POST("/signup", controller.SignUpHandler)
	//登陆
	v1.POST("/login", controller.LoginHandler)

	//community
	//获取帖子分类
	v1.GET("/community", controller.CommunityHandler)

	//refresh 刷新token
	v1.POST("/re", controller.RefreshToken)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.CreatePostHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
