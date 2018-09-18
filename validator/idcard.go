/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : id_card.go
 Time    : 2018/9/18 14:55
 Author  : yanue
 
 - 身份证验证相关
 
------------------------------- go ---------------------------------*/

package validator

import (
	"github.com/yanue/go-esport-common/errcode"
	"strconv"
	"strings"
	"time"
)

var (
	idCardArea = map[string]string{
		"11": "北京",
		"12": "天津",
		"13": "河北",
		"14": "山西",
		"15": "内蒙",
		"21": "辽宁",
		"22": "吉林",
		"23": "黑龙",
		"31": "上海",
		"32": "江苏",
		"33": "浙江",
		"34": "安徽",
		"35": "福建",
		"36": "江西",
		"37": "山东",
		"41": "河南",
		"42": "湖北",
		"43": "湖南",
		"44": "广东",
		"45": "广西",
		"46": "海南",
		"50": "重庆",
		"51": "四川",
		"52": "贵州",
		"53": "云南",
		"54": "西藏",
		"61": "陕西",
		"62": "甘肃",
		"63": "青海",
		"64": "宁夏",
		"65": "新疆",
		"71": "台湾",
		"81": "香港",
		"82": "澳门",
		"91": "国外",
	}
	idCardMinDate = time.Date(1890, 0, 0, 0, 0, 0, 0, time.Local)
	idCardMaxDate = time.Now()
	idCardWeight  = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}    //十七位数字本体码权重
	idCardCode    = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'} // 校验码
)

func (this *validator) IsIdCard(idCardNo string) int32 {
	// 二代身份证
	if len(idCardNo) != 18 {
		return errcode.Err_InValid_IdCardNo
	}

	//1. 区域校验
	if _, ok := idCardArea[idCardNo[0:2]]; !ok {
		return errcode.Err_InValid_IdCardNo
	}

	//2. 校验生日,包括格式和范围
	birth := idCardNo[6:14]
	if date, err := time.Parse("20060102", birth); err != nil {
		return errcode.Err_InValid_IdCardNo
	} else if date.After(idCardMaxDate) && date.Before(idCardMinDate) {
		return errcode.Err_InValid_IdCardNo
	}

	//3. 验证校验和
	sum := 0
	// 根据前17位计算校验码
	for i, char := range idCardNo[:len(idCardNo)-1] {
		charF, _ := strconv.ParseFloat(string(char), 64)
		sum += int(charF) * idCardWeight[i]
	}

	// 取模
	if idCardCode[sum%11] != idCardNo[len(idCardNo)-1] {
		return errcode.Err_InValid_IdCardNo
	}

	return errcode.No_Error
}

/**
 * 身份证号码与出生日期的验证
 * 前提是：身份证号和出生日期格式都已经过格式验证
 * 身份证号为15或18位，出生日期格式为XXXX-XX-XX
 */
func (this *validator) IsIdCardAndBirthday(idCardNo string, birthday string) int32 {
	if birthday == "" {
		return errcode.Err_InValid_Birthday
	}

	// 验证身份证
	if code := this.IsIdCard(idCardNo); code > 0 {
		return code
	}

	// 解析生日格式
	inputDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return errcode.Err_InValid_Birthday
	}

	nowDate := time.Now()

	//出生日期时间不能大于今天,请检查!
	if inputDate.Unix() > nowDate.Unix() {
		return errcode.Err_InValid_Birthday
	}

	var idBirthday string

	//15位身份证
	if len(idCardNo) == 15 {
		//从ID NO 中截取生日6位数字,前面加上19
		idBirthday = "19" + idCardNo[6:12]
	} else {
		//从ID NO 中截取生日8位数字
		idBirthday = idCardNo[6:14]
	}

	//日期字符串中的8位生日数字
	if idBirthday != strings.Replace(birthday, "-", "", -1) {
		return errcode.Err_InValid_Birthday
	}

	return errcode.No_Error
}

/*
*@note 是否为生日
*@param birthday 生日日期 (YYYY-mm-dd格式)
*@remark 错误码
 */
func (this *validator) IsBirthday(birthday string) int32 {
	if birthday == "" {
		return errcode.Err_InValid_Birthday
	}

	// 解析生日格式YYYY-mm-dd格式
	_, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return errcode.Err_InValid_Birthday
	}

	return errcode.No_Error
}
