package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS string = "Hi there!"
	var newS string
	newS = strings.Repeat(origS, 3)
	fmt.Printf("new Sting is :%s\n",newS)
}
