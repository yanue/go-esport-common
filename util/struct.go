/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : struct.go
 Time    : 2018/9/19 14:16
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"encoding/json"
	"github.com/yanue/go-esport-common"
	"reflect"
)

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
