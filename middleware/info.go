package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Info(c *gin.Context) {
	//fmt.Println("分组路由 中间件")
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
}
