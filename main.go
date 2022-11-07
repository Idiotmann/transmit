// Package main
// @Description:
package main

import (
	"github.com/Idiotmann/transmit/config"
	"github.com/Idiotmann/transmit/server"
	"os"
	"os/exec"
	"os/signal"
)

//var ui lorca.UI
//ui, _ = lorca.New("http://127.0.0.1:8080/static/index.html", "", 1000, 600, "--disable--sync", "--disable-translate")
////我们获取ctrl+c的中断，当中断时，关闭ui窗口  搜索golang 优雅退出相关的
//chSignal := make(chan os.Signal, 1) //只接受操作系统的信号 1 缓冲 通道
//signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
////syscall.SIGINT  syscall.SIGTERM 订阅信号终止和中断信号，发送给chsignal  https://www.jianshu.com/p/ae72ad58ecb6信号可以在这里面看
////syscall.SIGINT就是ctrl+c   syscall.SIGTERM就是系统结束进程的信号对应的就是ctrl+c 和kill
////select,等待第一个可以读或写的ch操作
////select{}直接将main线程挂起，永远停 只占用很少的资源 其他方式退出，或者唤醒
//select {
//case <-ui.Done(): //等待ui结束的信号，这个线程就不动了 阻塞
//case <-chSignal: //阻塞当前，等chSignal有信号，取出信号，ui关闭
//}
////为什么写后面？ 因为前面两句都是做同样的事情。所以让它放后面执行就行
//ui.Close()

func main() {
	//如果不加go，server就会一直执行
	go func() {
		server.Run()
	}()
	cmd := startBrowser()
	ch := listenToInterrupt()
	//等待终端信号
	select {
	case <-ch:
		cmd.Process.Kill()
	}

}
func startBrowser() *exec.Cmd {
	//var ui lorca.UI
	//ui, _ = lorca.New("http://127.0.0.1:8080/static/index.html", "", 1000, 600, "--disable--sync", "--disable-translate")
	////我们获取ctrl+c的中断，当中断时，关闭ui窗口  搜索golang 优雅退出相关的
	//chSignal := make(chan os.Signal, 1) //只接受操作系统的信号 1 缓冲 通道
	//signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	////syscall.SIGINT  syscall.SIGTERM 订阅信号终止和中断信号，发送给chsignal  https://www.jianshu.com/p/ae72ad58ecb6信号可以在这里面看
	////syscall.SIGINT就是ctrl+c   syscall.SIGTERM就是系统结束进程的信号对应的就是ctrl+c 和kill
	////select,等待第一个可以读或写的ch操作
	////select{}直接将main线程挂起，永远停 只占用很少的资源 其他方式退出，或者唤醒
	//select {
	//case <-ui.Done(): //等待ui结束的信号，这个线程就不动了 阻塞
	//case <-chSignal: //阻塞当前，等chSignal有信号，取出信号，ui关闭
	//}
	////为什么写后面？ 因为前面两句都是做同样的事情。所以让它放后面执行就行
	//ui.Close()
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+config.GetPort()+"/static/index.html")
	cmd.Start()
	return cmd
}

func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
