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
*@note 是否为密码格式
*@param pass 密码
*@remark 错误码
 */
func (this *validator) IsPassword(pass string) int32 {
	l := len(pass)

	if l < LenMinPassword || l > LenMaxPassword {
		return errcode.ErrInvalidPassword
	}

	return errcode.No_Error
}

/*
*@note 是否为Email
*@param email Email地址
*@remark 错误码
 */
func (this *validator) IsEmail(email string) int32 {
	if email == "" {
		return errcode.ErrInvalidEmail
	}

	if !regexpEmail.MatchString(email) {
		return errcode.ErrInvalidEmail
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
		return errcode.ErrInvalidUrl
	}

	// 通过原生url parse验证
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return errcode.ErrInvalidUrl
	}

	return errcode.No_Error
}
