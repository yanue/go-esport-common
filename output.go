/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : output.go
 Time    : 2018/9/18 11:05
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package common

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/yanue/go-esport-common/errcode"
	pb "github.com/yanue/go-esport-common/proto"
	"reflect"
	"strings"
)

var Output *result = new(result)

type result struct {
}

/*
 *@note  对消息进行统一格式封装
 *@param pb 消息结构
 *@return 字节数组或error
 */
func (this result) ProtoRight() {

}

/*
 *@note  对消息进行统一格式封装
 *@param pb 消息结构
 *@return 字节数组或error
 */
func (this result) ProtoError(code errcode.ErrNo) {
}

/*
 *@note  对消息进行统一格式封装
 *@param pb 消息结构
 *@return 字节数组或error
 */
func (this result) JsonRight() {

}

/*
 *@note  对消息进行统一格式封装
 *@param pb 消息结构
 *@return 字节数组或error
 */
func (this result) JsonError(code int32) {
	msg := errcode.GetErrStr(code)
	a := pb.PJsonResult{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

/*
 *@note  对消息进行统一格式封装
 *@param pb 消息结构
 *@return 字节数组或error
 */
func PBMarshal(pb proto.Message) ([]byte, error) {
	buf, _ := proto.Marshal(pb)
	tpName := reflect.TypeOf(pb).String()
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
func JsonMarshal(pb proto.Message) ([]byte, error) {
	buf, _ := proto.Marshal(pb)
	tpName := reflect.TypeOf(pb).String()
	pos := strings.LastIndex(tpName, ".")

	msg := &pb.PMessage{
		Magic: 9833,
		Type:  tpName[pos+1:],
		Data:  buf,
	}

	return json.Marshal(msg)
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

	Logs.Debug("msg:%+v", msg)

	pm := reflect.New(proto.MessageType("bbproto." + msg.Type).Elem()).Interface().(proto.Message)
	err = proto.Unmarshal(msg.Data, pm)
	if err != nil {
		return nil, err
	}

	return json.Marshal(pm)
}
