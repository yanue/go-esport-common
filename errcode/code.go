/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : init.go
 Time    : 2018/9/18 12:06
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package errcode

import "errors"

type ErrNo = int32

const (
	No_Error           ErrNo = 0
	err_Offset_Common  ErrNo = 1001000
	err_Offset_InValid ErrNo = 1101000
	err_Offset_Account ErrNo = 1201000
	err_Offset_Login   ErrNo = 1301000
)

var errMsgAll = map[ErrNo]map[int32]string{
	err_Offset_Common:  errMsgCommon,  // 公共错误码
	err_Offset_InValid: errMsgInValid, // 无效相关的错误码
	err_Offset_Account: errMsgAccount, // 账号相关错误码
	err_Offset_Login:   errMsgLogin,   // 登陆相关
}

func init() {
	// 初始化合并错误信息

}

func GetErrStr(code int32) string {
	var errMsg map[int32]string
	// 查找对应区间
	for errOffset, msgStep := range errMsgAll {
		// offset 同一区间间隔最大1000
		if code%errOffset < 1000 {
			// 找到了对应区间
			errMsg = msgStep
		}
	}

	reply, ok := errMsg[code]
	if ok {
		return reply
	} else {
		reply, _ := errMsg[Err_Com_Unknown]
		return reply
	}
}

func GetError(code int32) error {
	var errMsg map[int32]string
	// 查找对应区间
	for errOffset, msgStep := range errMsgAll {
		// offset 同一区间间隔最大1000
		if code%errOffset < 1000 {
			// 找到了对应区间
			errMsg = msgStep
		}
	}

	reply, ok := errMsg[code]
	if !ok {
		resp, _ := errMsg[Err_Com_Unknown]
		reply = resp
	}

	return errors.New(reply)
}
