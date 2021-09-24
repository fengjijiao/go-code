package main

import (
	"fmt"
	"os"
)

//当发生像数组下标越界或类型断言失败这样的运行错误时，Go运行时会触发运行时panic,伴随着程序的崩溃抛出一个runtime.Error接口类型的值。这个错误值有个RuntimeError()方法用于区别普通错误。
//panic可以之间从代码初始化：当错误条件（我们所测试的代码）很严苛且不可恢复，程序不能继续运行时，可以使用panic函数产生一个中止程序的运行时错误。panic接收一个任意类型的参数，通常是字符串，在程序死亡时被打印出来。Go运行时负责中止程序并给出调试信息。
func main() {
	fmt.Println("starting the program")
	panic("a severe error occurred: stopping the program!")
	fmt.Println("ending the program")
	/**
	starting the program
	panic: a severe error occurred: stopping the program!

	goroutine 1 [running]:
	main.main()
	        F:/GolandProjects/src/go-code/13.2运行时异常和panic.go:9 +0x65
	exit status 2
	*/
}
//一个检查程序是否被已知用户启动的具体例子
var user = os.Getenv("USER")

func check() {
	if user == "" {
		panic("unknown user: no value for $USER")
	}
}
//可以在导入包中的init函数中检查这些。
//当发生错误必须中止程序时，panic可以用于错误处理模式。
//if err != nil {
//	panic("Error occurred: " + err.Error())
//}
//Go panicking:
//在多层嵌套的函数中调用panic，可以马上中止当前函数的执行，所有的defer语句都会保证执行并把控制权交还给接收到panic的函数调用者。这样向上冒泡直到最顶层，并执行（每层的）defer，在栈顶处程序崩溃，并在命令行中传给panic的值报告错误情况：这个终止过程就叫panicking。
//标准库中有许多包含Must前缀的函数，像regexp.MustComplie和template.Must；当正则表达式或模板中转入的转换字符串导致错误时，这些函数会panic。
//不能随意地用panic中止程序，必须尽力补救错误让程序能继续执行。
