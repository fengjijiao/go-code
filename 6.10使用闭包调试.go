package main

import (
	"fmt"
	"log"
	"runtime"
)

/**
在分析和调试复杂的程序时，无数个函数在不同的代码文件中相互调用，如果这时候能够准确地知道哪个文件中的具体哪个函数正在执行，对于调试是十分有帮助的。你可以使用runtime或log包中的特殊函数来实现这样的功能。包runtime中断函数Caller()提供了相应的信息。因此可以在需要的时候实现一个where闭包函数来打印函数执行的位置。
 */
func main() {
	//1.
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d\n", file, line)
	}
	for i := 0; i < 30; i++ {
		where()
		//some code
		fmt.Printf("%i\n", i)
		where()
		//some code
	}
	//2.也可以设置log包中的flag参数来实现
	log.SetFlags(log.Llongfile)
	log.Printf("")
	//3.或使用一个更加简易版的where函数
	var where2 = log.Print
	for i := 0; i < 30; i++ {
		where2()
		//some code
		fmt.Printf("%i\n", i)
		where2()
		//some code
	}
}
