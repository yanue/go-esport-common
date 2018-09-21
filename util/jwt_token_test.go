/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : jwt_test.go
 Time    : 2018/9/20 18:30
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"fmt"
	pb "github.com/yanue/go-esport-common/proto"
	"testing"
)

func init() {
	JwtToken = &jwtToken{
		PJwtToken: &pb.PJwtToken{},
	}
}

func TestJwtToken_GenerateToken(t *testing.T) {
	// 生成
	a, e := JwtToken.Generate(10243560, pb.Os_WEB, pb.ELoginType_ACCOUNT)
	fmt.Println("a,e", a, e)
	// 验证
	c, e := JwtToken.Verify(a)
	fmt.Println("c,e", c, e)
}
