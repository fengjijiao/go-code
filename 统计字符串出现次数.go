package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is a example."
	fmt.Printf("%d\n",strings.Count(str,"a"))
}
