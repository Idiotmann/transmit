package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

//go:embed frontend/dist/*
//在打包go的可执行文件时，把这个目录下的所有文件打包进去
var FS embed.FS

func main() {
	go func() {
		gin.SetMode(gin.DebugMode) //gin的模式 debug 生产 测试
		router := gin.Default()    //新建路由 ，创建生成引擎中间件log和recovery
		//根目录 不需要了
		//router.GET("/", func(c *gin.Context) {
		//	//request
		//	c.Writer.Write([]byte("asda"))
		//	//response
		//})
		//把所有静态文件变成一个变量  获得一个子文件系统 子文件系统的根由第二个参数 dir 指定
		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		//把后面读取到的静态文件放在/static/下面
		router.StaticFS("/static", http.FS(staticFiles))
		//如果没有该地址
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
		router.Run(":8080") //如果不开协程，会一直执行，去监听，而不会往下去执行
	}()
	log.Println("123")
	var ui lorca.UI
	ui, _ = lorca.New("http://127.0.0.1:8080/static/index.html", "", 1000, 600, "--disable--sync", "--disable-translate")
	//我们获取ctrl+c的中断，当中断时，关闭ui窗口  搜索golang 优雅退出相关的
	chSignal := make(chan os.Signal, 1) //只接受操作系统的信号 1 缓冲 通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	//syscall.SIGINT  syscall.SIGTERM 订阅信号终止和中断信号，发送给chsignal  https://www.jianshu.com/p/ae72ad58ecb6信号可以在这里面看
	//syscall.SIGINT就是ctrl+c   syscall.SIGTERM就是系统结束进程的信号对应的就是ctrl+c 和kill
	//select,等待第一个可以读或写的ch操作
	//select{}直接将main线程挂起，永远停 只占用很少的资源 其他方式退出，或者唤醒
	select {
	case <-ui.Done(): //等待ui结束的信号，这个线程就不动了 阻塞
	case <-chSignal: //阻塞当前，等chSignal有信号，取出信号，ui关闭
	}
	//为什么写后面？ 因为前面两句都是做同样的事情。所以让它放后面执行就行
	ui.Close()

}
