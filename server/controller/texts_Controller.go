package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func TextsController(c *gin.Context) {
	var json struct {
		Raw string `json:"raw"` //小写不会被反序列话，因为是私有的，而暴露给外界的需要是小写的
	}
	//调用函数传一个空的结构题，就会把结构体填充 如果结果不为空，就是传的有问题
	//如果绑定成功就说明texe请求已经发给我了，执行下面的操作
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		exe, err := os.Executable() //获取当前执行文件的路径
		if err != nil {
			log.Fatal(err)
		}
		dir := filepath.Dir(exe) //获取当前执行文件的目录
		if err != nil {
			log.Fatal(err)
		}
		filename := uuid.New().String()          //创建一个文件名 随机
		uploads := filepath.Join(dir, "uploads") //拼接uploads的绝对路径
		err = os.MkdirAll(uploads, os.ModePerm)  //创建uploads目录    和 用户权限 直接写777不行，因为是八进制数字是511     目录权限 777 666 555 1 2 4
		if err != nil {
			log.Fatal(err)
		}

		fullPath := path.Join("uploads", filename+".txt") //拼接文件的绝对路径，（不含exe所在目录）
		//把内容写进去到目标路径中
		err = ioutil.WriteFile(filepath.Join(dir, fullPath), []byte(json.Raw), 0644) //将json.Raw写入文件
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"url": "/" + fullPath}) //返回文件的绝对路径（不含exe所在目录）

	}

}
