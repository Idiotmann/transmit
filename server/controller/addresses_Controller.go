package controller

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

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
