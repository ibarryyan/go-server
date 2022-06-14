package interceptor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// HttpInterceptor 自定义中间件
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前
		log.Print("--------------拦截器-------------")
		//定义错误,终止并返回该JSON
		//c.AbortWithStatusJSON(500, "error")
		requestURI := c.Request.RequestURI
		fmt.Println(requestURI)
		//通过请求
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}
