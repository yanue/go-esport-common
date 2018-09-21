/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : rsa.go
 Time    : 2018/9/19 16:42
 Author  : yanue

 - 对称加密解密: 服务端保存私钥(和公钥),客户端只拥有公钥

 - 使用场景:
   1. cli 公钥加密  						=> svr 私钥解密(RsaPrivateDecrypt)
   2. svr 私钥加密 (RsaPrivateEncrypt) 	=> cli 公钥解密

 - 服务端只关心函数: RsaPrivateEncrypt,RsaPrivateDecrypt

 - 分片加密解密: (1024bit)
   加密时分片长度: pri.PublicKey.N.BitLen()/8-11 = 117
   解密时分片长度: pri.PublicKey.N.BitLen()/8    = 128

 - 用法:
	util.Rsa.RsaEncryptPrivate()

 - openssl生成私钥
	openssl genrsa -out rsa_private_key.pem 1024

 - openssl生成公钥
	openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

------------------------------- go ---------------------------------*/

package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDTsw7li1oxyK5reumWnzQQ/kBlhFQN7PV09cTaokxUswkh9O8D
jAci/eJk9kXOctnwUHOMyjXvt+GSkeCGHWk9m+DaCdCt6MaKjTm1EUbmy5tS68Pa
NovDdlPXVnpt8gRornWmTFxQ7FAxOVWFguaQEHRL3JrfLMeM690GGIUgPQIDAQAB
AoGACTCju+lBwBmDIN1UGJrOqtIuv3lwIK6htTMaGZekEqU3B0dXvOKuSKGW22Up
gJ3RwjHr4jfBAagM3c0BXzEVqWfGoZbwfX8ZWlyxcYOV8SVeRQXxOJRDyU+tbtgm
JjHgMSfbBr/RIM9B69cxpIS3oEAJG/U5iafw9/ZSVzb0/KECQQD1NJJohsc/h7Ry
6CbJoPoqN4eyhgc6wFl21g3kVuQHISBhLy7rZUZd/Aka9A3Bw+09UgebfCFUFvJ5
iP7MUqu1AkEA3QThLb56oRIS1jgpfzdcLkPQmn2TJG8KThGryg1fQ3xW9vc6NG9K
lriog2M0DiWyyDaxdoZpfWQ3h6B3HnfHaQJAFDJOVNm1E6CD1msUtsrRkCSewq+T
bN1nAQjEgChAA+5QknCmdrESyK73uQadE3al1cUp5z6kKB7zvdrw0beFeQJBAJl/
5hQrEmgDcWmuH8Pm4vKOzrY9OJA5PmLyCumNV/g6xvtGwPnhwV/kZ8S4hVK+A+jh
c2bp+yHHFHnxjElwzuECQQDUn05MT2r1TascaAwmoqAclbs0lrgDiWGbCBD1Mco4
bbHRtA7qAVxiU+MmDiesBfyzpzL8+qKb9VSyZjv2AAa6
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDTsw7li1oxyK5reumWnzQQ/kBl
hFQN7PV09cTaokxUswkh9O8DjAci/eJk9kXOctnwUHOMyjXvt+GSkeCGHWk9m+Da
CdCt6MaKjTm1EUbmy5tS68PaNovDdlPXVnpt8gRornWmTFxQ7FAxOVWFguaQEHRL
3JrfLMeM690GGIUgPQIDAQAB
-----END PUBLIC KEY-----
`)

type rsaHelper struct {
}

var Rsa *rsaHelper = new(rsaHelper)

// 服务端-私钥加密
func (this *rsaHelper) RsaEncryptPrivate(cipherText string) (res string, err error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("private key error!")
		return
	}

	//解析PKCS1格式的私钥
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	// 分片
	var data = this.packageData([]byte(cipherText), pri.PublicKey.N.BitLen()/8-11)
	var plainData = make([]byte, 0, 0)

	// 每片解密
	for _, d := range data {
		var p, e = rsaPrivateEncrypt(pri, d)
		if e != nil {
			err = e
			return
		}
		plainData = append(plainData, p...)
	}

	// base64加密
	return base64.RawURLEncoding.EncodeToString(plainData), nil
}

// 服务端-私钥解密
func (this *rsaHelper) RsaDecryptPrivate(cipherText string) (res string, err error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("private key error!")
		return
	}

	//解析PKCS1格式的私钥
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	// base64解密
	originalData, err := base64.RawURLEncoding.DecodeString(cipherText)
	if err != nil {
		return
	}

	// 分片
	var data = this.packageData(originalData, pri.PublicKey.N.BitLen()/8)
	var plainData = make([]byte, 0, 0)

	// 每片解密
	for _, d := range data {
		var p, e = rsa.DecryptPKCS1v15(rand.Reader, pri, d)
		if e != nil {
			err = e
			return
		}
		plainData = append(plainData, p...)
	}

	// 解密
	return string(plainData), nil
}

// 客户端-公钥加密
func (this *rsaHelper) RsaEncryptPublic(plainText string) (res string, err error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("public key error")
		return
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	// 数据分片
	var data = this.packageData([]byte(plainText), pub.N.BitLen()/8-11)
	var cipherData = make([]byte, 0, 0)

	// 分片加密
	for _, d := range data {
		var c, e = rsa.EncryptPKCS1v15(rand.Reader, pub, d)
		if e != nil {
			err = e
			return
		}
		cipherData = append(cipherData, c...)
	}

	// base64加密
	return base64.RawURLEncoding.EncodeToString(cipherData), nil
}

// 客户端-公钥解密
func (this *rsaHelper) RsaDecryptPublic(cipherText string) (res string, err error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("public key error")
		return
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	// base64解密
	originalData, err := base64.RawURLEncoding.DecodeString(cipherText)
	if err != nil {
		return
	}

	// 分片
	var data = this.packageData(originalData, pub.N.BitLen()/8)
	var plainData = make([]byte, 0, 0)

	// 每片解密
	for _, d := range data {
		var p, e = rsaPublicDecrypt(pub, d)
		if e != nil {
			err = e
			return
		}
		plainData = append(plainData, p...)
	}

	return string(plainData), nil
}

// 私钥签名
func (this *rsaHelper) RsaSign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	//解析PKCS1格式的私钥
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hashed)
}

// 公钥验证
func (this *rsaHelper) RsaSignVerify(data []byte, signature []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

/*
*@note 分片处理
*@param originalData 原生数据
*@param packageSize 没片大小
*@return
 */
func (this *rsaHelper) packageData(originalData []byte, packageSize int) (r [][]byte) {
	var src = make([]byte, len(originalData))
	copy(src, originalData)

	r = make([][]byte, 0)

	if len(src) <= packageSize {
		return append(r, src)
	}

	for len(src) > 0 {
		var p = src[:packageSize]
		r = append(r, p)
		src = src[packageSize:]
		if len(src) <= packageSize {
			r = append(r, src)
			break
		}
	}

	return r
}
