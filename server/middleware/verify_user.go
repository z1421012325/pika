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
		c.Request.Header.Set(config.SET_USER_ID_NAME, Uid)
		c.Next()
	}
}

func respondWithError(code int, msg string, c *gin.Context) {
	c.JSON(code, err.ParamBindErrorResult(msg))
	c.Abort()
}
