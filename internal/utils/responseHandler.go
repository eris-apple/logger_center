package utils

import (
	"github.com/gin-gonic/gin"
)

func ResponseHandler(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func ErrorResponseHandler(ctx *gin.Context, code int, message string, error interface{}) {
	ctx.JSON(code, gin.H{
		"status":  "failed",
		"message": message,
		"error":   error,
	})
}
