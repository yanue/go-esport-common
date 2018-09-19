/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : init.go
 Time    : 2018/9/18 11:48
 Author  : yanue
 
 - 各种验证
 - 用法:
	validator.Verify.IsXxx()

------------------------------- go ---------------------------------*/

package validator

// 长度定义
const (
	LenMinAccount  = 3  // 最短用户名长度
	LenMinNickname = 2  // 最短昵称长度
	LenMinRealName = 2  // 最短实名长度
	LenMinPassword = 6  // 最短密码长度
	LenMaxAccount  = 20 // 最长用户名长度
	LenMaxNickname = 15 // 最长昵称长度
	LenMaxRealName = 20 // 最长实名长度
	LenMaxPassword = 20 // 最长密码长度
)

// 合并验证
type validator struct {
	*cRegexpName
	*cRegexpPhoneCode
}

// 验证器
var Verify *validator

func init() {
	Verify = new(validator)

	// 手机号相关验证
	Verify.cRegexpPhoneCode = &cRegexpPhoneCode{
		regexpWithAreaCode,
		regexpWithoutAreaCode,
	}

	// 名字相关验证
	Verify.cRegexpName = &cRegexpName{
		regexpAccount:  regexpAccount,
		regexpNickname: regexpNickname,
		regexpRealName: regexpRealName,
	}
}
