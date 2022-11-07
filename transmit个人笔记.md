



#                                        同步传输系统



[TOC]



## 前言

项目参考 博客

[(53条消息) 【GO语言】实现同步传输系统：局域网内手机和电脑互传文件互发消息。go语言练手项目_立志考博士的博客-CSDN博客_手机 局域网给电脑 发消息 server-client](https://blog.csdn.net/lwj8819/article/details/123729385?spm=1001.2014.3001.5502)

项目地址[FrankFang/synk (github.com)](https://github.com/FrankFang/synk)

参考项目地址[2394799692/transmit-doc: go语言练习项目，局域网传输文件 (github.com)](https://github.com/2394799692/transmit-doc)

参考项目地址https://github.com/2394799692/transmit-doc

第 6 ~ 10 节课的代码：https://github.com/FrankFang/synk-demo-1 

第 11 ~ 13 节课的代码：https://github.com/FrankFang/synk-demo-2 

第 14 ~ 16 节课：https://github.com/FrankFang/synk-demo-3 

第 17 ~ 21 节课：https://github.com/FrankFang/synk-demo-4 

第 22 ~ 25 节课：https://github.com/FrankFang/synk-demo-5 

第 26 ~ 27 节课：https://github.com/FrankFang/synk-demo-6

## 创建流程

### 一  创建仓库

我的项目地址[Idiotmann/transmit (github.com)](https://github.com/Idiotmann/transmit)

创建步骤

创建新的仓库transmit 公共为公共

步骤

打开链接[GitHub](https://github.com/)、

创建仓库  第一次流程

```sh
echo "# transmit" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main  //修改分支名字为main
git remote add origin https://github.com/Idiotmann/ttl.git //会出现问题
git push -u origin main //推送到origin主机的main 分支
```

![image-20221026210745229](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221026210745229.png)

出现问题后：解决[(53条消息) 解决报错：Failed to connect to github.com port 443 after 21098 ms: Timed out_前端小学生 (●—●)的博客-CSDN博客](https://blog.csdn.net/weixin_52796927/article/details/121664227)

```shell
git config --global https.proxy
git config --global --unset https.proxy
```

==已经存在的==

```shell
//已经存在的
git init
git add .
//自己起的名字
git commit -m "XXX"
git push -u origin xxx(分支名称)
//如果上面超时，就试一下下面的 如果还不行，就多试几次
git config --global https.proxy
git config --global --unset https.proxy
```

要建立两个分支

```
git branch //查看分支情况
git branch test//创建test分支
git branch -m xxx //修改分支名字
git checkout  xxx //更改分支

```



![image-20221103212206715](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221103212206715.png)

```shell
git init                                                  # 初始化本地git仓库（创建新仓库）
git config --global user.name "xxx"                       # 配置用户名
git config --global user.email "xxx@xxx.com"              # 配置邮件
git config --global color.ui true                         # git status等命令自动着色
git config --global color.status auto
git config --global color.diff auto
git config --global color.branch auto
git config --global color.interactive auto
git config --global --unset http.proxy                    # remove  proxy configuration on git
git clone git+ssh://git@192.168.53.168/VT.git             # clone远程仓库
git status                                                # 查看当前版本状态（是否修改）
git add xyz                                               # 添加xyz文件至index
git add .                                                 # 增加当前子目录下所有更改过的文件至index
git commit -m 'xxx'                                       # 提交
git commit --amend -m 'xxx'                               # 合并上一次提交（用于反复修改）
git commit -am 'xxx'                                      # 将add和commit合为一步
git rm xxx                                                # 删除index中的文件
git rm -r *                                               # 递归删除
git log                                                   # 显示提交日志
git log -1                                                # 显示1行日志 -n为n行
git log -5
git log --stat                                            # 显示提交日志及相关变动文件
git log -p -m
git show dfb02e6e4f2f7b573337763e5c0013802e392818         # 显示某个提交的详细内容
git show dfb02                                            # 可只用commitid的前几位
git show HEAD                                             # 显示HEAD提交日志
git show HEAD^                                            # 显示HEAD的父（上一个版本）的提交日志 ^^为上两个版本 ^5为上5个版本
git tag                                                   # 显示已存在的tag
git tag -a v2.0 -m 'xxx'                                  # 增加v2.0的tag
git show v2.0                                             # 显示v2.0的日志及详细内容
git log v2.0                                              # 显示v2.0的日志
git diff                                                  # 显示所有未添加至index的变更
git diff --cached                                         # 显示所有已添加index但还未commit的变更
git diff HEAD^                                            # 比较与上一个版本的差异
git diff HEAD -- ./lib                                    # 比较与HEAD版本lib目录的差异
git diff origin/master..master                            # 比较远程分支master上有本地分支master上没有的
git diff origin/master..master --stat                     # 只显示差异的文件，不显示具体内容
git remote add origin git+ssh://git@192.168.53.168/VT.git # 增加远程定义（用于push/pull/fetch）
git branch                                                # 显示本地分支
git branch --contains 50089                               # 显示包含提交50089的分支
git branch -a                                             # 显示所有分支
git branch -r                                             # 显示所有原创分支
git branch --merged                                       # 显示所有已合并到当前分支的分支
git branch --no-merged                                    # 显示所有未合并到当前分支的分支
git branch -m master master_copy                          # 本地分支改名
git checkout -b master_copy                               # 从当前分支创建新分支master_copy并检出
git checkout -b master master_copy                        # 上面的完整版
git checkout features/performance                         # 检出已存在的features/performance分支
git checkout --track hotfixes/BJVEP933                    # 检出远程分支hotfixes/BJVEP933并创建本地跟踪分支
git checkout v2.0                                         # 检出版本v2.0
git checkout -b devel origin/develop                      # 从远程分支develop创建新本地分支devel并检出
git checkout -- README                                    # 检出head版本的README文件（可用于修改错误回退）
git merge origin/master                                   # 合并远程master分支至当前分支
git cherry-pick ff44785404a8e                             # 合并提交ff44785404a8e的修改
git push origin master                                    # 将当前分支push到远程master分支
git push origin :hotfixes/BJVEP933                        # 删除远程仓库的hotfixes/BJVEP933分支
git push --tags                                           # 把所有tag推送到远程仓库
git fetch                                                 # 获取所有远程分支（不更新本地分支，另需merge）
git fetch --prune                                         # 获取所有原创分支并清除服务器上已删掉的分支
git pull origin master                                    # 获取远程分支master并merge到当前分支
git mv README README2                                     # 重命名文件README为README2
git reset --hard HEAD                                     # 将当前版本重置为HEAD（通常用于merge失败回退）
git rebase
git branch -d hotfixes/BJVEP933                           # 删除分支hotfixes/BJVEP933（本分支修改已合并到其他分支）
git branch -D hotfixes/BJVEP933                           # 强制删除分支hotfixes/BJVEP933
git ls-files                                              # 列出git index包含的文件
git show-branch                                           # 图示当前分支历史
git show-branch --all                                     # 图示所有分支历史
git whatchanged                                           # 显示提交历史对应的文件修改
git revert dfb02e6e4f2f7b573337763e5c0013802e392818       # 撤销提交dfb02e6e4f2f7b573337763e5c0013802e392818
git ls-tree HEAD                                          # 内部命令：显示某个git对象
git rev-parse v2.0                                        # 内部命令：显示某个ref对于的SHA1 HASH
git reflog                                                # 显示所有提交，包括孤立节点
git show HEAD@{5}
git show master@{yesterday}                               # 显示master分支昨天的状态
git log --pretty=format:'%h %s' --graph                   # 图示提交日志
git show HEAD~3
git show -s --pretty=raw 2be7fcb476
git stash                                                 # 暂存当前修改，将所有至为HEAD状态
git stash list                                            # 查看所有暂存
git stash show -p stash@{0}                               # 参考第一次暂存
git stash apply stash@{0}                                 # 应用第一次暂存
git grep "delete from"                                    # 文件中搜索文本“delete from”
git grep -e '#define' --and -e SORT_DIRENT
git gc
git fsck
```

这种上传老是超时，用ssh

前提是弄好了ssh的密匙

然后把关联的仓库地址删了，

查找关联的仓库`git remote -v`和删除`git remote rm origin`

然后关联ssh的仓库

 `git remote add origin git@github.com:Idiotmann/transmit.git`

`git push -u origin main`

快捷键 ctrl +左击 进去实现的代码

alt+左箭头 退出实现的代码

### 二：需求分析

电脑手机互传（不用微信qq 蓝牙 不注册 只扫码）

软件架构

![image-20221026213551344](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221026213551344.png)

### 三：使用的lorac

**用的一个插件**

安装 `go install github.com/silenceper/gowatch@latest` 再命令行输入gowatch 就会自动执行程序的代码

初始化

![image-20221026215027046](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221026215027046.png) 

`go mod init`后面最好加上一个包名提示版本

larca体验[zserge/lorca: Build cross-platform modern desktop apps in Go + HTML5 (github.com)](https://github.com/zserge/lorca)

怎么用，看例子

```go
ui, _ := lorca.New("", "", 480, 320)
defer ui.Close()

// Bind Go function to be available in JS. Go function may be long-running and
// blocking - in JS it's represented with a Promise.
ui.Bind("add", func(a, b int) int { return a + b })

// Call JS function from Go. Functions may be asynchronous, i.e. return promises
n := ui.Eval(`Math.random()`).Float()
fmt.Println(n)

// Call JS that calls Go and so on and so on...
m := ui.Eval(`add(2, 3)`).Int()
fmt.Println(m)

// Wait for the browser window to be closed
<-ui.Done()
```

调用本地的chrome-like浏览器

安装

`go get github.com/zserge/lorca`下载之后回更新你的go mod

试用 在例子里面找hello之类的

```go
package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
```

#### 出现问题

![image-20221026222414388](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221026222414388.png)

问题解决：[(53条消息) lorca在新版本chrome使用提示chrome正受到自动软件控制问题解决_终端小蛙的博客-CSDN博客](https://blog.csdn.net/qq_41898367/article/details/123907770)

找到本地的lorca包的lorca/ui.go的文件，注释掉"–enable-automation"

### 四：GIN体验

项目地址

[gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. (github.com)](https://github.com/gin-gonic/gin)

网址

[Gin Web Framework (gin-gonic.com)](https://gin-gonic.com/)  用的时候看实例

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()                    //创建默认的r
	router.GET("/ping", func(c *gin.Context) { //如果get到/ping，就返回这个 gin上下文（局部的全局变量）
		c.JSON(http.StatusOK, gin.H{ //h代表hash
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

```

![image-20221101120715484](C:\Users\idiotmannn\AppData\Roaming\Typora\typora-user-images\image-20221101120715484.png)

- 执行gowatch,可以监听代码所作的修改，直接看实时结果

### 五 构建main

`go mod init github.com/Idiotmann/XXXX自己的建的库名字`

构建目录及文件

如果没有导入，go get一下

```go
package main

import (
	"github.com/zserge/lorca"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var ui lorca.UI
	ui, _ = lorca.New("http://baidu.com", "", 1000, 600, "--disable--sync", "--disable-translate")
	//我们获取ctrl+c的中断，当中断时，关闭ui窗口  搜索golang 优雅退出相关的
	chSignal := make(chan os.Signal, 1) //只接受操作系统的信号 1 缓冲 通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTRAP)
	//syscall.SIGINT  syscall.SIGTRAP 订阅信号终止和中断信号，发送给chsignal     //https://www.jianshu.com/p/ae72ad58ecb6信号可以在这里面看
	//syscall.SIGINT就是ctrl+c   syscall.SIGTRAP就是系统结束进程的信号
	//select,等待
	select {
	case <-ui.Done(): //ui结束
	case <-chSignal: //阻塞当前，等chSignal有信号，取出信号，ui关闭
	}
	//为什么写后面？ 因为前面两句都是做同样的事情。所以让它放后面执行就行
	ui.Close()
}

```



![image-20221101155256468](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221101155256468.png)

功能: 关闭主线程，ui自动退出。关闭ui后，主线程自动退出

新手：

抄袭老手怎么写，尽量避免原创

尽量从网上抄代码

golang 官网 stackoverflow 大佬博客

Print 单纯打印 Printf  format格式化打印

如何看select



```go
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
```

### 六：添加静态前端文件路由

```go
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

```

然后在后面下载前端代码，放到frontend文件夹下

进入前端文件夹

`npm install`

前端工具  找到部署静态站点

[开始 | Vite 官方中文文档 (vitejs.dev)](https://cn.vitejs.dev/guide/)

![image-20221102184750785](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221102184750785.png)

默认的base是/ 我们是 /static

在 `vite.config.js` 中设置正确的 `base`

给的已经写死在上面了

```
npm run build
```

执行文件后，页面空白

`https://www.yuque.com/u29422/dg5uy8/eei6ie#fF05l`

解决办法：在 Windows 注册表中找到 HKEY_CLASSES_ROOT/.js 然后将右边的 text/plain 改为 application/javascript 最后关闭注册表，即可。

一：前端，下载src代码，编译成dist

二：直接下载dist覆盖到里面（缺点不能改）只能一次性的去部署，想长期维护，就要下载src

创建新的文件.gitignore

添加dist node_modules

git不提交的文件git status -sb 可以看提交的部分



### 七 实现接口API

#### 一 上传 

![image-20221102203204930](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221102203204930.png)

讲的go知识

因为go值是拷贝的，所以想要改变它本来的，就要传的是它的地址，而不是它的值

&和*在哪使用

&是个值

*后面是个类型

按F12 打开开发者工具

```go
// Package main
// @Description:
package main

import (
   "embed"
   "github.com/gin-gonic/gin"
   "github.com/google/uuid"
   "io/fs"
   "io/ioutil"
   "log"
   "net/http"
   "os"
   "os/exec"
   "os/signal"
   "path"
   "path/filepath"
   "strings"
)

//go:embed frontend/dist/*

//将整个目录下的文件嵌入到exe中，就不用读其他的文件，不用安装环境
//在打包go的可执行文件时，把这个目录下的所有文件打包进去
var FS embed.FS

func main() {
   go func() {
      gin.SetMode(gin.DebugMode) //gin的模式 debug 生产 测试
      router := gin.Default()    //新建路由 ，创建生成引擎中间件log和recovery
      //根目录 不需要了
      //router.GET("/", func(c *gin.Context) {
      // //request
      // c.Writer.Write([]byte("asda"))
      // //response
      //})
      //把所有静态文件变成一个变量  获得一个子文件系统 子文件系统的根由第二个参数 dir 指定
      staticFiles, _ := fs.Sub(FS, "frontend/dist")
      //获取接口
      router.POST("/api/v1/texts", TextsController)
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
            // Stat() (FileInfo, error)
            // Read([]byte) (int, error)
            // Close() error
            //}
            reader, err := staticFiles.Open("index.html")
            if err != nil {
               log.Fatal(err)
            }
            defer reader.Close()
            stat, err := reader.Stat()
            //statistics 统计   返回的接口类型
            //type FileInfo interface {
            // Name() string       // base name of the file
            // Size() int64        // length in bytes for regular files; system-dependent for others
            // Mode() FileMode     // file mode bits
            // ModTime() time.Time // modification time
            // IsDir() bool        // abbreviation for Mode().IsDir()
            // Sys() any           // underlying data source (can return nil)
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
      router.Run(":8089") //如果不开协程，会一直执行，去监听，而不会往下去执行
   }()

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
   cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8089/static/index.html")
   cmd.Start()

   chSignal := make(chan os.Signal, 1)
   signal.Notify(chSignal, os.Interrupt)

   select {
   case <-chSignal:
      cmd.Process.Kill()
   }

}
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
```

它在启动后输入数据后，会在`C:\Users\idiotmannn\AppData\Local\Temp\GoLand\uploads`创建得到的文件，因为在调试时，创建临时的exe文件

创建一个在项目所在目录，不打开命令行的

```shell
 go build -ldflags -H=windowsgui -o transmit.exe main.go

```



#### 二：局域网

```go
func AddressesController(c *gin.Context) {
	addr, _ := net.InterfaceAddrs() //获取本地的ip,多个ip地址，要找到能过访问网络的ip
	var result []string
	for _, address := range addr {
		//检查ip地址是否是回环地址  address.(*net.IPNet)断言是这个类型的
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
}
```

```go
type IPNet struct {
    IP   IP     // 网络地址
    Mask IPMask // 子网掩码
}
type Addr interface {
    Network() string // 网络名
    String() string  // 字符串格式的地址
}

type IPMask []byte
//IPMask代表一个IP地址的掩码。

type IP []byte
//IP类型是代表单个IP地址的[]byte切片。本包的函数都可以接受4字节（IPv4）和16字节（IPv6）的切片作为输入。
//注意，IP地址是IPv4地址还是IPv6地址是语义上的属性，而不取决于切片的长度：16字节的切片也可以是IPv4地址。

func (ip IP) IsLoopback() bool
//如果ip是环回地址，则返回真。
```

代码

```go
// Package main
// @Description:
package main

import (
   "embed"
   "github.com/gin-gonic/gin"
   "github.com/google/uuid"
   "github.com/zserge/lorca"
   "io/fs"
   "io/ioutil"
   "log"
   "net"
   "net/http"
   "os"
   "os/signal"
   "path"
   "path/filepath"
   "strings"
   "syscall"
)

//go:embed frontend/dist/*

//将整个目录下的文件嵌入到exe中，就不用读其他的文件，不用安装环境
//在打包go的可执行文件时，把这个目录下的所有文件打包进去
var FS embed.FS

func main() {
   go func() {
      gin.SetMode(gin.DebugMode) //gin的模式 debug 生产 测试
      router := gin.Default()    //新建路由 ，创建生成引擎中间件log和recovery
      //根目录 不需要了
      //router.GET("/", func(c *gin.Context) {
      // //request
      // c.Writer.Write([]byte("asda"))
      // //response
      //})
      //把所有静态文件变成一个变量  获得一个子文件系统 子文件系统的根由第二个参数 dir 指定
      staticFiles, _ := fs.Sub(FS, "frontend/dist")

      //获取接口
      router.GET("/api/v1/addresses", AddressesController)
      router.POST("/api/v1/texts", TextsController)
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
            // Stat() (FileInfo, error)
            // Read([]byte) (int, error)
            // Close() error
            //}
            reader, err := staticFiles.Open("index.html")
            if err != nil {
               log.Fatal(err)
            }
            defer reader.Close()
            stat, err := reader.Stat()
            //statistics 统计   返回的接口类型
            //type FileInfo interface {
            // Name() string       // base name of the file
            // Size() int64        // length in bytes for regular files; system-dependent for others
            // Mode() FileMode     // file mode bits
            // ModTime() time.Time // modification time
            // IsDir() bool        // abbreviation for Mode().IsDir()
            // Sys() any           // underlying data source (can return nil)
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
   //chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
   //cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/static/index.html")
   //cmd.Start()

   //chSignal := make(chan os.Signal, 1)
   //signal.Notify(chSignal, os.Interrupt)
   //
   //select {
   //case <-chSignal:
   // cmd.Process.Kill()
   //}

}
func AddressesController(c *gin.Context) {
   addr, _ := net.InterfaceAddrs() //获取本地的ip,多个ip地址，要找到能过访问网络的ip
   var result []string
   for _, address := range addr {
      //检查ip地址是否是回环地址  address.(*net.IPNet)断言是这个类型的
      if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
         //To4将一个IPv4地址转换为4字节表示。如果ip不是IPv4地址，To4会返回nil。
         //To16将一个IP地址转换为16字节表示。如果ip不是一个IP地址（长度错误），To16会返回nil。
         if ipnet.IP.To4() != nil {
            result = append(result, ipnet.IP.String())
         }
      }
   }
   c.JSON(http.StatusOK, gin.H{"addresses": result}) //如果地址没问题就把
}
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
```

#### 三：文件下载



#### 四：websocket

[gorilla/websocket: A fast, well-tested and widely used WebSocket implementation for Go. (github.com)](https://github.com/gorilla/websocket)

库的地址

服务器往浏览器发用websocket

浏览器往服务器发用http

发布订阅

![image-20221107105935703](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221107105935703.png)

如果有人订阅了，就把客户添加到我的列表

如果取消订阅，就从客户列表删除

如果广播一个消息，就遍历客户课表，发消息发送给客户

查看使用示例



发布订阅模式

```go
package ws

import "sync"

type Hub struct {
   // Registered clients.
   clients map[*Client]bool
   // Inbound messages from the clients.
   broadcast chan []byte
   // Register requests from the clients.
   register chan *Client
   // Unregister requests from clients.
   unregister chan *Client
}

func NewHub() *Hub {
   return &Hub{
      broadcast:  make(chan []byte),
      register:   make(chan *Client),
      unregister: make(chan *Client),
      clients:    make(map[*Client]bool),
   }
}

var once sync.Once
var singleton *Hub

//找hub的例子
func (h *Hub) Run() {
   for {
      select {
      case client := <-h.register:
         h.clients[client] = true
      case client := <-h.unregister:
         if _, ok := h.clients[client]; ok {
            delete(h.clients, client)
            close(client.send)
         }
      case message := <-h.broadcast:
         for client := range h.clients {
            select {
            case client.send <- message:
            default:
               close(client.send)
               delete(h.clients, client)
            }
         }
      }
   }
}
```

对client的定义

```go
package ws

import (
   "bytes"
   "github.com/gorilla/websocket"
   "log"
   "time"
)

const (
   // Time allowed to write a message to the peer.
   writeWait = 10 * time.Second
   // Time allowed to read the next pong message from the peer.
   pongWait = 60 * time.Second

   // Send pings to peer with this period. Must be less than pongWait.
   pingPeriod = (pongWait * 9) / 10

   // Maximum message size allowed from peer.
   maxMessageSize = 512
)

var (
   newline = []byte{'\n'}
   space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
   ReadBufferSize:  1024,
   WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
   hub *Hub
   // The websocket connection.
   conn *websocket.Conn
   // Buffered channel of outbound messages.
   send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
   defer func() {
      c.hub.unregister <- c
      c.conn.Close()
   }()
   c.conn.SetReadLimit(maxMessageSize)
   c.conn.SetReadDeadline(time.Now().Add(pongWait))
   c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
   for {
      _, message, err := c.conn.ReadMessage()
      if err != nil {
         if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
            log.Printf("error: %v", err)
         }
         break
      }
      message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
      c.hub.broadcast <- message
   }
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
   ticker := time.NewTicker(pingPeriod)
   defer func() {
      ticker.Stop()
      c.conn.Close()
   }()
   for {
      select {
      case message, ok := <-c.send:
         c.conn.SetWriteDeadline(time.Now().Add(writeWait))
         if !ok {
            c.conn.WriteMessage(websocket.CloseMessage, []byte{})
            return
         }

         w, err := c.conn.NextWriter(websocket.TextMessage)
         if err != nil {
            return
         }
         w.Write(message)

         // Add queued chat messages to the current websocket message.
         n := len(c.send)
         for i := 0; i < n; i++ {
            w.Write(newline)
            w.Write(<-c.send)
         }

         if err := w.Close(); err != nil {
            return
         }
      case <-ticker.C:
         c.conn.SetWriteDeadline(time.Now().Add(writeWait))
         if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
            return
         }
      }
   }
}
```

创建控制入口

```go
package ws

import (
   "github.com/gin-gonic/gin"
   "github.com/gorilla/websocket"
   "log"
   "net/http"
)

var wsupgrader = websocket.Upgrader{
   ReadBufferSize:  1024,
   WriteBufferSize: 1024,
   CheckOrigin: func(r *http.Request) bool {
      return true
   },
}

func wshandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
   conn, err := wsupgrader.Upgrade(w, r, nil)
   if err != nil {
      log.Println(err)
      return
   }
   client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
   client.hub.register <- client

   // Allow collection of memory referenced by the caller by doing all work in
   // new goroutines.
   go client.writePump()
   go client.readPump()
}
func HttpController(c *gin.Context, hub *Hub) {
   wshandler(hub, c.Writer, c.Request)
}
```

端口配置

```go
package config

func GetPort() string {
   return "27149"
}
```

记住将front目录移动到server下面

把27149加入白名单

ctrl +shift+i 

最后代码看我的github

![image-20221107145440505](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221107145440505.png)

![image-20221107145519169](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221107145519169.png)



现在的问题：

能够上传，无法下载，去查查哪里有问题

![image-20221107152531629](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221107152531629.png)

![image-20221107152547022](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221107152547022.png)

上面可以看到存到了发送的数据，但是整个下载功能都是有问题的

原因前端接口写错了

![image-20221107155402127](https://gitee.com/idiotmannn/picture/raw/master/imag/image-20221107155402127.png)

修改成这样

## 使用的库及框架

### 一 gin web框架

快速入门

```go
package main
import "github.com/gin-gonic/gin"
func main() {
    
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

二：优雅退出

```go
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
```