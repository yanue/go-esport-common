/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : name.go
 Time    : 2018/9/14 17:44
 Author  : yanue
 
 - 用户名校验相关
 
------------------------------- go ---------------------------------*/

package validator

import (
	"github.com/yanue/go-esport-common/errcode"
	"regexp"
	"strings"
	"unicode/utf8"
)

// unicode编码大全参考地址：http://unicode-table.com/cn/#control-character
// 控制字符：\u000d\u000a(换行)
// 基本拉丁字母：\u0020-\u007f(\u0020:空格)(\u0041-\u005a:大写)(\u0061-\u007a:小写)
// 拉丁文补充1：\u00a1-\u00ff
// 拉丁文扩展A：\u0100-\u017f
// 拉丁文扩展B：\u0180-\u024f
// 国际音标扩展：\u0250-\u02af
// 占位修饰符号：\u02b0-\u02ff
// 结合附加符号：\u0300-\u036f
// 常用标点：\u2010, \u2012-\u2027, \u2030-\u205e
// 上标及下标 - 杂项符号和箭头：\u2070-\u2bff(表情符号)
// 中日韩符号和标点：\u3000-\u303f
// 中日韩统一表意文字：\u4e00-\u9fff(简体中文、繁体中文范围)
// 半角及全角形式：\uff00-\uffef
// 谚文音节：\uac00-\ud7af(韩文范围)
// 多文种补充平面：\u10000-\u1ffff(扩展的表情符号)
// 日文平假名/片假名: \u3040-\u30ff

type cRegexpName struct {
	regexpAccount  *regexp.Regexp
	regexpNickname *regexp.Regexp
	regexpRealName *regexp.Regexp
}

// 验证规则
var (
	singleWithSpace = regexp.MustCompile(` {2,}`)
	regexpAccount   = regexp.MustCompile("^[a-z]{1}([a-z0-9]|_)+$")
	regexpNickname  = regexp.MustCompile("[\u00A0\u1680\u180E\u2000-\u200B\u202F\u205F\u3000\uFEFF\n\b]")
	regexpRealName  = regexp.MustCompile("^[\u4e00-\u9fff\uac00-\ud7af\u0041-\u005a\u0061-\u007a]+$")
)

/*
*@note 是否为账号名格式
*@param account 账号名
*@return 错误码
*/
func (this *cRegexpName) IsAccount(account string) int32 {
	str := strings.ToLower(account)

	if !this.regexpAccount.MatchString(str) {
		return errcode.ErrAccountFmtAccount
	}

	l := len(str)
	if l < LenMinAccount || l > LenMaxAccount {
		return errcode.ErrAccountLenRealname
	}

	return errcode.No_Error
}

/*
*@note 是否为昵称格式
*@param nickname 昵称
*@return 错误码
*/
func (this *cRegexpName) IsNickname(nickname string) int32 {
	var i, n, c int

	// 分拆中文长度
	for i < len(nickname) {
		_, size := utf8.DecodeRuneInString(nickname[i:])
		c++
		i += size
		if size < 3 {
			n++
		} else {
			n += 2
		}
	}

	// 判断长度
	if c < LenMinNickname || n > LenMaxNickname {
		return errcode.ErrAccountLenNickname
	}

	// 去空验证
	if !this.regexpNickname.MatchString(nickname) && !singleWithSpace.MatchString(nickname) {
		return errcode.No_Error
	}

	return errcode.ErrAccountFmtNickname
}

/*
*@note 是否为真实姓名格式
*@param nickname 昵称
*@return 错误码
*/
func (this *cRegexpName) IsRealName(realName string) int32 {
	strlen := len(realName)

	if strlen < LenMinRealName || strlen > LenMaxRealName {
		return errcode.ErrAccountLenRealname
	}

	// 去空验证
	if !this.regexpRealName.MatchString(realName) && !singleWithSpace.MatchString(realName) {
		return errcode.No_Error
	}

	return errcode.ErrAccountFmtRealname
}
