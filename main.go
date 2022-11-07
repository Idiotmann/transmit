package main

import (
	"github.com/Idiotmann/transmit/config"
	"github.com/Idiotmann/transmit/server"
	"github.com/zserge/lorca"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//如果不加go，server就会一直执行
	go func() {
		server.Run()
	}()
	startBrowser()

}
func startBrowser() {
	var ui lorca.UI
	ui, _ = lorca.New("http://127.0.0.1:"+config.GetPort()+"/static/index.html", "", 1000, 600, "--disable--sync", "--disable-translate")
	chSignal := make(chan os.Signal, 1) //只接受操作系统的信号 1 缓冲 通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ui.Done(): //等待ui结束的信号，这个线程就不动了 阻塞
	case <-chSignal: //阻塞当前，等chSignal有信号，取出信号，ui关闭
	}
	//为什么写后面？ 因为前面两句都是做同样的事情。所以让它放后面执行就行
	ui.Close()

}
