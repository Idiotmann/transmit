package server

import (
	"embed"
	"github.com/Idiotmann/transmit/config"
	"github.com/Idiotmann/transmit/server/controller"
	"github.com/Idiotmann/transmit/server/ws"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dist/*
//将整个目录下的文件嵌入到exe中，就不用读其他的文件，不用安装环境
//在打包go的可执行文件时，把这个目录下的所有文件打包进去
var FS embed.FS

//启动gin服务
func Run() {
	gin.SetMode(gin.DebugMode) //gin的模式 debug 生产 测试
	router := gin.Default()    //新建路由 ，创建生成引擎中间件log和recovery
	//把所有静态文件变成一个变量  获得一个子文件系统 子文件系统的根由第二个参数 dir 指定
	hub := ws.NewHub()
	go hub.Run()
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	router.GET("uploads/:path", controller.UploadsController)
	router.GET("/api/v1/addresses", controller.AddressesController)
	//获取接口
	router.GET("/api/v1/qrcodes", controller.QrcodesController)
	router.POST("/api/v1/files", controller.FilesController)
	//下载文件
	router.POST("/api/v1/texts", controller.TextsController)
	router.GET("/ws", func(c *gin.Context) {
		ws.HttpController(c, hub)
	})

	//把后面读取到的静态文件放在/static/下面
	router.StaticFS("/static", http.FS(staticFiles))
	//如果没有该地址  最后的路由
	router.NoRoute(func(c *gin.Context) {
		//获取用户访问的路径
		path := c.Request.URL.Path
		//如果路径是以static开头的,想访问静态文件
		if strings.HasPrefix(path, "/static/") {
			//去读文件  fs.FS接口只有一个方法，即打开一个命名文件   Open 方法返回一个 fs.File 接口类型
			//三种方法Read 方法进行读操作   Stat 方法判断文件是什么类型，也可能需要获取文件的一些元数据信息返回一个 FileInfo类型，它也是一个接口
			//Go 中，你应该始终记住，打开文件，进行操作后，记得关闭文件，否则会泄露文件描述符。所以，fs.File 的第是三个方法就是 Close 方法，它的签名和 io.Closer 是一致的。
			//type File interface {
			//	Stat() (FileInfo, error)
			//	Read([]byte) (int, error)
			//	Close() error
			//}
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			//statistics 统计   返回的接口类型
			//type FileInfo interface {
			//	Name() string       // base name of the file
			//	Size() int64        // length in bytes for regular files; system-dependent for others
			//	Mode() FileMode     // file mode bits
			//	ModTime() time.Time // modification time
			//	IsDir() bool        // abbreviation for Mode().IsDir()
			//	Sys() any           // underlying data source (can return nil)
			//}
			if err != nil {
				log.Fatal(err)
			}
			//content-type 解析https://www.runoob.com/http/http-content-type.html
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)
		} else {
			//否则想访问动态文件
			//现在没有就404 里面直接404 就行
			c.Status(http.StatusNotFound)
		}

	})
	router.Run(":" + config.GetPort()) //如果不开协程，会一直执行，去监听，而不会往下去执行
}
