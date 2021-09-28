package main

import "fmt"

//如何使用内建函数recover终止panic过程（13.3）
func protect(g func()) {
	defer func() {
		fmt.Println("done")
		if x:= recover(); x != nil {
			fmt.Printf("run time panic: %v\n", x)
		}
	}()
	fmt.Println("start")
	g()
}
