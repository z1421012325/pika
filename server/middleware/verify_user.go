package middleware

import (
	"fmt"

	"pika/config"
	err "pika/server/code"
	"pika/tools"

	"github.com/gin-gonic/gin"
)

// 验证用户登录中间件

func VerifyUserLogin() gin.HandlerFunc {

	fmt.Println(" Verify User Login Middleware ...")

	return func(c *gin.Context) {

		StrToken := tools.GetRequestToken(c, config.SET_TOKEN_NAME)
		Uid, ok := tools.CheckToken(StrToken)
		if !ok {
			// 用户未登录
			respondWithError(201, "API token required", c)
			return
		}
		// 如果通过验证,在header中设置uid,返回进行删除
		c.Request.Header.Set(config.SET_USER_ID_NAME, Uid)
		c.Next()
		// todo 测试是否在返回前被删除
		c.Request.Header.Del(config.SET_TOKEN_NAME)
	}
}

func respondWithError(code int, msg string, c *gin.Context) {
	c.JSON(code, err.ParamBindErrorResult())
	c.Abort()
}
