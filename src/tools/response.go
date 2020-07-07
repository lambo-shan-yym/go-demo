package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func response(ctx *gin.Context, statusCode int, code int, msg string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code": code,
		"msg":  msg,
		"result": data,
	})
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	response(ctx, http.StatusOK, 0, "", data)
}

func FailResponse(ctx *gin.Context, code int, msg string, data interface{}) {
	response(ctx, http.StatusOK, code, msg, data)
}
