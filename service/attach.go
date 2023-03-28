package service

import (
	"Gin_WebSocket_IM/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Upload
// @Summary 发送图片
// @Tags 信息模块
// @Router /attach/upload [post]
func Upload(c *gin.Context) {
	w := c.Writer
	r := c.Request
	srcFile, head, err := r.FormFile("file")
	if err != nil {
		utils.RespFail(w, err.Error())
		return
	}
	suffix := ".png"
	ofilName := head.Filename
	tem := strings.Split(ofilName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstFile, err := os.Create("./asset/upload/" + fileName)
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	_, err = io.Copy(dstFile, srcFile)

	url := "./asset/upload/" + fileName
	utils.RespOK(w, url, "发送图片成功")
}
