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

func TestPasswordGenerate(t *testing.T) {
	for i := 0; i < 10; i++ {
		a, e := Password.Generate("111111")
		t.Logf("a:%v %v", a, e)
	}
}

func TestPasswordVerify(t *testing.T) {
	a := Password.Verify("123456", "$2a$10$S3P7gltm08db5JmXlD1UPu7WJ9jnQ4a1laefrut9Acgq8njAJPAFC")
	t.Logf("a:%v", a)
	a = Password.Verify("1234546", "$2a$10$S3P7gltm08db5JmXlD1UPu7WJ9jnQ4a1laefrut9Acgq8njAJPAFC")
	t.Logf("a:%v", a)
}
