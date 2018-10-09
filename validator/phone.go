/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : phone.go
 Time    : 2018/9/14 17:20
 Author  : yanue
 
 - 手机校验相关
 
------------------------------- go ---------------------------------*/

package validator

import (
	"github.com/yanue/go-esport-common/errcode"
	"regexp"
	"strings"
)

type cRegexpPhoneCode struct {
	regexpWithAreaCode    map[string]*regexp.Regexp
	regexpWithoutAreaCode map[string]*regexp.Regexp
}

const (
	AreaCode_CN = "+86"
	AreaCode_HK = "+852"
	AreaCode_MO = "+853"
	AreaCode_TW = "+886"
	AreaCode_US = "+1"
)

// 手机号验证规则
var regexpWithAreaCode = map[string]*regexp.Regexp{
	AreaCode_CN: regexp.MustCompile(`^(\+86)((1[1-9][0-9])|(14[5,7])|(17[0-9]))[0-9]{8}$`), //中国
	AreaCode_HK: regexp.MustCompile(`^(\+852)(9|6|5)[0-9]{7}$`),                            //香港
	AreaCode_MO: regexp.MustCompile(`^(\+853)(66|68)[0-9]{5}$`),                            //澳门
	AreaCode_TW: regexp.MustCompile(`^(\+886)9[0-9]{8}$`),                                  //台湾
	AreaCode_US: regexp.MustCompile(`^(\+1)[0-9]{10}$`),                                    //美国
}

// 手机号验证规则
var regexpWithoutAreaCode = map[string]*regexp.Regexp{
	AreaCode_CN: regexp.MustCompile(`^(\+86)?((1[1-9][0-9])|(14[5,7])|(17[0-9]))[0-9]{8}$`), //中国
	AreaCode_HK: regexp.MustCompile(`^(\+852)?(9|6|5)[0-9]{7}$`),                            //香港
	AreaCode_MO: regexp.MustCompile(`^(\+853)?(66|68)[0-9]{5}$`),                            //澳门
	AreaCode_TW: regexp.MustCompile(`^(\+886)?9[0-9]{8}$`),                                  //台湾
	AreaCode_US: regexp.MustCompile(`^(\+1)?[0-9]{10}$`),                                    //美国
}

/*
 *@note 精确验证手机号格式
 *@param phone 手机号码
 *@param areaCode 国家码
 *@return 错误码
 */
func (this *cRegexpPhoneCode) IsPhoneWithCode(phone string, areaCode string) int32 {

	for k, v := range this.regexpWithAreaCode {
		// 匹配国家编码
		if k != areaCode {
			continue
		}

		if v.MatchString(phone) {
			return errcode.No_Error
		}
	}

	return errcode.ErrInvalidPhone
}

/*
 *@note 验证是否手机号 - 不含国家码
 *@return 错误码
 */
func (this *cRegexpPhoneCode) IsPhone(phone string) int32 {
	for _, v := range this.regexpWithoutAreaCode {
		if v.MatchString(phone) {
			return errcode.No_Error
		}
	}

	return errcode.ErrInvalidPhone
}

/*
 *@note 分割区号和手机号
 *@param phoneWithCode 带区号的手机号码
 *@return 分割区号,手机号
 */
func (this *cRegexpPhoneCode) PhoneSplitCode(phoneWithCode string) (string, string) {
	for k, v := range this.regexpWithAreaCode {
		if v.MatchString(phoneWithCode) {
			return k, phoneWithCode[len(k):]
		}
	}

	return "", ""
}

/*
 *@note 通过手机号猜出区号
 *@param phone 不带区号的手机号码
 *@return 分割区号,手机号
 */
func (this *cRegexpPhoneCode) PhoneGuessCode(phone string) []string {
	reply := []string{}

	for k, v := range this.regexpWithoutAreaCode {
		if v.MatchString(phone) {
			// 有区号的手机号码必然匹配不了其他区域手机号码
			if strings.HasPrefix(phone, k) {
				return []string{phone}
			}

			reply = append(reply, k+phone)
		}
	}

	return reply
}

/*
 *@note 填充手机号码的区号
 *@param phones 传递进来的手机号码
 *@return 添加区号后的手机号码
 */
func (this *cRegexpPhoneCode) PhoneFillAreaCode(phones []string) []string {
	cLen := len(phones)

	for i := 0; i < cLen; i++ {
		for k, v := range this.regexpWithoutAreaCode {
			c := phones[i]
			if v.MatchString(c) {
				if !strings.HasPrefix(c, k) {
					phones[i] = k + c
				}

				break
			}
		}
	}

	return phones
}
