package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is a example of a string"
	fmt.Printf("T/F? Does the string \"%s\" has prefix %s?", str, "Th")
	//strings.HasPrefix判断某字符串是否已特定字符开头
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//同样的还有strings.HasSuffix用于判断某字符串是否以特定字符结尾
}
