/**
* @Author: ASlowPerson  
* @Date: 19-6-23 下午9:51
*/
package main

import (
	"crypto/des"
	"bytes"
	"crypto/cipher"
	"errors"
	"fmt"
)

/*
	背景：DES + CBC
	1.加密算法：DES
	des: 秘钥：8字节，分组长度：8字节
	2.分组模式：CBC
	cbc: 1.长度与算法相同(8字节) 2. 需要填充
*/

//输入明文，输出密文
func desCBCEncrypt(plainText, key []byte) ([]byte, error) {
	// 第一步：创建des密码接口, 输入秘钥，返回接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	// 创建一个8字节的初始化向量
	iv := bytes.Repeat([]byte("1"), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 第三步：填充
	plainText, err = paddingData(plainText, blockSize)
	if err != nil {
		return nil, err
	}

	// 第四步：加密
	blockMode.CryptBlocks(plainText, plainText)
	return plainText, nil
}

//输入密文，得到明文
func desCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	//第一步：创建des密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//第二步：创建cbc分组
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	mode := cipher.NewCBCDecrypter(block, iv)
	//第三步：解密
	mode.CryptBlocks(encryptData, encryptData)

	//第四步: 去除填充
	encryptData, err = unPaddingData(encryptData)
	if err != nil {
		return nil, err
	}

	return encryptData, nil
}

// 填充数据
func paddingData(src []byte, blockSize int) ([]byte, error) {
	if src == nil {
		return nil, errors.New("src长度不能小于0")
	}
	fmt.Println("调用paddingData")
	// 1.得到分组之后剩余的长度
	leftNumber := len(src) % blockSize

	// 2.得到需要填充的个数
	needNumber := blockSize - leftNumber

	// 3.创建一个slice
	newSlice := bytes.Repeat([]byte{byte(needNumber)}, needNumber) //b := byte(needNumber)
	fmt.Printf("newSclie : %v\n", newSlice)
	// 4.将新切片追加到src
	src = append(src, newSlice...)
	return src, nil
}

// 解密之后去除填充的数据
func unPaddingData(src []byte) ([]byte, error) {
	if src == nil {
		return nil, errors.New("src长度不能小于0")
	}
	fmt.Println("调用paddingData")
	// 1.获取最后一个字符
	lastChar := src[len(src)-1]

	// 2.得到需要填充的个数
	num := int(lastChar)

	//3. 截取切片(左闭右开)
	return src[:len(src)-num], nil
}

func main() {
	src := "离离原上草，一岁一枯荣"
	key := "12345678"
	encryptData, err := desCBCEncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println("desCBCEncrypt err:", err)
		return
	}
	fmt.Printf("加密后的数据:%x\n", encryptData)
	//调用解密函数
	plainText, err := desCBCDecrypt(encryptData, []byte(key))
	if err != nil {
		fmt.Println("desCBCDecrypt err:", err)
		return
	}

	fmt.Printf("解密后的数据: %s\n", plainText)
	fmt.Printf("解密后的数据 hex : %x\n", plainText)
}
