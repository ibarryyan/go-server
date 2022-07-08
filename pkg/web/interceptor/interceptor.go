package interceptor

import (
	"count_num/pkg/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var (
	// 排除的URL,暂不支持 /*
	Exclude = []string{
		"/user/login",
		"/user/register",
	}
)

// HttpInterceptor 自定义拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := range Exclude {
			if c.Request.URL.String() == Exclude[i] {
				//通过请求
				c.Next()
				return
			}
		}
		token := c.GetHeader("token")
		user := auth.GetToken(c, token)
		isPass := auth.CheckEnforce(cast.ToString(user.Role), c.Request.URL.String(), "")
		if isPass {
			c.Next()
			return
		}
		//定义错误,终止并返回该JSON 不能使用 c.JSON(403,"no auth")
		c.AbortWithStatusJSON(403, "no auth")
	}
}
