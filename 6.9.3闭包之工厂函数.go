package main

import (
	"fmt"
	"strings"
)

/**
可以返回其他函数的函数和接收其他函数作为参数的函数均被称之为高阶函数，是函数式语言的特点。闭包在GO语言中很常见，常用于Goroutine和管道操作。
 */
func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Printf("%s\n", addBmp("file"))
	fmt.Printf("%s\n", addJpeg("file"))
	/*
	file.bmp
	file.jpeg
	*/
}

/**
一个返回值为另一个函数的函数可以被称之为工厂函数，这在需要创建一系列相似的函数时非常有用；书写一个工厂函数而不是针对每种情况都书写一个函数。
 */
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
