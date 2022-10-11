// Package hsystem 系统操作
package hsystem

import (
	"os"
	"runtime"
	"syscall"
)

var globalPrintFile *os.File

// PrintOutAndErrToFile 将标准输出和错误输出写入到文件
func PrintOutAndErrToFile(fileName string) error {
	// 因很少运行在 Windows 上, 暂时不对 Windows 做支持
	if runtime.GOOS == "windows" {
		panic("Not currently supported windows")
	}
	var err error
	// 这里文件句柄不使用局部变量, 防止gc 时被回收
	globalPrintFile, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	if err = syscall.Dup2(int(globalPrintFile.Fd()), int(os.Stderr.Fd())); err != nil {
		return err
	}
	if err = syscall.Dup2(int(globalPrintFile.Fd()), int(os.Stdout.Fd())); err != nil {
		return err
	}
	return nil
}
