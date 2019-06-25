/**
* @Author: ASlowPerson
* @Date: 19-6-25 下午6:30
 */
package main

import (
	"fmt"
	"net/http"
)

/*
	定义简单的https服务器
*/
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world,this is my self-signed certificate!"))
	})

	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error {
	// 参数１：ip+端口
	// 参数２：自签名证书
	// 参数３：私钥文件
	// 参数４：处理函数
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("服务器启动失败！", err)
		return
	}
	fmt.Println("服务器启动...")
}
