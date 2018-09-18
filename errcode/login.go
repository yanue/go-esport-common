/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : account.go
 Time    : 2018/9/18 17:53
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package errcode

var errMsgLogin = map[int32]string{
	// login service
	Err_Login_Type:                         "请使用其他第三方登录",
	Err_Login_Account_Psw:                  "账号或密码不匹配",
	Err_Login_Account_Len:                  "仅支持6-12位字符",
	Err_Login_Nickname_Len:                 "昵称长度错误",
	Err_Login_Realname_Len:                 "真名长度错误",
	Err_Login_Account_Fmt:                  "仅支持数字、字母、下划线，且必须以字母开头",
	Err_Login_Nickname_Fmt:                 "昵称格式错误",
	Err_Login_Realname_Fmt:                 "真名格式错误",
	Err_Login_Password_Len:                 "密码长度错误",
	Err_Login_PhoneCode_Exists:             "手机账号已存在",
	Err_Login_PhoneCode_Not_Exists:         "手机账号不存在",
	Err_Login_PhoneCode_Area:               "不可以用该区号的手机号码注册",
	Err_Login_PhoneCode_Format:             "手机号码格式错误",
	Err_Login_PhoneCode_InBlacklist:        "该手机号码不能注册",
	Err_Login_Username_Exists:              "BB号已存在",
	Err_Login_Account_Exist:                "用户已存在",
	Err_Login_Account_Not_Exist:            "用户不存在",
	Err_Login_Account_Has_Keyword:          "您输入的BB号含有敏感词，请更换BB号",
	Err_Login_Account_Bind_Type:            "不支持该类型的绑定方式",
	Err_Login_Account_Bind_Exist:           "用户已绑定该登录方式",
	Err_Login_Account_Bind_Not_Exist:       "用户没有绑定该登录方式",
	Err_Login_Account_Disable:              "该账号已经被封，请联系客服代表",
	Err_Login_AppKey_NotExist:              "应用信息有误，请联系客服代表",
	Err_Login_Verify_WeChat:                "微信验证失败",
	Err_Login_Verify_WeiBo:                 "微博验证失败",
	Err_Login_Verify_Facebook:              "脸书验证失败",
	Err_Login_Verify_Google:                "谷歌验证失败",
	Err_Login_Bind_PUID_Exists:             "该登录方式已绑定其他账号，无法再绑定",
	Err_Login_Token_Update:                 "登录信息授权失败，您需要重新登录",
	Err_Loing_Token_Format:                 "登录信息格式错误，您需要重新登录",
	Err_Login_Token_UserID:                 "登录信息与用户信息不匹配，您需要重新登录",
	Err_Login_Token_Expire:                 "登录信息已过期，您需要重新获取授权",
	Err_Login_Token_Reflash:                "登录信息已过期，您需要重新登录",
	Err_Login_Token_Verify:                 "您的账号已在其它设备上登录，请重新登录",
	Err_Login_Token_Inside_Get:             "只有边缘服务器才能获取token",
	Err_Login_AuthCode_Reg_Already_Exist:   "该手机号码已经注册过，请使用其他手机号码注册",
	Err_Login_AuthCode_Reset_Pwd_Not_Exist: "没有查询到该手机号码的注册信息，请核对您绑定的手机号码",
	Err_Login_AuthCode_Bind_Already_Exist:  "该手机号码已经注册过，请使用其他手机号码绑定",
	Err_Login_AuthCode_Verify:              "短信验证码验证失败",
	Err_Login_Register_Sign:                "签名验证失败",
	Err_Login_UnBind_Only_One:              "您尚未绑定其它登录方式，无法取消当前绑定方式",
	Err_Login_Reset_Wrong_PhoneCode:        "您输入的手机号码已绑定其他账号",
}

// login service err code
const (
	Err_Login_Type                         = err_Offset_Login + iota // "请使用其他第三方登录",
	Err_Login_Account_Psw                                            // "账号或密码不匹配",
	Err_Login_Account_Len                                            // "仅支持6-12位字符",
	Err_Login_Nickname_Len                                           // "昵称长度错误",
	Err_Login_Realname_Len                                           // "真名长度错误",
	Err_Login_Account_Fmt                                            // "仅支持数字、字母、下划线，且必须以字母开头",
	Err_Login_Nickname_Fmt                                           // "昵称格式错误",
	Err_Login_Realname_Fmt                                           // "真名格式错误",
	Err_Login_Password_Len                                           // "密码长度错误",
	Err_Login_PhoneCode_Exists                                       // "手机账号已存在",
	Err_Login_PhoneCode_Not_Exists                                   // "手机账号不存在",
	Err_Login_PhoneCode_Area                                         // "不可以用该区号的手机号码注册",
	Err_Login_PhoneCode_Format                                       // "手机号码格式错误",
	Err_Login_PhoneCode_InBlacklist                                  // "该手机号码不能注册",
	Err_Login_Username_Exists                                        // "BB号已存在",
	Err_Login_Account_Exist                                          // "用户已存在",
	Err_Login_Account_Not_Exist                                      // "用户不存在",
	Err_Login_Account_Has_Keyword                                    // "您输入的BB号含有敏感词，请更换BB号",
	Err_Login_Account_Bind_Type                                      // "不支持该类型的绑定方式",
	Err_Login_Account_Bind_Exist                                     // "用户已绑定该登陆方式",
	Err_Login_Account_Bind_Not_Exist                                 // "用户没有绑定该登陆方式",
	Err_Login_Account_Disable                                        // "该账号已经被封，请联系客服代表",
	Err_Login_AppKey_NotExist                                        // "应用信息有误，请联系客服代表",
	Err_Login_Verify_WeChat                                          // "微信验证失败",
	Err_Login_Verify_WeiBo                                           // "微博验证失败",
	Err_Login_Verify_Facebook                                        // "脸书验证失败",
	Err_Login_Verify_Google                                          // "谷歌验证失败",
	Err_Login_Bind_PUID_Exists                                       // "该登陆方式已绑定其他账号，无法再绑定",
	Err_Login_Create_UUID                                            // "创建用户失败",
	Err_Login_Create_ChatID                                          // "创建聊天信息失败",
	Err_Login_Token_Update                                           // "登陆信息授权失败，您需要重新登陆",
	Err_Loing_Token_Format                                           // "登陆信息格式错误，您需要重新登陆",
	Err_Login_Token_UserID                                           // "登陆信息与用户信息不匹配，您需要重新登陆",
	Err_Login_Token_Expire                                           // "登陆信息已过期，您需要重新获取授权",
	Err_Login_Token_Reflash                                          // "登录信息已过期，您需要重新登录",
	Err_Login_Token_Verify                                           // "您的账号已在其它设备上登陆，请重新登陆",
	Err_Login_Token_Inside_Get                                       // "只有边缘服务器才能获取token",
	Err_Login_AuthCode_Reg_Already_Exist                             // "该手机号码已经注册过，请使用其他手机号码注册",
	Err_Login_AuthCode_Reset_Pwd_Not_Exist                           // "没有查询到该手机号码的注册信息，请核对您绑定的手机号码",
	Err_Login_AuthCode_Bind_Already_Exist                            // "该手机号码已经注册过，请使用其他手机号码绑定",
	Err_Login_AuthCode_Verify                                        // "短信验证码验证失败",
	Err_Login_Register_Sign                                          // "签名验证失败",
	Err_Login_UnBind_Only_One                                        // "您尚未绑定其它登陆方式，无法取消当前绑定方式",
	Err_Login_Reset_Wrong_PhoneCode                                  // "您输入的手机号码已绑定其他账号",
)
