/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : util_test.go
 Time    : 2018/9/14 17:43
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import "testing"

func TestStrTo_Exist(t *testing.T) {
	var a StrTo = "12321"
	c, d := a.Int64()
	t.Logf("a:%v-%v", c, d)
}
