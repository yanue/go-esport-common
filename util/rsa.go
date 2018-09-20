/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : rsa.go
 Time    : 2018/9/19 16:42
 Author  : yanue

 - 对称加密解密: 服务端保存私钥(和公钥),客户端只拥有公钥

 - 使用场景:
   1. cli 公钥加密  						=> svr 私钥解密(RsaPrivateDecrypt)
   2. svr 私钥加密 (RsaPrivateEncrypt) 	=> cli 公钥解密

 - 服务端只关心函数: RsaPrivateEncrypt,RsaPrivateDecrypt

 - 分片加密解密: (2048bit)
   加密时分片长度: pri.PublicKey.N.BitLen()/8-11 = 245
   解密时分片长度: pri.PublicKey.N.BitLen()/8    = 256

 - openssl生成私钥
	openssl genrsa -out rsa_private_key.pem 2048

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
//openssl genrsa -out rsa_private_key.pem 2048
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEArZ2Uc3jYTB+XS3NfWAaNWbT/5LXbKnd5N7gtDN0PrMf8GlDI
f8Rf60zJRPD7lfGOnVhVJgHzLpxkq/NlNWb0rTQQ5RHt71mmiGLTiPlM6OLeB4g+
fAHjRyWf0ijbRnX0tVkRlDOWllcSYHFy0jrTVJp9foTHcntYqnk5a0Q7/S+IGEBr
SyBVbfgkthjPCKEI1mTxc07OOQhtCzo7pDI98D9mD7Yx28NYIxOzeISYCS+VqWrk
plaI1uWUNE/nNhh2qYffg3Lf50I8UnOcHGesTjmnUskUfW00dRCVlgGFLilrptuL
OUbXcaqbdkkLreqlMH43j5WogDFRGhT6GX3FrwIDAQABAoIBAGs6gT6Ua5sQg+Qw
3LlESrcWKFn8y+E9qxtz4DcqrYy8c4NZU4w+IDP21/SBlhF1AO1DakuwNp8aLr9Y
87B45zO2jZy9ZyRGTam8yAO4Xf0UaadjSZxTmikOHtixLUUmgz4iRc4v0pkGLC5u
w0j+1hlR1aJ3pauabRfVCVar7RUg6p0ZkZmnTBMalRKpHY8Gzgb66oA4v/A56VFK
4cIy/pIV49VHdE5djWuDW/BDJvuv8Y/xhgWp4XGKjRyayRrcDSAHy0T5rdYg/H84
qROTglfaEY4MJxMhYrPD2GfO7MUSgB545MwqpRA8E7XioGmRBEwNYyaOAaPZV9A0
1eJOoeECgYEA26ZGkalZfVBCmDCl0HWPqYRO3ehLm6u0+r0NiS59aG7mwlWBUmDU
1X67O9yKBuj1A7h6EGUQnhi1mKcD4fy9r2mMCNMC0rGS+f27VmLUhFst23KLVPz6
QvN2YmrdPcnvxbUHtyNTsSZLdZL/H99hS5yK+GY9yqKMt96soMGzI/8CgYEAylkE
ndAKq7PLKkru2OeGJIj4pr8p07V01J0D+qkiEtcsmIALPAxCPicWs45ZRqLRtPyD
oRAJcM8GA9Eywwogw1dAnDau/Roq4BikQm9Dff42cmBHIGr2JDp6nsOzfS6cKjxt
AHIqayt7dQir1H4/5akdPZsj8M1wgu83tmBonlECgYB0h41u38qTWg5KkZyWsJgM
Fh6FSiU6rGjykXPp8Jkl25hfR1+5pZekwHxy8Ljlm5fJZoiTxBqB1ZgaKZk8voqf
0j4xvEkGIKFaMYu+8+XNZlY401cqOqBG/sUyx4Eis8yaNkWmmn5fQHLOKLNjZG5I
3/82c3+azowbTG6HRtxUXwKBgB5vDxuxS7mRfDArPwtOn0VleIiT3fWiqCTGTO/p
el99D48MSyRH77qrZGWzNkhCeuoOxLl30QOvj4cJcuoU3uKif+w+6UjWI7a63hHD
7FHJ52SCiJAeplDCnui8JIXiech8eCSGB01BJ/ttR3LZXkDrk6NNbzVroM2Ar091
5qZRAoGBAKcByMleCzYvrrxZ6vVIg/MI6yjxiypOqUc/kcrzn4gY3j/sszPoqLjX
4ZfHICipLwWHv4fUvu+5z+RgSQ51cNxZCyMgL+23h9CnkpkvN/ZOUbYHNbn5pzJS
EwNHBgYK/oZ8+uxjulmA+kNEaq+1kuUq1d/WDCBD+5tXYYVoTZcU
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArZ2Uc3jYTB+XS3NfWAaN
WbT/5LXbKnd5N7gtDN0PrMf8GlDIf8Rf60zJRPD7lfGOnVhVJgHzLpxkq/NlNWb0
rTQQ5RHt71mmiGLTiPlM6OLeB4g+fAHjRyWf0ijbRnX0tVkRlDOWllcSYHFy0jrT
VJp9foTHcntYqnk5a0Q7/S+IGEBrSyBVbfgkthjPCKEI1mTxc07OOQhtCzo7pDI9
8D9mD7Yx28NYIxOzeISYCS+VqWrkplaI1uWUNE/nNhh2qYffg3Lf50I8UnOcHGes
TjmnUskUfW00dRCVlgGFLilrptuLOUbXcaqbdkkLreqlMH43j5WogDFRGhT6GX3F
rwIDAQAB
-----END PUBLIC KEY-----
`)

// 服务端-私钥加密
func RsaEncryptPrivate(cipherText string) (res string, err error) {
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
	var data = packageData([]byte(cipherText), pri.PublicKey.N.BitLen()/8-11)
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
func RsaDecryptPrivate(cipherText string) (res string, err error) {
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
	var data = packageData(originalData, pri.PublicKey.N.BitLen()/8)
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
func RsaEncryptPublic(plainText string) (res string, err error) {
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
	var data = packageData([]byte(plainText), pub.N.BitLen()/8-11)
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
func RsaDecryptPublic(cipherText string) (res string, err error) {
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
	var data = packageData(originalData, pub.N.BitLen()/8)
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
func RsaSign(data []byte) ([]byte, error) {
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
func RsaSignVerify(data []byte, signature []byte) error {
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
func packageData(originalData []byte, packageSize int) (r [][]byte) {
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