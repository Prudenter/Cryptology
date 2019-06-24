/**
* @Author: ASlowPerson
* @Date: 19-6-24 下午9:00
 */

package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

/*
	定义函数,实现私钥签名
*/
func rsaSignData(fileName string, src []byte) ([]byte, error) {
	// 1.调用函数,获取私钥
	priKey, err := readRsaPriKey(fileName)
	if err != nil {
		fmt.Println("readRsaPriKey err: ", err)
		return nil, err
	}

	// 2.将待签名数据通过哈希函数生成其哈希值,返回的是个[32]byte数组
	hash := sha256.Sum256(src)

	// 3.调用函数,对传入的数据进行数字签名
	// func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
	// 参数1：随机数
	// 参数2: 私钥
	// 参数3: 计算哈希的算法
	// 参数4: 原文的哈希值
	// 注意,这里第四个参数要求是个切片,但我们得到的hash是个数组,所以需要进行格式转换
	signData, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Println("SignPKCS1v15 err: ", err)
		return nil, err
	}
	return signData, nil
}

/*
	定义函数,实现公钥验证签名
*/
func rsaVerifyData(fileName string, src []byte, signData []byte) bool {
	// 1.调用函数,获取公钥
	pubKey, err := readRsaPubKey(fileName)
	if err != nil {
		fmt.Println("readRsaPubKey err: ", err)
		return false
	}

	// 2.将接收到的明文数据通过相同的哈希函数生成其哈希值,用于数据校验
	hash := sha256.Sum256(src)

	// 3.调用函数,对传入的数字签名进行认证校验
	// func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
	// 参数1: 公钥
	// 参数2: 哈希算法
	// 参数3: 本地对原文生成的哈希
	// 参数4: 待验证数字签名
	// 注意,这里第三个参数要求是个切片,但我们得到的hash是个数组,所以需要进行格式转换
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], signData)
	if err != nil {
		fmt.Println("SignPKCS1v15 err: ", err)
		return false
	}
	return true
}

func main() {
	// 明文数据
	src := []byte("实际上，数字签名和非对称加密有着非常紧密的联系，简而言之，数字签名就是通过将非对称加密反过来用而实现的。")
	// 调用函数,进行数字签名
	signData, err := rsaSignData("rsaPriKey.pem", src)
	if err != nil {
		fmt.Println("数字签名失败! ", err)
		return
	}
	fmt.Printf("签名后的数据为:%x\n", signData)

	//调用函数,进行签名认证
	result := rsaVerifyData("rsaPublicKey.pem", src, signData)
	if !result {
		fmt.Println("数字签名认证失败!")
		return
	}
	fmt.Println("数字签名认证成功!")
}
