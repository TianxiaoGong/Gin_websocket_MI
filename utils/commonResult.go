package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0 //操作成功
	FAILED  int = 1 //操作失败
)

// Success 普通成功返回
func Success(ctx *gin.Context, msg string, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  msg,
		"data": v,
	})
}

// Failed 操作失败返回
func Failed(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  msg,
	})
}
