/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : common.go
 Time    : 2018/9/18 17:52
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package errcode

// invalid err code
const (
	Err_InValid_Custom   = err_Offset_InValid + iota // 自定义参数
	Err_InValid_Param                                // 无效的参数
	Err_InValid_Data                                 // 无效的数据
	Err_InValid_Account                              // 无效的账号
	Err_InValid_Phone                                // 无效的手机号
	Err_InValid_Email                                // 无效的邮箱
	Err_InValid_Url                                  // 无效的网址
	Err_InValid_Img                                  // 无效的图片
	Err_InValid_Sign                                 // 无效的签名
	Err_InValid_Date                                 // 无效的日期
	Err_InValid_IdCardNo                             // 无效的身份证号
)

var errMsgInValid = map[int32]string{
	Err_InValid_Custom:   "自定义参数",
	Err_InValid_Param:    "无效的参数",
	Err_InValid_Data:     "无效的数据",
	Err_InValid_Account:  "无效的账号",
	Err_InValid_Phone:    "无效的手机号",
	Err_InValid_Email:    "无效的邮箱",
	Err_InValid_Url:      "无效的网址",
	Err_InValid_Img:      "无效的图片",
	Err_InValid_Sign:     "无效的签名",
	Err_InValid_Date:     "无效的日期",
	Err_InValid_IdCardNo: "无效的身份证号",
}

// err msg
var errMsgCommon = map[int32]string{
	No_Error:                       "操作成功",
	Err_Com_Unknown:                "未知错误",
	Err_Com_Redis:                  "获取用户数据失败，请稍后重试",
	Err_Com_Request_Method:         "HTTP请求类型错误",
	Err_Com_URL_Param:              "参数有误",
	Err_Com_Inside_Service_Verify:  "服务校验失败",
	Err_Com_Token_Verify:           "登录信息已过期，您需要重新登录",
	Err_Com_Net_ReadData:           "读取网络数据失败",
	Err_Com_Upload_File:            "上传文件失败",
	Err_Com_Marshal:                "网络传输失败，请稍后重试",
	Err_Com_UnMarshal:              "网络传输失败，请稍后重试",
	Err_Com_Set_Consul:             "修改配置信息失败",
	Err_Com_Cookies_Param:          "cookies参数错误",
	Err_Com_RemoteCall:             "服务通信失败，请稍后重试",
	Err_Com_Email_Format:           "邮箱名格式错误",
	Err_Com_Mobile_Or_Email_Format: "手机号码或邮箱名格式错误",
	Err_Com_Keyword_Len:            "关键字长度错误",
	Err_Com_Too_Many_Param_Value:   "参数值太多",
	Err_Com_Req_Frequent:           "请求太频繁，请稍后再试",
	Err_Com_Rds:                    "获取数据失败，请稍后再试",
	Err_Com_ForceBindPhoneCode:     "系统账户升级，你需要重新登录后绑定手机号才能继续使用APP",
	Err_Com_Betch_Task_Failed:      "批量任务处理失败",
	Err_Com_Ots:                    "获取数据失败，请稍后再试",
	Err_Com_Unknown_Request:        "请求无法处理",
}

// common err code
const (
	Err_Com_Unknown = err_Offset_Common + iota
	Err_Com_Redis
	Err_Com_Request_Method
	Err_Com_URL_Param
	Err_Com_Inside_Service_Verify
	Err_Com_Token_Verify
	Err_Com_Net_ReadData
	Err_Com_Upload_File
	Err_Com_Marshal
	Err_Com_UnMarshal
	Err_Com_Set_Consul
	Err_Com_RemoteCall
	Err_Com_Cookies_Param
	Err_Com_Email_Format
	Err_Com_Mobile_Or_Email_Format
	Err_Com_Keyword_Len
	Err_Com_Too_Many_Param_Value  // "参数值太多",
	Err_Com_Req_Frequent          // "请求太频繁，请稍后再试",
	Err_Com_Rds                   // "获取数据失败，请稍后再试",
	Err_Com_ForceBindPhoneCode    // "系统账户升级，你需要重新登录后绑定手机号才能继续使用APP",
	Err_Com_Betch_Task_Failed     // "批量任务处理失败",
	Err_Com_Ots                   // "获取数据失败，请稍后再试",
	Err_Com_Unknown_Request       // "请求无法处理",
)
