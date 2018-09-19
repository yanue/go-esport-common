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
end.
`

func TestRsaSign(t *testing.T) {
	a, err := RsaSign([]byte(plainTxt))
	fmt.Println("a", err)
	b, err := RsaDecryptPublic(a)
	fmt.Println("b", err, string(b))
}

func TestRsaSign1(t *testing.T) {
	a, err := RsaEncryptPublic([]byte(plainTxt))
	fmt.Println("a", err)
	b, err := RsaDecryptPrivate(a)
	fmt.Println("b", err, string(b))
}
