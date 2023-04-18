package middlewares

import (
	"net/http"
	"sesi_12/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifytoken, err := helpers.VerifyToken(c)
		_ = verifytoken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauhtnticated",
				"massage": err.Error(),
			})
			return
		}
		c.Set("userData", verifytoken)
		c.Next()
	}
}
