/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : rsa_test.go
 Time    : 2018/9/19 16:51
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"fmt"
	"testing"
)

var plainTxt = `
start.
SA加解密中必须考虑到的密钥长度、明文长度和密文长度问题。明文长度需要小于密钥长度，而密文长度则等于密钥长度。因此当加密内容长度大于密钥长度时，有效的RSA加解密就需要对内容进行分段。
对应2048位密钥时,加密时分片长度是245,解密时分片长度是256
end.
`

func TestRsaEncryptPrivate(t *testing.T) {
	// 私钥加密,公钥解密
	a, err := RsaEncryptPrivate(plainTxt)
	fmt.Println("a", err, a)
	b, err := RsaDecryptPublic(a)
	fmt.Println("b", err, string(b))
}

func TestRsaEncryptPublic(t *testing.T) {
	// 公钥加密,私钥解密
	a, err := RsaEncryptPublic(plainTxt)
	fmt.Println("a", err, a)
	b, err := RsaDecryptPrivate(a)
	fmt.Println("b", err, string(b))
}

func TestRsaSign(t *testing.T) {
	// 私钥签名,公钥验证
	a, err := RsaSign([]byte(plainTxt))
	fmt.Println("a", err)
	b := RsaSignVerify([]byte(plainTxt), a)
	fmt.Println("b", b)
}
