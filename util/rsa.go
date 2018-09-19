/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : rsa.go
 Time    : 2018/9/19 16:42
 Author  : yanue

 - 对称加密解密: 服务端保存私钥(和公钥),客户端只拥有公钥
 - 使用场景:
   1. cli 公钥加密  						=> svr 私钥解密(RsaPrivateDecrypt)
   2. svr 私钥加密 (RsaPrivateEncrypt) 	=> cli 公钥解密
 - 服务端只关心函数: RsaPrivateEncrypt,RsaPrivateDecrypt

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
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
)

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDAV7367zhejx2JkYiGzQpS+NON4RSZ3cK9BqyphRws6H2F2pWY
wpIXMQqtMPCDoOqYHcehzj77CD3PzF4tlKaGByg/oKGXkA9yvI/qygoRMyYadN3c
CKpe3gYt3glkfg7bmfrA9N15K0wI0UZ3O2xStCGVAeGGhupBP4Tf88r2vwIDAQAB
AoGABsYQRroN/iaEH8gkSrsF1g41RRXcJ98LcKS/h+jMKTi20vPzuMGBa5eqjJbg
oYIVQO4rjaM5zJVpt9u2pVxK0PXGK9uA6vnMqZ6YoX9xk2WmT7mLHooB3NwblUn4
wTuvznWGgd7v8COSg4JUBu1IiefgemzCblYW0f/VMQ8XzQECQQD77BcAy49RJh2H
F0BV3rUmpHyrU5be5HrdKOrs25zUNx/mKoolbgz+K6BDl1uPxeHFccXUsT6mYfAg
5/D7am8fAkEAw3TEnKTz+qAw6OjnrCzRtVtO1FEG7IBcBUB9VaqKYy7mJbr8Fcdf
XTFeswEMcLNly18/u4NZH2frOVK9sNmkYQJAL+NfNA19/uUJ8+YdmrUtJl1aPY80
PMad/HCMx92vYD/iVnR7skXLPn24h2C8TQZGtqu7+YR/7kzrwrWUf5Zp7QJBAKhw
Upej19YegsfVHwHTxg2CWJbEFTFvFN45y0kuJQCAhDnzwBaMsHRBfZjJyIy/LXRr
6yKPeRiFl8LYuTxU80ECQFYQc4M/XdS4BpTX2M0PmktS8XBqfENMPGm6OTRrGZdF
xu04jk+yaFQPX9dgbkuNX1bm7p11z/tHtWgXE4gc580=
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDAV7367zhejx2JkYiGzQpS+NON
4RSZ3cK9BqyphRws6H2F2pWYwpIXMQqtMPCDoOqYHcehzj77CD3PzF4tlKaGByg/
oKGXkA9yvI/qygoRMyYadN3cCKpe3gYt3glkfg7bmfrA9N15K0wI0UZ3O2xStCGV
AeGGhupBP4Tf88r2vwIDAQAB
-----END PUBLIC KEY-----
`)

// 服务端-私钥加密
func RsaEncryptPrivate(ciphertext []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	//解析PKCS1格式的私钥
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 分片
	var data = packageData(ciphertext, pri.PublicKey.N.BitLen()/8-11)
	var plainData = make([]byte, 0, 0)
	fmt.Println("data", len(data))
	// 每片解密
	for _, d := range data {
		var p, e = rsaPrivateEncrypt(pri, d)
		if e != nil {
			return nil, e
		}
		plainData = append(plainData, p...)
	}

	// 解密
	return plainData, nil
}

// 服务端-私钥解密
func RsaDecryptPrivate(ciphertext []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	//解析PKCS1格式的私钥
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 分片
	var data = packageData(ciphertext, pri.PublicKey.N.BitLen()/8)
	var plainData = make([]byte, 0, 0)

	// 每片解密
	for _, d := range data {
		var p, e = rsa.DecryptPKCS1v15(rand.Reader, pri, d)
		if e != nil {
			return nil, e
		}
		plainData = append(plainData, p...)
	}

	// 解密
	return plainData, nil
}

// 客户端-公钥加密
func RsaEncryptPublic(plaintext []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	// 数据分片
	var data = packageData(plaintext, pub.N.BitLen()/8-11)
	var cipherData = make([]byte, 0, 0)

	// 分片加密
	for _, d := range data {
		var c, e = rsa.EncryptPKCS1v15(rand.Reader, pub, d)
		if e != nil {
			return nil, e
		}
		cipherData = append(cipherData, c...)
	}

	return cipherData, nil
}

// 客户端-公钥解密
func RsaDecryptPublic(ciphertext []byte) (res []byte, err error) {
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

	// 分片
	var data = packageData(ciphertext, pub.N.BitLen()/8)
	var plainData = make([]byte, 0, 0)

	// 每片解密
	for _, d := range data {
		var p, e = rsaPublicDecrypt(pub, d)
		if e != nil {
			return nil, e
		}
		plainData = append(plainData, p...)
	}

	// 解密
	return plainData, nil
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

// pure-go realization of RSA_public_decrypt
func rsaPublicDecrypt(pubKey *rsa.PublicKey, data []byte) ([]byte, error) {
	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(data)
	e := big.NewInt(int64(pubKey.E))
	c.Exp(m, e, pubKey.N)
	out := c.Bytes()
	skip := 0
	for i := 2; i < len(out); i++ {
		if i+1 >= len(out) {
			break
		}
		if out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}

	return out[skip:], nil
}

// pure-go realization of RSA_private_encrypt
func rsaPrivateEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {
	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// rfc2313, section 8: The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		err = fmt.Errorf("input size")
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		err = fmt.Errorf("encryption")
		return
	}
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}
