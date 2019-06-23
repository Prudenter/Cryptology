/**
* @Author: ASlowPerson  
* @Date: 19-6-23 下午9:59
*/
package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"fmt"
	"crypto/rand"
)

//创建秘钥对，自己指定位数，位数越大，越安全，同时效率也越低
func generateRsaKeyPair(bit int) error {
	fmt.Println("创建私钥..")
	// 1.创建私钥，GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥。
	priKey, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		fmt.Println("GenerateKey err: ", err)
		return err
	}
	// 2.对私钥进行编码，生成der格式的字符串
	derText, err := x509.MarshalPKCS8PrivateKey(priKey)
	if err != nil {
		fmt.Println("MarshalPKCS8PrivateKey err: ", err)
		return err
	}
	// 3.将der字符串拼装到pem格式的数据块中
	block := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   derText,
	}
	file, err := os.Create("rsaPriKey.pem")
	if err != nil {
		fmt.Println("Create err: ", err)
		return err
	}
	defer file.Close()
	// 4.将pem格式的数据块进行base64编码，得到最终私钥
	err = pem.Encode(file, &block)
	if err != nil {
		fmt.Println("Encode err: ", err)
		return err
	}
	// 5.通过私钥得到公钥
	fmt.Println("创建公钥..")
	pubKey := priKey.PublicKey
	// 6.对公钥进行编码，生成der格式的字符串
	// 注意，这里需要传递公钥指针，否则编码时会报错
	derText, err = x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		fmt.Println("MarshalPKIXPublicKey err: ", err)
		return err
	}
	// 7.将der字符串拼装到pem格式的数据块中
	block = pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   derText,
	}
	// 8.将pem格式的数据块进行base64编码，得到最终公钥
	file, err = os.Create("rsaPublicKey.pem")
	if err != nil {
		fmt.Println("Create err: ", err)
		return err
	}
	defer file.Close()
	err = pem.Encode(file, &block)
	if err != nil {
		fmt.Println("Encode err: ", err)
		return err
	}
	return nil
}

func main() {
	bits := 1024
	err := generateRsaKeyPair(bits)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}

