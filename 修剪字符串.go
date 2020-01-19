package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "   This is a example string. "
	//去除两边空格
	fmt.Printf("%s\n",strings.TrimSpace(str))
	//去除两边指定字符
	fmt.Printf("%s\n",strings.Trim(str,"T"))
}
