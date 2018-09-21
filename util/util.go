/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : util.go
 Time    : 2018/9/14 17:32
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


/*
*@note 得到当前exe程序执行的不目录
 */
func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
