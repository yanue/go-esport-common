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
		PJwtToken: &pb.PJwtToken{
			// header
			Header: &pb.PJwtHeader{
				Alg: "HS256",        // HMAC SHA256
				Typ: jwtPayloadType, // JWT | proto
			},

			// Payload
			Payload: &pb.PJwtPayload{
				Device: &pb.PDevice{},
			},
		},
	}
}

/*
*@note 生成jwt token
*@param uid 用户ID
*@param os 	0: "ANDROID", 1: "IOS",	2: "WEB",
*@param loginType 0: "ACCOUNT", 1: "PHONE",2: "WECHAT",3: "QQ",
*@param imei 设备唯一标识
*@return 生成的jwt
 */
func (this *jwtToken) Generate(uid int, loginType pb.ELoginType, device *pb.PDevice) (token, payloadStr string, err error) {
	// payload
	this.Payload = &pb.PJwtPayload{
		Uid:       int32(uid),
		Time:      time.Now().Unix(),
		LoginType: loginType,
		Device:    device,
	}

	// 转成符合 JWT 标准的字符串
	var header []byte
	var payload []byte

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		header, err = proto.Marshal(this.Header)
	} else {
		header, err = json.Marshal(this.Header)
	}
	if err != nil {
		err = errors.New("marshal header")
		return
	}

	headerString := base64.RawURLEncoding.EncodeToString(header)

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		payload, err = proto.Marshal(this.Payload)
	} else {
		payload, err = json.Marshal(this.Payload)
	}

	if err != nil {
		err = errors.New("marshal payload")
		return
	}

	payloadStr = base64.RawURLEncoding.EncodeToString(payload)

	format := headerString + "." + payloadStr
	signature := this.getHmacCode(format)

	return format + "." + signature, payloadStr, nil
}

/*
*@note 验证token签名是否正确,并将内容解析出来
*@param tokenStr - jwt token字符串
*@return *jwtToken
 */
func (this *jwtToken) Verify(tokenStr string) (token *jwtToken, payload string, err error) {
	var jwtHeader = this.Header
	var result = this.Payload

	// 分拆
	arr := strings.Split(tokenStr, ".")
	if len(arr) != 3 {
		err = errors.New("invalid token")
		return
	}

	// 验证签名是否正确
	format := arr[0] + "." + arr[1]
	signature := this.getHmacCode(format)
	if signature != arr[2] {
		err = errors.New("invalid signature")
		return
	}

	header, err := base64.RawURLEncoding.DecodeString(arr[0])
	if err != nil {
		err = errors.New("invalid header base64")
		return
	}

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		err = proto.Unmarshal(header, jwtHeader)
	} else {
		err = json.Unmarshal(header, jwtHeader)
	}

	if err != nil {
		err = errors.New("invalid header Unmarshal")
		return
	}

	payloadBuf, err := base64.RawURLEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("invalid payload base64")
		return
	}

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		err = proto.Unmarshal(payloadBuf, result)
	} else {
		err = json.Unmarshal(payloadBuf, result)
	}

	if err != nil {
		err = errors.New("invalid payload Unmarshal")
		return
	}

	this.Header = jwtHeader
	this.Payload = result

	return this, arr[1], nil
}

/*
*@note 验证token签名是否正确,并将内容解析出来
*@param tokenStr - jwt token字符串
*@return *jwtToken
 */
func (this *jwtToken) ParsePayload(payloadStr string) (result *pb.PJwtPayload, err error) {
	result = &pb.PJwtPayload{}

	payload, err := base64.RawURLEncoding.DecodeString(payloadStr)
	if err != nil {
		err = errors.New("invalid payload base64")
		return
	}

	// 以proto格式解析
	if this.Header.Typ == "proto" {
		err = proto.Unmarshal(payload, result)
	} else {
		err = json.Unmarshal(payload, result)
	}

	if err != nil {
		err = errors.New("invalid payload Unmarshal")
		return
	}

	return result, nil
}

// 生成签名
func (this *jwtToken) getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte(jwtSalt))
	h.Write([]byte(s))
	key := h.Sum(nil)
	return hex.EncodeToString(key)
}
