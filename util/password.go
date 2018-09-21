/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : password.go
 Time    : 2018/9/21 16:41
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"github.com/yanue/go-esport-common"
	"golang.org/x/crypto/bcrypt"
)

type passHelper struct {
}

var Password *passHelper = new(passHelper)

const passSalt = "!@:\"#$%^&*<>?{}$^&*I@!" // 密码加盐

/*
*@note 根据明文生成加密密码
*@param password 明文密码
@return 加密密码
 */
func (this *passHelper) Generate(password string) (string, error) {
	// 虽然每次生成都不同，只需保存一份即可
	hash, err := bcrypt.GenerateFromPassword([]byte(password+passSalt), bcrypt.DefaultCost)
	if err != nil {
		common.Logs.Info("pass verify fail: ", err)
		return "", err
	}

	return string(hash), nil
}

/*
*@note 密码验证
*@param password 明文密码
@param hash 通过PasswordGenerate生成的,保存在数据库的密码(虽然每次生成都不同，只需保存一份即可)
@return bool
 */
func (this *passHelper) Verify(password string, hash string) bool {
	// 正确密码验证
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+passSalt))
	if err != nil {
		//common.Logs.Info("pass verify fail: ", err)
		return false
	}

	return true
}
