/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : account.go
 Time    : 2018/9/18 17:53
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package errcode

var errMsgAccount = map[int32]string{
	// account service
	Err_Account_Old_Psw:                       "旧密码校验失败",
	Err_Account_SetUserEx:                     "设置用户信息失败",
	Err_Account_Binding_Get:                   "获取账户绑定信息失败",
	Err_Account_Bind_Phone:                    "绑定手机账号失败",
	Err_Account_Bind_WeChat:                   "绑定微信账号失败",
	Err_Account_Bind_WeiBo:                    "绑定微博账号失败",
	Err_Account_Bind_Facebook:                 "绑定脸书账号失败",
	Err_Account_Bind_Google:                   "绑定谷歌账号失败",
	Err_Account_GetLevel:                      "获取会员等级失败",
	Err_Account_UserInfo_NotExist:             "用户信息不存在",
	Err_Account_Blacklist_Add:                 "添加黑名单失败",
	Err_Account_Blacklist_Update:              "更新黑名单失败",
	Err_Account_Blacklist_Del:                 "删除黑名单失败",
	Err_Account_Blacklist_Get:                 "获取黑名单失败",
	Err_Account_Blacklist_Count:               "获取黑名单数量失败",
	Err_Account_Nickname_Sensitive_Word:       "昵称包含敏感词",
	Err_Account_Realname_Sensitive_Word:       "真实名称包含敏感词",
	Err_Account_UserResume_Sensitive_Word:     "备注包含敏感词",
	Err_Account_AuthToken:                     "验证令牌失败",
	Err_Account_AuthPwd:                       "验证密码失败",
	Err_Account_ReadNameAuthChecking:          "实名认证已在审核中",
	Err_Account_ReadNameAuthOk:                "实名认证已认证",
	Err_Account_CheckUrl:                      "超链接格式错误",
	Err_Account_Username_Len:                  "仅支持6-12位字符",
	Err_Account_Nickname_Len:                  "仅支持3-24位字符",
	Err_Account_Realname_Len:                  "真名长度错误",
	Err_Account_Username_Fmt:                  "仅支持数字、字母、下划线，且必须以字母开头",
	Err_Account_Nickname_Fmt:                  "昵称中包含不被支持的字符",
	Err_Account_Realname_Fmt:                  "真名格式错误",
	Err_Account_UserResume_Fmt:                "个人简介错误",
	Err_Account_IdCard_Fmt:                    "身份证格式错误",
	Err_Account_Birthday_Fmt:                  "生日格式错误",
	Err_Account_Above_Vip:                     "当前等级超过VIP等级",
	Err_Account_AreaCode_Fmt:                  "区域代码格式错误",
	Err_Account_Username_Sensitive_Word:       "BB号包含敏感词",
	Err_Account_IdCardNo_Exist:                "提交身份证号码重复",
	Err_Account_Email_Fmt:                     "邮箱账号格式错误",
	Err_Account_Check_IdCard_Date:             "请校验身份证日期",
	Err_Account_Vo_Member:                     "提升权限验证失败",
	Err_Account_Position_Sensitive_Word:       "职位包含敏感词",
	Err_Account_Level_Not_Vo:                  "请求用户非企业用户",
	Err_Account_Level_Not_Vipp:                "请求用户非认证用户",
	Err_Account_Change_Vipp_Level:             "当前等级不支持直接修改",
	Err_Account_User_Disabled:                 "该用户已被封号",
	Err_Account_Nickname_Conflict:             "该昵称已被占用",
	Err_Account_Nickname_Change_Limited:       "昵称修改次数已达上限",
	Err_Account_MedalName_Exists:              "勋章名称已存在",
	Err_Account_MedalFormat:                   "勋章格式错误",
	Err_Account_SystemMedal:                   "系统配置勋章不可删除",
	Err_Account_MedalRelation:                 "未获得这个勋章，无法进行佩戴",
	Err_Account_Medal_NotExists:               "勋章不存在",
	Err_Account_Medal_Duplicated:              "勋章不可重复佩戴",
	Err_Account_Certificate_NotExists:         "认证信息不存在",
	Err_Account_Certificate_Duplicated:        "请不要重复认证",
	Err_Account_AuthName_Sensitive_Word:       "认证信息包含敏感词",
	Err_Account_AdditionalInfo_Sensitive_Word: "额外信息包含敏感词",
	Err_Account_Job_Sensitive_Word:            "从事职业包含敏感词",
	Err_Account_Introduce_Sensitive_Word:      "自我介绍包含敏感词",
	Err_Account_Certificate_Type_Mismatch:     "认证条件不满足",
	Err_Account_Modify_Certificate_Condition:  "个人认证属性，不支持修改",
	Err_Account_Not_TeamMember:                "该帐号没有队员权限",
	Err_Account_Change_Nick_Once_PerMonth:     "每个月最多可修改一次昵称",
}

// account service err code
const (
	Err_Account_Old_Psw = err_Offset_Account + iota
	Err_Account_SetUserEx
	Err_Account_Binding_Get
	Err_Account_Bind_Phone
	Err_Account_Bind_WeChat
	Err_Account_Bind_WeiBo
	Err_Account_Bind_Facebook
	Err_Account_Bind_Google
	Err_Account_GetLevel
	Err_Account_UserInfo_NotExist
	Err_Account_Blacklist_Add
	Err_Account_Blacklist_Update
	Err_Account_Blacklist_Del
	Err_Account_Blacklist_Get
	Err_Account_Blacklist_Count
	Err_Account_Nickname_Sensitive_Word
	Err_Account_Realname_Sensitive_Word
	Err_Account_UserResume_Sensitive_Word
	Err_Account_AuthToken
	Err_Account_AuthPwd
	Err_Account_ReadNameAuthChecking
	Err_Account_ReadNameAuthOk
	Err_Account_CheckUrl
	Err_Account_Username_Len
	Err_Account_Nickname_Len
	Err_Account_Realname_Len
	Err_Account_Username_Fmt
	Err_Account_Nickname_Fmt
	Err_Account_Realname_Fmt
	Err_Account_UserResume_Fmt
	Err_Account_IdCard_Fmt
	Err_Account_Birthday_Fmt
	Err_Account_Above_Vip
	Err_Account_AreaCode_Fmt
	Err_Account_Username_Sensitive_Word
	Err_Account_IdCardNo_Exist
	Err_Account_Email_Fmt
	Err_Account_Check_IdCard_Date
	Err_Account_Vo_Member
	Err_Account_Position_Sensitive_Word
	Err_Account_Level_Not_Vo
	Err_Account_Level_Not_Vipp
	Err_Account_Change_Vipp_Level
	Err_Account_User_Disabled
	Err_Account_Nickname_Conflict
	Err_Account_Nickname_Change_Limited
	Err_Account_MedalName_Exists
	Err_Account_MedalFormat
	Err_Account_SystemMedal
	Err_Account_MedalRelation
	Err_Account_Medal_NotExists
	Err_Account_Medal_Duplicated
	Err_Account_Certificate_NotExists
	Err_Account_Certificate_Duplicated
	Err_Account_AuthName_Sensitive_Word
	Err_Account_AdditionalInfo_Sensitive_Word
	Err_Account_Job_Sensitive_Word
	Err_Account_Introduce_Sensitive_Word
	Err_Account_Certificate_Type_Mismatch
	Err_Account_Modify_Certificate_Condition
	Err_Account_Not_TeamMember
	Err_Account_Change_Nick_Once_PerMonth
)
