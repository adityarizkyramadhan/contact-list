package utils

import "github.com/gin-gonic/gin"

func ResponseFail(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, gin.H{
		"status":  "fail",
		"message": message,
		"error":   err.Error(),
	})
}

func ResponseSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}
