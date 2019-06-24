/**
* @Author: ASlowPerson
* @Date: 19-6-24 下午7:52
 */
package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

/*
	定义函数,通过公钥文件获取公钥
*/
func readRsaPubKey(fileName string) (*rsa.PublicKey, error) {
	// 1.读取公钥文件
	info, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ReadFile err :", err)
		return nil, err
	}
	// 2.base64解码，得到block
	block, _ := pem.Decode(info)
	// 3.得到der数据
	der := block.Bytes
	// 4.对der数据进行解码,得到公钥
	pubInterface, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		fmt.Println("ParsePKIXPublicKey err :", err)
		return nil, err
	}
	// 类型断言
	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("pubKey not ok!")
	}
	return pubKey, nil
}

/*
	定义函数,通过私钥文件获取私钥
*/
func readRsaPriKey(fileName string) (*rsa.PrivateKey, error) {
	// 1.读取私钥文件
	info, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ReadFile err :", err)
		return nil, err
	}
	// 2.base64解码，得到block
	block, _ := pem.Decode(info)
	// 3.得到der数据
	der := block.Bytes
	// 4.对der数据进行解码,得到私钥
	priInterface, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err :", err)
		return nil, err
	}
	// 类型断言
	priKey, ok := priInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("priKey not ok!")
	}
	return priKey, nil
}
