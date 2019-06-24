/**
* @Author: ASlowPerson
* @Date: 19-6-24 下午7:46
 */
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

/*
	定义函数,使用公钥加密
*/
func rsaEncryptData(fileName string, src []byte) ([]byte, error) {
	// 调用函数,获取公钥
	pubKey, err := readRsaPubKey(fileName)
	if err != nil {
		fmt.Println("readRsaPubKey err: ", err)
		return nil, err
	}

	// 加密
	// EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
	encryptInfo, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if err != nil {
		fmt.Println("EncryptPKCS1v15 err: ", err)
		return nil, err
	}
	// 返回数据
	return encryptInfo, nil
}

/*
	定义函数,使用私钥解密
*/
func rsaDecryptData(fileName string, src []byte) ([]byte, error) {
	// 调用函数,获取私钥
	priKey, err := readRsaPriKey(fileName)
	if err != nil {
		fmt.Println("readRsaPriKey err: ", err)
		return nil, err
	}

	// 解密
	// DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, src)
	if err != nil {
		fmt.Println("DecryptPKCS1v15 err: ", err)
		return nil, err
	}
	// 返回数据
	return plainText, nil
}

func main() {
	// 明文数据
	src := []byte("go语言是世界上最好的语言!")
	// 调用函数加密
	cipherText, err := rsaEncryptData("rsaPublicKey.pem", src)
	if err != nil {
		fmt.Println("加密失败!", err)
		return
	}
	fmt.Println("加密后的数据为: ", cipherText)
	// 调用函数解密
	plainText, err := rsaDecryptData("rsaPriKey.pem", cipherText)
	if err != nil {
		fmt.Println("解密失败!", err)
		return
	}

	fmt.Printf("解密后的数据为: %s\n", plainText)
}
