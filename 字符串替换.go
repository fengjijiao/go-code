package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is a example."
	fmt.Printf("%s\n",strings.Replace(str, "example","apple",-1))//-1为替换所有
}
