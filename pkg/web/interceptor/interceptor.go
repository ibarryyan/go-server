package interceptor

import (
	"github.com/gin-gonic/gin"
	"log"
)

// HttpInterceptor 自定义拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前
		log.Print("--------------拦截器-------------")
		//定义错误,终止并返回该JSON
		//c.AbortWithStatusJSON(500, "error")
		//requestURI := c.Request.RequestURI
		//fmt.Println(requestURI)
		//通过请求
		c.Next()
	}
}
