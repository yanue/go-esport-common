/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : init.go
 Time    : 2018/9/18 11:48
 Author  : yanue

 - 其他验证

------------------------------- go ---------------------------------*/
package validator

import (
	"github.com/yanue/go-esport-common/errcode"
	"net/url"
	"regexp"
)

var (
	regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

/*
*@note 是否为Email
*@param email Email地址
*@remark 错误码
 */
func (this *validator) IsEmail(email string) int32 {
	if email == "" {
		return errcode.Err_Empty_Email
	}

	if !regexpEmail.MatchString(email) {
		return errcode.Err_InValid_Email
	}

	return errcode.No_Error
}

/*
*@note 是否为网址
*@param url url地址
*@remark 错误码
 */
func (this *validator) IsUrl(uri string) int32 {
	if uri == "" {
		return errcode.Err_InValid_Url
	}

	// 通过原生url parse验证
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return errcode.Err_InValid_Url
	}

	return errcode.No_Error
}
