package jwt

import (
	"github.com/aircjm/gocard/common/responseStatus"
	"github.com/aircjm/gocard/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = responseStatus.SUCCESS
		token := c.GetHeader("X-Token")
		if token == "" {
			code = responseStatus.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = responseStatus.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = responseStatus.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != responseStatus.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  responseStatus.GetStatusMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
