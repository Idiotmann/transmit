package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		gin.SetMode(gin.DebugMode) //gin的模式 debug 生产 测试
		router := gin.Default()    //新建路由 ，创建生成引擎中间件log和recovery
		//根目录
		router.GET("/", func(c *gin.Context) {
			//request
			c.Writer.Write([]byte("asda"))
			//response
		})
		router.Run(":8080") //如果不开协程，会一直执行，去监听，而不会往下去执行
	}()
	log.Println("123")
	var ui lorca.UI
	ui, _ = lorca.New("http://127.0.0.1:8080", "", 1000, 600, "--disable--sync", "--disable-translate")
	//我们获取ctrl+c的中断，当中断时，关闭ui窗口  搜索golang 优雅退出相关的
	chSignal := make(chan os.Signal, 1) //只接受操作系统的信号 1 缓冲 通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	//syscall.SIGINT  syscall.SIGTERM 订阅信号终止和中断信号，发送给chsignal  https://www.jianshu.com/p/ae72ad58ecb6信号可以在这里面看
	//syscall.SIGINT就是ctrl+c   syscall.SIGTERM就是系统结束进程的信号对应的就是ctrl+c 和kill
	//select,等待第一个可以读或写的ch操作
	//select{}直接将main线程挂起，永远停 只占用很少的资源 其他方式退出，或者唤醒
	select {
	case <-ui.Done(): //等待ui结束的信号，这个线程就不动了
	case <-chSignal: //阻塞当前，等chSignal有信号，取出信号，ui关闭
	}
	//为什么写后面？ 因为前面两句都是做同样的事情。所以让它放后面执行就行
	ui.Close()

}
