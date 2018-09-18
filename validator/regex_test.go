/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : init.go
 Time    : 2018/9/18 11:48
 Author  : yanue

 - 验证测试

------------------------------- go ---------------------------------*/

package validator

import (
	"fmt"
	"github.com/yanue/go-esport-common/errcode"
	"testing"
)

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

	fmt.Println("init.", Verify)
}

func Test_GetPhoneCodeInfo(t *testing.T) {
	code1 := "13512345678"
	code2 := "+8612512345678"
	code3 := "+8613512345678"
	code4 := "+8613813813813"
	code5 := "+886900001287"
	code6 := "900001287"

	area1, val1 := Verify.PhoneSplitCode(code1)
	if "" != area1 && "" != val1 {
		t.Logf("err val 1")
	}

	area2, val2 := Verify.PhoneSplitCode(code2)
	if "" != area2 && "" != val2 {
		t.Logf("err val 2")
	}

	area3, val3 := Verify.PhoneSplitCode(code3)
	if area3 != "+86" || val3 != "13512345678" {
		t.Logf("err val 3")
	}

	area4, val4 := Verify.PhoneSplitCode(code4)
	if area4 != "+86" || val4 != "13813813813" {
		t.Logf("err val 4")
	}

	area5, val5 := Verify.PhoneSplitCode(code5)
	if area5 != "+886" || val5 != "900001287" {
		t.Logf("err val 5")
	}

	area6 := Verify.PhoneGuessCode(code6)
	if len(area6) < 1 || area6[0] != "+886900001287" {
		t.Logf("err val 6")
	}
	t.Logf("done")
}

func Test_GuessFullPhoneCode(t *testing.T) {
	code1 := "13512345678"
	code2 := "+8612512345678"
	code3 := "+8613512345678"
	code4 := "+8613813813813"

	val1 := Verify.PhoneGuessCode(code1)
	if nil == val1 || val1[0] != "+8613512345678" {
		t.Logf("err val 1")
	}

	val2 := Verify.PhoneGuessCode(code2)
	if nil != val2 {
		t.Logf("err val 2")
	}

	val3 := Verify.PhoneGuessCode(code3)
	if nil == val3 || val3[0] != "+8613512345678" {
		t.Logf("err val 3")
	}

	val4 := Verify.PhoneGuessCode(code4)
	if nil == val4 || val4[0] != "+8613813813813" {
		t.Logf("err val 4")
	}
	t.Logf("done")
}

func Test_FillAreaCode(t *testing.T) {
	type xyz struct {
		In  []string
		Out []string
	}

	testData := &xyz{
		[]string{"13510989512", "+8613512345678", "91234567", "+97691234567"},
		[]string{"+8613510989512", "+8613512345678", "+97691234567", "+97691234567"}}

	reply := Verify.PhoneFillAreaCode(testData.In)

	for i := 0; i < len(testData.Out); i++ {
		if testData.Out[i] != reply[i] {
			t.Logf("in %d output(%s) != expect(%s)", i, reply[i], testData.Out[i])
		}
	}
	t.Logf("done")
}

func TestCRegexpName_IsAccount(t *testing.T) {
	xxoo := []string{
		"a12345",
		"A12345",
		"abcdefghi",
		"ABCDEFGHI",
		"1a345opq",
		" a345opq",
		"a 345opq",
		"a345opq ",
		"中文",
		"テスト",
		"회원님을",
		"aaa01234,sdfsdf",
		"a1234",
		"a123456789012",
	}

	for _, v := range xxoo {
		res := Verify.IsAccount(v)
		fmt.Println("isAccount : ", v, res)
	}
	t.Logf("done")
}

func TestCRegexpName_IsNickname(t *testing.T) {
	xxoo := []string{
		"\t\r\n\b",
		"a",
		"中文一二三四五六七八九十一二三四五六七八",
		"asdfadsfasdfasdfq2rsdfasdfdfed中ta",
		"a-a",
		"a－a",
		"a 345opq",
		"a345opq ",
		"a  345opq",
		" a345opq",
		"一二三四五六七",
		"中文三",
		"繁體義薄雲天",
		"aa123456789_",
		"frank",
		"”’‘’，。",
		".、？/",
		"+=== 啊 {}[]\\|:;",
		"#$%^&*()_",
		"`~!@",
		"テスト",
		"회원님을",
		"Leo Xu (徐川洋）",
		"sdf,sdfdf",
		"宇智波の赵四",
		"01234567890123456789一二",
	}

	for _, v := range xxoo {
		code := Verify.IsNickname(v)
		t.Logf("isAccount :%v,%v,%v", code, v, errcode.GetErrStr(code))
	}
	t.Logf("done")
}

func TestValidator_IsIdCard(t *testing.T) {
	xxoo := []string{
		"370882197811107818",
		"44152119870528881x",
		"130825199105138665",
		"44152119870528881234234",
		"absdfasdf234324",
		"210224199007216341",
	}

	for _, v := range xxoo {
		code := Verify.IsIdCard(v)
		t.Logf("IsIdCard :%v,%v,%v", code, v, errcode.GetErrStr(code))
	}

	t.Logf("done")
}

func TestValidator_IsBirthday(t *testing.T) {
	xxoo := []string{
		"2016-12-30",
		"20161230",
		"abcwerer3",
	}

	for _, v := range xxoo {
		code := Verify.IsBirthday(v)
		t.Logf("IsBirthday :%v,%v,%v", code, v, errcode.GetErrStr(code))
	}

	t.Logf("done")
}

func TestValidator_IsEmail(t *testing.T) {
	xxoo := []string{
		"yanue@outlook.com",
		"fbg@live.com",
		"abcwerer3",
		"sdfsdfsf",
		"sdfsdf@com",
		"1234234",
		"sdfsdf.com",
	}

	for _, v := range xxoo {
		code := Verify.IsEmail(v)
		t.Logf("IsEmail :%v,%v,%v", code, v, errcode.GetErrStr(code))
	}

	t.Logf("done")
}

func TestValidator_IsUrl(t *testing.T) {
	xxoo := []string{
		"http://www.baidu.com",
		"http://www.baidu.com/1234/abc",
		"http://www.baiducom",
		"http://www.baiducom/sdfdfdsfsdfdsf",
		"sdfsdf@com",
		"1234234",
		"sdfsdf.com",
	}

	for _, v := range xxoo {
		code := Verify.IsUrl(v)
		t.Logf("IsUrl :%v,%v,%v", code, v, errcode.GetErrStr(code))
	}

	t.Logf("done")
}

func TestValidator_IsIdCardAndBirthday(t *testing.T) {
	xxoo := map[string]string{
		"1987-05-28": "441521198705288816",
		"1967-04-01": "130503670401001",
		"1987-05-26": "441521198705288816",
		"19870528":   "441521198705288816",
	}

	for bir, id := range xxoo {
		code := Verify.IsIdCardAndBirthday(id, bir)
		t.Logf("IsUrl :%v,%v,%v,%v", code, id, bir, errcode.GetErrStr(code))
	}

	t.Logf("done")
}
