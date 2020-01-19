package main

import (
	"fmt"
)

func main() {
	var a,b int = 2, 0
	switch a {
		case 0:
			b = 9
		case 1:
			b = 6
		case 2:
			b = 3
			fallthrough //当使用此关键词时将会继续执行后续分支代码
		case 3:
			b = 0
		default:
			b = -3
	}
	fmt.Printf("%d\n", b)
	//不提供任何被判断的值
	switch {
		case a>=0 && a<3:
			fmt.Println("a属于[0,3)")
		case a>=3 && a<6:
			fmt.Println("a属于[3,6)")
		case a>=6 && a<9:
			fmt.Println("a属于[6,9)")
		default:
			fmt.Println("a所有判断条件都不满足")
	}
	//包含一个初始化语句
	switch result :=calculate(); {
		case result < 0:
			fmt.Println("result < 0")
		case result > 0:
			fmt.Println("result > 0")
		default:
			fmt.Println("resutl = 0")
	}
}

func calculate() int {
return -3
}
