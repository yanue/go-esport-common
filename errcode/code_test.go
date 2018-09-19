/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : code_test.go
 Time    : 2018/9/18 17:43
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package errcode

import "testing"

func TestGetError(t *testing.T) {

}

func TestConvertJsonFile(t *testing.T) {
	ConvertJsonFile("errcode.json")
}
