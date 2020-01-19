package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is a example."
	//按空格作为分隔符
	strings.Fields(str)
	//按自定义字符
	var slic []string= strings.split(str," ")
	//拼接slice
	strings.Join(slic, "aaa")//aaa为拼接字符
}
