package main

import (
	"fmt"
)

func main() {
	//1.无限循环
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}
	//2.无限循环0
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}
	//3.
	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
	//4.
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
