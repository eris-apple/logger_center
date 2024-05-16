package utils

import (
	"github.com/gin-gonic/gin"
)

func ResponseHandler(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func ErrorResponseHandler(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"status": "failed",
		"error":  err.Error(),
	})
}
