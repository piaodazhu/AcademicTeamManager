package middleware

import (
	"atm/common"
	"atm/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Result(response.ErrCodeNoLogin, nil, c)
			c.Abort()
			return
		}
		userid, err := common.VerifyToken(token)
		if userid != int64(uid) || err != nil {
			response.Result(response.ErrCodeTokenExpire, nil, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
