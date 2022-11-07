package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadsController(c *gin.Context) {
	//获取路由里面的path
	if path := c.Param("path"); path != "" {
		target := filepath.Join(GetUploadsDir(), path)
		//服务器通过header告诉浏览器，回送数据的信息
		c.Header("Content-Description", "File Transfer") //
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Type", "application/octet-stream") //回送数据的类型
		c.File(target)                                       //给前端发送任意类型的文件文件 写入
	} else {
		c.Status(http.StatusNotFound)
	}
}
func GetUploadsDir() (uploads string) {
	exe, err := os.Executable() //返回当前运行程序的路径 而os.Getwd()会输出实际的工作目录
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe) // filepath.Dir()函数用于返回指定路径中除最后一个元素以外的所有元素。filepath.Dir("/Geeks/GFG/gfg.org") /Geeks/GFG
	uploads = filepath.Join(dir, "uploads")
	return
}
