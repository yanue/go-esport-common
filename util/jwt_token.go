/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : this.go
 Time    : 2018/9/20 18:26
 Author  : yanue

 - 功能介绍: 用于 jwt token 生成及验证
 - header和payload部分格式,默认是proto,也可以是json
 - 生成的token是base64url格式,可用于url参数输入方式

------------------------------- go ---------------------------------*/

package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	pb "github.com/yanue/go-esport-common/proto"
	"strings"
	"time"
)

const (
	jwtSalt        = "A`0#!t45+Dylnr?k^9Sp*k>*rC&+W2!0" // jwt 盐
	jwtPayloadType = "proto"                            // 格式化方式: proto|JWT,默认proto
)

// JWT 完整的本体
type jwtToken struct {
	*pb.PJwtToken
}

var JwtToken *jwtToken

func init() {
	// 初始化
	JwtToken = &jwtToken{
		PJwtToken: &pb.PJwtToken{},
	}
}

/*
*@note 生成jwt token
*@param uid 用户ID
*@param os 	0: "ANDROID", 1: "IOS",	2: "WEB",
*@param loginType 0: "ACCOUNT", 1: "PHONE",2: "WECHAT",3: "QQ",
*@return 生成的jwt
 */
func (this *jwtToken) Generate(uid int, os pb.Os, loginType pb.ELoginType) (string, error) {
	// header
	this.Header = &pb.PJwtHeader{
		Alg: "HS256",        // HMAC SHA256
		Typ: jwtPayloadType, // JWT | proto
	}

	// payload
	this.Payload = &pb.PJwtPayload{
		Uid:       int32(uid),
		Time:      time.Now().Unix(),
		Os:        os,
		LoginType: loginType,
	}

	// encode
	return this.encode()
}

/*
*@note 验证token签名是否正确,并将内容解析出来
*@param tokenStr - jwt token字符串
*@return *jwtToken
 */
func (this *jwtToken) Verify(tokenStr string) (*jwtToken, error) {
	var jwtHeader = &pb.PJwtHeader{}
	var result = &pb.PJwtPayload{}

	// 分拆
	arr := strings.Split(tokenStr, ".")
	if len(arr) != 3 {
		return this, errors.New("invalid token")
	}

	// 验证签名是否正确
	format := arr[0] + "." + arr[1]
	signature := this.getHmacCode(format)
	if signature != arr[2] {
		return this, errors.New("invalid signature")
	}

	header, err := base64.RawURLEncoding.DecodeString(arr[0])
	if err != nil {
		return this, errors.New("invalid header")
	}

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		err = proto.Unmarshal(header, jwtHeader)
	} else {
		err = json.Unmarshal(header, jwtHeader)
	}

	if err != nil {
		return this, errors.New("invalid header")
	}

	payload, err := base64.RawURLEncoding.DecodeString(arr[1])
	if err != nil {
		return this, errors.New("invalid payload")
	}

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		err = proto.Unmarshal(payload, result)
	} else {
		err = json.Unmarshal(payload, result)
	}

	if err != nil {
		return this, errors.New("invalid payload")
	}

	this.Header = jwtHeader
	this.Payload = result

	return this, nil
}

// 转成符合 JWT 标准的字符串
func (this *jwtToken) encode() (string, error) {
	var err error

	// 以proto格式解析
	var header []byte
	if this.Header.Typ == "proto" {
		header, err = proto.Marshal(this.Header)
	} else {
		header, err = json.Marshal(this.Header)
	}
	if err != nil {
		return "", errors.New("marshal header")
	}

	headerString := base64.RawURLEncoding.EncodeToString(header)
	var payload []byte

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		payload, err = proto.Marshal(this.Payload)
	} else {
		payload, err = json.Marshal(this.Payload)
	}

	if err != nil {
		return "", errors.New("marshal payload")
	}

	payloadString := base64.RawURLEncoding.EncodeToString(payload)

	format := headerString + "." + payloadString
	signature := this.getHmacCode(format)

	return format + "." + signature, nil
}

// 生成签名
func (this *jwtToken) getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte(jwtSalt))
	h.Write([]byte(s))
	key := h.Sum(nil)
	return hex.EncodeToString(key)
}
