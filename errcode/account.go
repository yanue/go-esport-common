/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : account.go
 Time    : 2018/9/18 17:53
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package errcode

var errMsgAccount = map[int32]string{
	// account service
	ErrAccountNotExist:                   "账号不存在",
	ErrAccountGetUserInfo:                "获取用户信息失败",
	ErrAccountExist:                      "账号已存在",
	ErrAccountDisabled:                   "该用户已被封号",
	ErrAccountPassIncorrect:              "密码不正确",
	ErrAccountPassNotSet:                 "你还未完善密码",
	ErrAccountSetUserProfile:             "设置用户信息失败",
	ErrAccountSensitiveNick:              "昵称包含敏感词",
	ErrAccountSensitiveRealname:          "真实名称包含敏感词",
	ErrAccountSensitiveAccount:           "账号包含敏感词",
	ErrAccountLenAccount:                 "仅支持6-15位字符",
	ErrAccountLenNickname:                "仅支持3-24位字符",
	ErrAccountLenRealname:                "真名长度错误",
	ErrAccountLenPass:                    "密码长度错误",
	ErrAccountFmtAccount:                 "仅支持数字、字母、下划线，且必须以字母开头",
	ErrAccountFmtNickname:                "昵称中包含不被支持的字符",
	ErrAccountFmtRealname:                "真名格式错误",
	ErrAccountAuthtoken:                  "验证令牌失败",
	ErrAccountAuthpwd:                    "验证密码失败",
	ErrAccountNicknameConflict:           "该昵称已被占用",
	ErrAccountNicknameChangeLimited:      "昵称修改次数已达上限",
	ErrAccountNicknameChangeOncePermonth: "每个月最多可修改一次昵称",
	ErrAccountVerificationChecking:       "实名认证审核中",
	ErrAccountVerificationOk:             "实名认证已认证",
	// login service
	ErrAccountPhoneExist:          "手机账号已存在",
	ErrAccountPhoneNotExist:       "手机账号不存在",
	ErrAccountPhoneAreaNotSupport: "不可以用该区号的手机号码注册",
	ErrAccountPhoneInBlacklist:    "该手机号码不能注册",
	// 验证码
	ErrVerifyCodeRegAlreadyExist:  "该手机号码已经注册过，请使用其他手机号码注册",
	ErrVerifyCodeResetPwdNotExist: "没有查询到该手机号码的注册信息，请核对您绑定的手机号码",
	ErrVerifyCodeBindAlreadyExist: "该手机号码已经注册过，请使用其他手机号码绑定",
	// 第三方登陆
	ErrAccountLoginType:         "请选择登陆方式",
	ErrAccountBindCancelOnlyOne: "您尚未绑定其它登录方式，无法取消当前绑定方式",
	ErrAccountBindGet:           "获取账户绑定信息失败",
	ErrAccountBindType:          "不支持该类型的绑定方式",
	ErrAccountBindPhone:         "绑定手机账号失败",
	ErrAccountBindWechat:        "绑定微信账号失败",
	ErrAccountBindQQ:            "绑定QQ账号失败",
	ErrAccountBindWeibo:         "绑定微博账号失败",
	ErrAccountBindExist:         "该登录方式已绑定其他账号",
	ErrAccountBindExistPhone:    "该手机号已绑定其他账号",
	ErrAccountBindNotExist:      "用户没有绑定该登录方式",
	ErrAccountVerifyWechat:      "微信验证失败",
	ErrAccountVerifyQQ:          "QQ验证失败",
	ErrAccountTokenGenerate:     "生成token失败,请重试!",
	ErrAccountTokenUpdate:       "登录信息授权失败，您需要重新登录",
	ErrAccountTokenFormat:       "登录信息格式错误，您需要重新登录",
	ErrAccountTokenGet:          "获取登陆信息失败，您需要重新登录",
	ErrAccountTokenExpire:       "登录信息已过期，您需要重新获取授权",
	ErrAccountTokenRefresh:      "登录信息已过期，您需要重新登录", // 第三方登陆
	ErrAccountTokenNotEqual:     "您的账号已在其它设备上登录，请重新登录",
	ErrAccountTokenVerify:       "登录信息已过期，您需要重新登录",
}

// account service err code
const (
	// account service
	ErrAccountNotExist                   = errOffsetAccount + iota //账号不存在",
	ErrAccountGetUserInfo                                          //获取用户信息失败",
	ErrAccountExist                                                //账号已存在",
	ErrAccountDisabled                                             //该用户已被封号",
	ErrAccountPassIncorrect                                        //密码不正确",
	ErrAccountPassNotSet                                           //你还未完善密码
	ErrAccountSetUserProfile                                       //设置用户信息失败",
	ErrAccountSensitiveNick                                        //昵称包含敏感词",
	ErrAccountSensitiveRealname                                    //真实名称包含敏感词",
	ErrAccountSensitiveAccount                                     //账号包含敏感词",
	ErrAccountLenAccount                                           //仅支持6-15位字符",
	ErrAccountLenNickname                                          //仅支持3-24位字符",
	ErrAccountLenRealname                                          //真名长度错误",
	ErrAccountLenPass                                              //密码长度错误",
	ErrAccountFmtAccount                                           //仅支持数字、字母、下划线，且必须以字母开头",
	ErrAccountFmtNickname                                          //昵称中包含不被支持的字符",
	ErrAccountFmtRealname                                          //真名格式错误",
	ErrAccountAuthtoken                                            //验证令牌失败",
	ErrAccountAuthpwd                                              //验证密码失败",
	ErrAccountNicknameConflict                                     //该昵称已被占用",
	ErrAccountNicknameChangeLimited                                //昵称修改次数已达上限",
	ErrAccountNicknameChangeOncePermonth                           //"每个月最多可修改一次昵称",
	ErrAccountVerificationChecking                                 //实名认证审核中",
	ErrAccountVerificationOk                                       //实名认证已认证",
	// login service
	ErrAccountPhoneExist           //手机账号已存在",
	ErrAccountPhoneNotExist        //手机账号不存在",
	ErrAccountPhoneAreaNotSupport  //不可以用该区号的手机号码注册",
	ErrAccountPhoneInBlacklist     //该手机号码不能注册",
	// 验证码
	ErrVerifyCodeRegAlreadyExist   //该手机号码已经注册过，请使用其他手机号码注册",
	ErrVerifyCodeResetPwdNotExist  //没有查询到该手机号码的注册信息，请核对您绑定的手机号码",
	ErrVerifyCodeBindAlreadyExist  //该手机号码已经注册过，请使用其他手机号码绑定",
	// 第三方登陆
	ErrAccountLoginType          //请选择登陆方式",
	ErrAccountBindCancelOnlyOne  //您尚未绑定其它登录方式，无法取消当前绑定方式",
	ErrAccountBindGet            //获取账户绑定信息失败",
	ErrAccountBindType           //不支持该类型的绑定方式",
	ErrAccountBindPhone          //绑定手机账号失败",
	ErrAccountBindWechat         //绑定微信账号失败",
	ErrAccountBindQQ             //绑定QQ账号失败",
	ErrAccountBindWeibo          //绑定微博账号失败",
	ErrAccountBindExist          //该登录方式已绑定其他账号",
	ErrAccountBindExistPhone     //该手机号已绑定其他账号",
	ErrAccountBindNotExist       //用户没有绑定该登录方式",
	ErrAccountVerifyWechat       //微信验证失败",
	ErrAccountVerifyQQ           //qq验证失败",
	ErrAccountTokenGenerate      //登录信息授权失败，您需要重新登录",
	ErrAccountTokenUpdate        //登录信息授权失败，您需要重新登录",
	ErrAccountTokenFormat        //登录信息格式错误，您需要重新登录",
	ErrAccountTokenGet           //获取登陆信息失败，您需要重新登录,
	ErrAccountTokenExpire        //登录信息已过期，您需要重新获取授权",
	ErrAccountTokenRefresh       //登录信息已过期，您需要重新登录",
	ErrAccountTokenNotEqual      //您的账号已在其它设备上登录，请重新登录
	ErrAccountTokenVerify        //登录信息已过期，您需要重新登录,
)
