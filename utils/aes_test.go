package utils

import (
	"fmt"
	"testing"
)

func Test_a(t *testing.T)  {
	//设置摘要
	//明文
	origData := []byte("123456")

	//加密
	en := AESEncrypt(origData)
	fmt.Printf("%s, %d, \n", en, len(en))

	//encodeString := base64.StdEncoding.EncodeToString(en)
	//fmt.Printf("11-> %s\n", encodeString)
	//dd, err := base64.StdEncoding.DecodeString(encodeString)
	//fmt.Printf("22-> %s, %v\n", string(dd), err)
	//if (err != nil) {
	//	fmt.Printf("22-> %v\n", string(dd))
	//}else {
	//	//log.Fatalf("22-> %v", string(dd))
	//}
	//de2 := AESDecrypt(dd, key)
	////fmt.Println(string(de))
	//fmt.Println(string(de2))
	//fmt.Println("-----------------------------")

	//for len(en) > 0 {
	//	ch, size := utf8.DecodeRune(en)
	//	en = en[size:]
	//	fmt.Printf("%c ", ch)
	//}
	//fmt.Println()
	//解密
	de := AESDecrypt(en)
	//fmt.Println(string(de))
	fmt.Println(de)

	fmt.Println(CheckPWD("123456",
		"x04jpoIrc8/mvNRqAG59Wg=="))

	fmt.Println("------------------------------------")

}

func TestMd5(t *testing.T) {
	b := []byte("1234567890abcdef")
	t.Logf("text: %s, md5: %s", b, Md5(b))
}

func TestEncrypt(t *testing.T) {
	//要加密的字符串
	cipherkey := []byte("1234dsa5678dsadddsadsadsad90abcdef")
	// 加密后的string
	ciphertext := AESEncrypt2(cipherkey, []byte("text123dsadsadsad4"))
	t.Logf("ciphertext: %s", ciphertext)

	rawtext, err := AESDecrypt2(cipherkey, ciphertext)
	t.Logf("rawtext: %s, error: %v", rawtext, err)
	//if string(rawtext) != "text1234" {
	//	t.Fatalf("expect: %s, but get: %s", "text1234", rawtext)
	//}
}
