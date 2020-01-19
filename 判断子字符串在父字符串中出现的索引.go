package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is a example string."
	//起始位置
	fmt.Printf("%t\n", strings.Index(str, "example"))
	//最后位置
	fmt.Printf("%t\n", strings.LastIndex(str, "example"))
	//非ascii编码
	fmt.Printf("%t\n", strings.IndexRuna(str, "中文"))
}
