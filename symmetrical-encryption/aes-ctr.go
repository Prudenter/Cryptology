/**
* @Author: ASlowPerson  
* @Date: 19-6-23 下午9:37
*/
package main

import (
	"crypto/aes"
	"bytes"
	"crypto/cipher"
	"fmt"
)

/*
	背景：AES + CTR
	1.加密算法：AES
	秘钥：16
	分组长度：16
	2.分组模式：CTR
	不需要填充
	需要提供数字
*/

// 输入明文，输出密文
func aesCtrEncrypt(plainText, key []byte) ([]byte, error) {
	// 第一步：创建aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("NewCipher err:", err)
		return nil, err
	}
	// 打印aes分组长度
	fmt.Println("blocksize:", block.BlockSize())

	//第二步：创建分组模式ct
	// 创建iv,与算法长度一致，16字节
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	// Stream接口代表一个流模式的加/解密器
	stream := cipher.NewCTR(block, iv)

	// 第三步：加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)
	return dst, nil
}

// 输入密文，得到明文
func aesCtrDecrypt(encrptData []byte, key []byte) ([]byte, error) {
	return aesCtrEncrypt(encrptData, key)
}

func main() {
	// 明文
	src := "Stream接口代表一个流模式的加/解密器。"
	// 对称秘钥，ase,16字节
	key := "1234567887654321"
	// 调用加密函数
	encrptData, err := aesCtrEncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println("aesCtrEncrypt err:", err)
		return
	}
	fmt.Printf("加密后的数据:%x\n", encrptData)

	// 调用解密函数
	plainText, err := aesCtrDecrypt(encrptData, []byte(key))
	if err != nil {
		fmt.Println("aesCtrDecrypt err:", err)
		return
	}
	fmt.Printf("解密后的数据: %s\n", plainText)
}
