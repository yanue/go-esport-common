/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : init.go
 Time    : 2018/9/18 11:48
 Author  : yanue
 
 - 各种验证
 - 用法:
	validator.Verify.IsXxx()

------------------------------- go ---------------------------------*/

package validator

type validator struct {
	*cRegexpName
	*cRegexpPhoneCode
}

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
