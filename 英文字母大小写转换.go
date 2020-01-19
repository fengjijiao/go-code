package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Hey, how are you george?"
	var lower,upper string
	fmt.Printf("original string is:%s\n",orig)
	lower = strings.ToLower(orig)
	fmt.Printf("lowercase string is:%s\n",lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("uppercase string is:%s\n",upper)
}
