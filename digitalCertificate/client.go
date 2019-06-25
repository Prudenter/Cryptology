/**
* @Author: ASlowPerson
* @Date: 19-6-25 下午6:58
 */
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	自定义http客户端，访问https服务器
*/
func main() {
	/*
		访问服务器时,服务器会返回自己的证书,, 但是我的client没有本法认证这个证书，所以请求失败了。
		因此我们需要告诉client哪些证书是我们认证的
		因为我们是自签名的数字证书,"server.crt"就是当前服务器的根证书
	*/

	// 1.读取证书
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	// 2.创建ca池
	caPool := x509.NewCertPool()

	// 3.将我们认可的根证书添加到ca池中
	ok := caPool.AppendCertsFromPEM(caCert)
	if !ok {
		fmt.Println("添加到ca池失败!")
		return
	}

	// 4.创建tls结构,填入pool
	// tls,早期叫ssl,安全套接层,现在叫tls,传输层安全性协议
	config := tls.Config{
		// 将拼凑好的ca池配置到tls配置结构中
		RootCAs: caPool,
	}

	// 5.创建client通道
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &config,
		},
	}

	//　6.发起get请求
	// 注意,这里需要根据证书中之前定义好的域名去系统中配置好,linux下是去编辑配置 /etc/hosts文件
	response, err := client.Get("https://www.test.com:8080")

	if err != nil {
		fmt.Println("Get err:", err)
		return
	}

	// 7.获取body数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll err:", err)
		return
	}
	defer response.Body.Close()

	fmt.Printf("body:%s\n", body)
}
