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

func ErrorResponseHandler(ctx *gin.Context, code int, err interface{}) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"status": "failed",
		"error":  err,
	})
}

func ErrorResponseValidationHandler(ctx *gin.Context, code int, err interface{}, val interface{}) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"status":     "failed",
		"error":      err,
		"validation": val,
	})
}
