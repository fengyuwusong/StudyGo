package main

import (
	"os"
	"io"
)

// defer介绍 : 用于关闭资源等...
//多个资源可用匿名函数 defer func(){...}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close() //会在方法关闭前执行
	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
