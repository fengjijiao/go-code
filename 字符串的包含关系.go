package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is a example string."
	fmt.Printf("%t\n",strings.Contains(str,"string"))
}
