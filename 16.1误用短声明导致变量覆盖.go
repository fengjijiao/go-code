package main

import "fmt"

func main() {
	var remember bool = false
	if 1 == 1 {
		remember := true
	}
	println(remember)
	//在这段代码中remember永远不可能为true，由于使用了短声明，if语句内部的新变量remember将覆盖外面的remember变量并且该值变为了true，但在if语句外面，变量remember的值依然是false。

	//此错误也容易在for循环中出现，尤其当函数返回一个具名变量时难以察觉：
	//func shadow() (err error)
}

//func shadow() (err error) {
//	x, err := check1()//x是新创建变量，err是被赋值
//	if err != nil {
//		return//正确的返回了err
//	}
//	if y, err := check2(x); err != nil {
//		return//if语句中的err覆盖了外面的err，所以错误的返回了nil
//	}else {
//		fmt.Println(y)
//	}
//	return
//}
