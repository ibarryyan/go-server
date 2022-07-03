package interceptor

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	// 排除的URL,暂不支持 /*
	Exclude = []string{
		"/user/login",
	}
)

// HttpInterceptor 自定义拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 example 变量
		// c.Set("example", "12345")
		// 请求前
		log.Print("--------------拦截器-------------")
		for i := range Exclude {
			if c.Request.URL.String() == Exclude[i] {
				//通过请求
				c.Next()
				return
			}
		}
		//定义错误,终止并返回该JSON 不能使用 c.JSON(403,"no auth")
		c.AbortWithStatusJSON(403, "no auth")
	}
}
