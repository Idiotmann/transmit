package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func QrcodesController(c *gin.Context) {
	//查询字符串
	if content := c.Query("content"); content != "" {
		//变成二维码切片
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}
		c.Data(http.StatusOK, "image/png", png) //为啥不用file 因为是用来展示的，不是用来下载的 c.File("image/png")是用来下载的图片
	} else {
		c.Status(http.StatusBadRequest)
	}
}
