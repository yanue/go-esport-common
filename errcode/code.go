/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : init.go
 Time    : 2018/9/18 12:06
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package errcode

import (
	"errors"
	"fmt"
	"os"
)

type ErrNo = int32

const (
	No_Error         ErrNo = 0
	errOffsetCommon  ErrNo = 1001000
	errOffsetInvalid ErrNo = 1101000
	errOffsetSms     ErrNo = 1201000
	errOffsetAccount ErrNo = 2101000
)

var errMsgAll = map[ErrNo]map[int32]string{
	errOffsetCommon:  errMsgCommon,  // 公共错误码
	errOffsetInvalid: errMsgInValid, // 无效相关的错误码
	errOffsetSms:     errMsgSms,     // 短信消息
	errOffsetAccount: errMsgAccount, // 账号相关错误码
}

func init() {
	// 初始化合并错误信息

}

/*
*@note 根据错误码获取错误信息
*@param code 错误码
*@param extraMsg 额外自定义信息
*@return 对应错误信息
*/
func GetErrMsg(code int32, extraMsg ... string) string {
	if code == 0 {
		return errMsgCommon[code]
	}

	var errMsg map[int32]string
	// 查找对应区间
	for errOffset, msgStep := range errMsgAll {
		// offset 同一区间间隔最大1000
		if code%errOffset < 1000 {
			// 找到了对应区间
			errMsg = msgStep
			break
		}
	}

	reply, ok := errMsg[code]
	if !ok {
		resp, _ := errMsg[ErrCommonUnknownError]
		reply = resp
	}

	if len(extraMsg) > 0 {
		if len(reply) > 0 {
			reply = reply + "[" + extraMsg[0] + "]"
		} else {
			reply = extraMsg[0]
		}
	}

	return reply
}

/*
*@note 根据错误码获取错误
*@param code 错误码
*@param extraMsg 额外自定义信息
*@return
*/
func GetError(code int32, extraMsg ... string) error {
	return errors.New(GetErrMsg(code, extraMsg...))
}

/*
*@note 转换错误码到json文件
*@param code 错误码
*@param extraMsg 额外自定义信息
*@return
*/
func ConvertJsonFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("//自动生成，不要手动修改！\n")
	file.WriteString("var Result = {\n")
	for _, msgStep := range errMsgAll {
		for code, msg := range msgStep {
			fmt.Println("code, msg", fmt.Sprintf("\t%d: %q,\n", code, msg))
			file.WriteString(fmt.Sprintf("\t%d: %q,\n", code, msg))
		}
	}
	file.WriteString("};")

	return nil
}
