package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if context.Request.Method == "OPTIONS" {
		context.AbortWithStatus(204)
		return
	}

	context.Next()
}

/**
* Handle requests errors
 */
func ErrorHandler(context *gin.Context) {
	context.Next()

	// TODO: Handle it in a better way
	if len(context.Errors) > 0 {
		context.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": context.Errors,
		})
	}
}
