package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 25; i++ {
		fmt.Println(strings.Repeat("G",i+1))
	}
}
