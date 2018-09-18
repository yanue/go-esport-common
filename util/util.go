/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : util.go
 Time    : 2018/9/14 17:32
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"bytes"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/yanue/go-esport-common"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

const DEF_HTTP_TIMEOUT = time.Second * 10

/*
*@note 得到当前exe程序执行的不目录
 */
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

/**
 * 反序列化
 *
 *@param buf 二进制数据
 *@param result 反序列化的对象数据
 *@return PResult, error
 */
func Unmarshal(buf []byte, result proto.Message) int32 {
	pMsg := &PMessage{}
	err := proto.Unmarshal(buf, pMsg)
	if err != nil {
		common.Logs.Warn("Unmarshal 1 err=%s", err.Error())
		return Err_Com_UnMarshal
	}

	switch pMsg.Type {
	case "PResult":
		res := &PResult{}
		err = proto.Unmarshal(pMsg.Data, res)
		if err != nil {
			common.Logs.Warn("Unmarshal 2 err=%s", err.Error())
			return Err_Com_UnMarshal
		}
		return res.ErrorCode
	default:
		if result == nil {
			return Err_Com_UnMarshal
		}

		err = proto.Unmarshal(pMsg.Data, result)
		if err != nil {
			common.Logs.Warn("Unmarshal 3 err=%s", err.Error())
			return Err_Com_UnMarshal
		}
	}
	return No_Error
}

/*
 *@note WritePbRespone http服务器针对proto.Message数据的回包，
 *@note 根据客户端的Content-Type判断回json格式还是proto格式
 *@param w http.ResponseWriter
 *@param r *http.Request
 *@param pb proto.Message
 */
func WritePbRespone(w http.ResponseWriter, r *http.Request, pb proto.Message) {
	if r.Header.Get("Content-Type") == "application/json" ||
		r.Header.Get("Request-From") == "x-web-json" {
		w.Header().Set("Content-Type", "application/json")
		reply, _ := json.Marshal(pb)
		common.Logs.Debug("resp json, method:%v url:%v, %v", r.Method, r.RequestURI, string(reply))
		w.Write(reply)
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
		common.Logs.Debug("resp pb, method:%v url:%v, %v", r.Method, r.RequestURI, SimpleJson(pb))
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
func WriteErrRespone(w http.ResponseWriter, r *http.Request, code int32) {
	if r.Header.Get("Content-Type") == "application/json" ||
		r.Header.Get("Request-From") == "x-web-json" {
		w.Header().Set("Content-Type", "application/json")
		common.Logs.Debug("resp: " + string(JsonResult.Get(code)))
		w.Write(JsonResult.Get(code))
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
		common.Logs.Debug("resp: " + string(JsonResult.Get(code)))
		w.Write(PbResult.Get(code))
	}
}

///*
// *@note WritePbRespone http服务器针对已序列化的proto.message数据的回包，
// *@note 根据客户端的Content-Type判断回json格式还是proto格式
// *@param w http.ResponseWriter
// *@param r *http.Request
// *@param data 已序列化的proto.message数据
// */
//func WritePbOctRespone(w http.ResponseWriter, r *http.Request, data []byte) {
//	if r.Header.Get("Content-Type") == "application/json" {
//		w.Header().Set("Content-Type", "application/json")
//		var info proto.Message
//		PBUnMarshal(data, info)
//		reply, _ := json.Marshal(info)
//		common.Logs.Debug(string(reply))
//		w.Write(reply)
//	} else {
//		var info proto.Message
//		PBUnMarshal(data, info)
//		reply, _ := json.Marshal(info)
//		common.Logs.Debug(string(reply))

//		w.Header().Set("Content-Type", "application/octet-stream")
//		w.Write(data)
//	}
//}

/*
*@note 结构转Map
*@param obj 结构体
*@return map[string]interface{}
 */
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

/*
*@note 结构转[]byte
*@param v 结构体
*@return []byte
 */
func Struct2Byte(v interface{}) []byte {
	reply, err := json.Marshal(v)
	if err != nil {
		common.Logs.Debug("Struct2Byte(%v) fail, err=%s", v, err.Error())
		return nil
	}

	return reply
}

/*
*@note 结构转字符串
*@param v 结构体
*@return stirng
 */
func Struct2String(v interface{}) string {
	reply, err := json.Marshal(v)
	if err != nil {
		common.Logs.Debug("Struct2String(%v) fail, err=%s", v, err.Error())
		return ""
	}

	return string(reply)
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
		return nil, Err_Com_Marshal
	}

	return data, No_Error
}

func Byte2Proto(data []byte, pb proto.Message) int32 {
	err := proto.Unmarshal(data, pb)
	if err != nil {
		common.Logs.Warn("Byte2Proto() fail, err=%s", err.Error())
		return Err_Com_UnMarshal
	}

	return No_Error
}

func ReverseStrings(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
*@note 测试函数执行时间
*@param funName 函数名
*@param start 开始时间
 */
func TestTimeout(funName string, start time.Time) {
	dis := time.Now().Sub(start).Seconds()
	common.Logs.Warn("funName:%s, timeout:%fs", funName, dis)
}

/*
*@note 统计长度
*@param s 内容
*@return 内容长度
 */
func StatsLen(s string) int32 {
	var count int32

	for _, c := range []rune(s) {
		if (c >= 0 && c <= 255) || (c >= 0xff61 && c <= 0xff9f) {
			count = count + 1
		} else {
			count = count + 2
		}
	}

	return count
}

var (
	jsonMarshaler = jsonpb.Marshaler{}
)

func SimpleJson(v interface{}) string {
	var err error
	var jsonBytes []byte
	if pb, ok := v.(proto.Message); ok {
		var buf bytes.Buffer
		err = jsonMarshaler.Marshal(&buf, pb)
		jsonBytes = buf.Bytes()
	} else {
		jsonBytes, err = json.Marshal(v)
	}
	if err != nil {
		common.Logs.Warn("SimpleJson: marshal error: %v", err)
		return ""
	}
	return *(*string)(unsafe.Pointer(&jsonBytes))
}
