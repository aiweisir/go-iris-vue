package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/**
AES加密解密
使用CBC模式+PKCS7 填充方式实现AES的加密和解密
*/
const key = "ARRSWdczx13213EDDSWQ!!@W"

// 校验密码
func CheckPWD(password, enPassword string) bool {
	de := AESDecrypt(enPassword)
	if password == de {
		return true
	}
	return false
}

// -----------------------------------------------------------
// ----------------------- 解密 ------------------------------
// -----------------------------------------------------------
// 先base64转码，再解密
func AESDecrypt(baseStr string) string {
	crypted, err := base64.StdEncoding.DecodeString(baseStr)
	if err != nil {
		fmt.Println("base64 encoding 错误")
	}

	block, _ := aes.NewCipher([]byte(key))
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, []byte(key)[:blockSize])
	originPWD := make([]byte, len(crypted))
	blockMode.CryptBlocks(originPWD, crypted)
	originPWD = pkcs7_unPadding(originPWD)
	return string(originPWD)
}

// 补码
func pkcs7_unPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}

// -----------------------------------------------------------
// ----------------------- 加密 ------------------------------
// -----------------------------------------------------------
// 加密后再base64编码成string
func AESEncrypt(originPWD []byte) string {
	//获取block块
	block, _ := aes.NewCipher([]byte(key))
	//补码
	originPWD = pkcs7_padding(originPWD, block.BlockSize())
	//加密模式，
	blockMode := cipher.NewCBCEncrypter(block, []byte(key)[:block.BlockSize()])
	//创建明文长度的数组
	crypted := make([]byte, len(originPWD))
	//加密明文
	blockMode.CryptBlocks(crypted, originPWD)

	return base64.StdEncoding.EncodeToString(crypted)
}

// 补码
func pkcs7_padding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padtext...)
}
