package main

import (
	"io"
	"os"
)

//如何拷贝一个文件到另一个文件？最简单的方式就是使用io包

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	//当打开目标文件时发生了错误，那么defer仍然能够确保src.Close()执行。如果不这么做，文件会一直保持打开状态并占用资源。
	return io.Copy(dst, src)
}

func main() {
	CopyFile("target.txt", "source.txt")
}
