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

func TestJwtToken_GenerateToken(t *testing.T) {
	// 生成
	a, p, e := JwtToken.Generate(1024356, pb.ELoginType_ACCOUNT, &pb.PDevice{
		Imei: "19AAB430-9CB8-4325-ACC5-D7D386B68960",
		// 操作系统类型
		Os: pb.Os_IOS,
		// 操作系统版本
		OsVersion: "12.0.2",
		// 设备型号，如iPhone 6s
		Model: "iPhone 6s",
	})
	fmt.Println("a,e", a, p, e)
	// 验证
	c, p, e := JwtToken.Verify(a)
	fmt.Println("c,e", c, p, e)
	// CgVIUzI1NhIFcHJvdG8.EOTCPhjGxabdBSAC
	// CgVIUzI1NhIFcHJvdG8.EOTCPiAC.c6f1c6d547b05fdab6b63a050df2653be946d6f7f190ada757db3cbdcca51830 <nil>
	// CgVIUzI1NhIFcHJvdG8.EOTCPiAC.c6f1c6d547b05fdab6b63a050df2653be946d6f7f190ada757db3cbdcca51830 <nil>
	d, e := JwtToken.ParsePayload(p)
	fmt.Println("d,e", d, e)
}
