/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : response.go
 Time    : 2018/9/19 14:29
 Author  : yanue
 
 - 响应相关
 
------------------------------- go ---------------------------------*/

package util

import (
	"bytes"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/yanue/go-esport-common"
	"github.com/yanue/go-esport-common/errcode"
	pb "github.com/yanue/go-esport-common/proto"
	"net/http"
	"reflect"
	"strings"
	"unsafe"
)

var (
	jsonMarshal = jsonpb.Marshaler{}
)

type responseHelper struct {
}

var Response *responseHelper = new(responseHelper)

/*
 *@note WritePbRespone http服务器针对proto.Message数据的回包，
 *@note 根据客户端的Content-Type判断回json格式还是proto格式
 *@param w http.ResponseWriter
 *@param r *http.Request
 *@param pb proto.Message
 */
func (this *responseHelper) WritePbResponse(w http.ResponseWriter, r *http.Request, pb proto.Message) {
	if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Request-From") == "x-web-json" {
		w.Header().Set("Content-Type", "application/json")
		reply, _ := json.Marshal(pb)
		common.Logs.Debug("resp json, method:%v url:%v, %v", r.Method, r.RequestURI, string(reply))
		w.Write(reply)
	} else {
		w.Header().Set("Content-Type", "application/protobuf")
		common.Logs.Debug("resp pb, method:%v url:%v, %v", r.Method, r.RequestURI, Proto2Json(pb))
		buf, _ := PBMarshal(pb)
		w.Write(buf)
	}
}

/*
 *@note WritePbRespone http服务器针对错误码的回包，
 *@note 根据客户端的Content-Type判断回json格式还是proto格式
 *@param w http.ResponseWriter
 *@param r *http.Request
 *@param code 错误码，定义见PbResult
 */
func (this *responseHelper) WriteErrResponse(w http.ResponseWriter, r *http.Request, code int32) {
	// json格式
	if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Request-From") == "x-web-json" {
		w.Header().Set("Content-Type", "application/json")

		res := pb.PJsonResult{
			ErrorCode: code,
			ErrorMsg:  errcode.GetErrMsg(code),
		}

		js, _ := json.Marshal(res)
		common.Logs.Debug("resp: " + string(js))

		w.Write(js)
	} else {
		// protobuf格式
		w.Header().Set("Content-Type", "application/protobuf")
		res := pb.PbResult{
			ErrorCode: code,
			ErrorMsg:  errcode.GetErrMsg(code),
		}

		// protobuf转换方法
		protoMar := func(v interface{}) ([]byte, error) {
			return PBMarshal(v.(*pb.PbResult))
		}

		pbuf, _ := protoMar(res)
		common.Logs.Debug("resp: " + string(pbuf))

		w.Write(pbuf)
	}
}

/*
 *@note 检测是否为PResult
 *@param msg PMessage结构
 *@return error 是PResult返回error，不是PResult返回nil
 */
func (this *responseHelper) CheckPResult(msg *pb.PMessage) int32 {
	if msg.Type == "PResult" {
		err := int32(errcode.ErrCommonUnknownError)

		var ret pb.PbResult
		if proto.Unmarshal(msg.Data, &ret) == nil {
			err = ret.ErrorCode
		}

		common.Logs.Warn("CheckPResult err=%s", errcode.GetErrMsg(err))
		return err
	}

	return errcode.No_Error
}

/**
 * 反序列化
 *
 *@param buf 二进制数据
 *@param result 反序列化的对象数据
 *@return pb.PbResult, error
 */
func PbUnmarshal(buf []byte, result proto.Message) int32 {
	pMsg := &pb.PMessage{}
	err := proto.Unmarshal(buf, pMsg)
	if err != nil {
		common.Logs.Warn("Unmarshal 1 err=%s", err.Error())
		return errcode.ErrCommonUnmarshal
	}

	switch pMsg.Type {
	case "pb.PbResult":
		res := &pb.PbResult{}
		err = proto.Unmarshal(pMsg.Data, res)
		if err != nil {
			common.Logs.Warn("Unmarshal 2 err=%s", err.Error())
			return errcode.ErrCommonUnmarshal
		}
		return res.ErrorCode
	default:
		if result == nil {
			return errcode.ErrCommonUnmarshal
		}

		err = proto.Unmarshal(pMsg.Data, result)
		if err != nil {
			common.Logs.Warn("Unmarshal 3 err=%s", err.Error())
			return errcode.ErrCommonUnmarshal
		}
	}

	return errcode.No_Error
}

/*
 *@note  对消息进行统一格式解封装，输出json
 *@param in 输入数据流
 *@return out 输出或出错信息
 */
func PBUnMarshal2Json(in []byte) (out []byte, err error) {
	msg := &pb.PMessage{}
	err = proto.Unmarshal(in, msg)
	if err != nil {
		return nil, err
	}

	common.Logs.Debug("msg:%+v", msg)

	pm := reflect.New(proto.MessageType("proto." + msg.Type).Elem()).Interface().(proto.Message)
	err = proto.Unmarshal(msg.Data, pm)
	if err != nil {
		return nil, err
	}

	return json.Marshal(pm)
}

/*
 *@note  对消息进行统一格式封装
 *@param pb 消息结构
 *@return 字节数组或error
 */
func PBMarshal(in proto.Message) ([]byte, error) {
	buf, _ := proto.Marshal(in)
	tpName := reflect.TypeOf(in).String()
	pos := strings.LastIndex(tpName, ".")

	msg := &pb.PMessage{
		Magic: 9833,
		Type:  tpName[pos+1:],
		Data:  buf,
	}

	return proto.Marshal(msg)
}

/*
 *@note  对消息进行统一格式封装
 *@param pb 消息结构
 *@return 字节数组或error
 */
func JsonMarshal(in proto.Message) ([]byte, error) {
	buf, _ := proto.Marshal(in)
	tpName := reflect.TypeOf(in).String()
	pos := strings.LastIndex(tpName, ".")

	msg := &pb.PMessage{
		Magic: 9833,
		Type:  tpName[pos+1:],
		Data:  buf,
	}

	return json.Marshal(msg)
}

/*
*@note proto转换成json字符串
*@param v 结构体proto
*@return string
 */
func Proto2Json(v interface{}) string {
	var err error
	var jsonBytes []byte

	// 判断是否proto类型
	if pbf, ok := v.(proto.Message); ok {
		var buf bytes.Buffer
		err = jsonMarshal.Marshal(&buf, pbf)
		jsonBytes = buf.Bytes()
	} else {
		// 普通json
		jsonBytes, err = json.Marshal(v)
	}

	if err != nil {
		common.Logs.Warn("Proto2Json: marshal error: %v", err)
		return ""
	}

	return *(*string)(unsafe.Pointer(&jsonBytes))
}

/*
*@note proto结构转[]byte
*@param pb 结构体
*@return []byte, int32
 */
func Proto2Byte(pb proto.Message) ([]byte, int32) {
	data, err := proto.Marshal(pb)
	if err != nil {
		common.Logs.Warn("Proto2Byte() fail, err=%s", err.Error())
		return nil, errcode.ErrCommonMarshal
	}

	return data, errcode.No_Error
}

/*
*@note []byte结构转proto
*@param pb 结构体
*@param data []byte
*@return int32
 */
func Byte2Proto(data []byte, pb proto.Message) int32 {
	err := proto.Unmarshal(data, pb)
	if err != nil {
		common.Logs.Warn("Byte2Proto() fail, err=%s", err.Error())
		return errcode.ErrCommonUnmarshal
	}

	return errcode.No_Error
}
