package middlewares

import (
	"gobackend/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrInvalidToken, "error token")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
