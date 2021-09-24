package main

import (
	"fmt"
	"log"
)

//正如名字一样，这个（recover）内建函数被用于从panic或错误场景中恢复；让程序可以从panicking重新获得控制权，停止终止过程，进而恢复正常执行。
//recover只能用在defer修饰的函数中使用：用于取得panic调用中传递过来的错误值，如果是正常执行，调用recover会返回nil,且没有其他效果。
//总结：panic会导致栈被展开直到defer修饰的recover()被调用或者程序中止。
//下面的例子中的protect函数调用函数参数g来保护调用者防止从g中抛出的运行时panic，并展示panic中的信息：
func protect(g func()) {
	defer func() {
		log.Println("done")//println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g()
}
//这与Java和.NET这样的语言中的catch块类似
//log包实现了简单的日志功能：默认的log对象向标准错误输出中写入并打印每条日志信息的日期和时间。除了Println和Printf函数，其他的致命性函数都会都会在写完日志信息后调用os.Exit(1)，那些退出函数也是如此。
//而panic效果的函数会在写完日志信息后调用panic;可以在程序必须中止或发生了临界错误时使用它们，就像当web服务器不能启动时那样。
//log包用那些方法（methods）定义了一个Logger接口类型。
//这是一个展示panic，defer和recover怎么结合使用的完整例子：
func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e:= recover(); e != nil {
			fmt.Printf("panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("calling test\r\n")
	test()
	fmt.Printf("test completed\r\n")
}
//defer-panic-recover在某种意义上也是一中像if, for这样的控制流机制。
//Go标准库中许多地方都用了这个机制，例如：json包中的解码和regexp包中的Compile函数。Go库的原则是即使在包的内部使用了panic，在它的对外接口（API）中也必须用recover处理成返回显示的错误。
