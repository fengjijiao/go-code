package main

import (
	"errors"
	"fmt"
)

//Go有一个预先定义的error接口类型
//type error interface {
//	Error() string
//}
//错误值用来表示异常状态；我们可以在5.看到标准用法
//处理文件操作的例子可以在12章找到
//15章将有网络操作的例子。
//error包中有一个errorString结构体实现了error接口。当程序处于错误状态时可以用os.Exit(1)来中止。

//13.1.1定义错误
//任何时候当你需要一个新的错误类型，都可以用errors包的errors.New函数接收合适的错误信息来创建，像下面这样：
err := errors.New("math - square root of negative number")

var errNotFound error = errors.New("not found error")

func main() {
	fmt.Printf("error: %v\n", errNotFound)
}

//可以把它用于计算平方根函数的参数测试：
//func Sqrt(f float64) (float64, error) {
//	if f < 0 {
//		return 0, errors.New("math - square root of negative number")
//	}
//	//implementation of Sqrt
//}
//可以像下面这样调用
//if f, err := Sqrt(-1); err != nil {
//	fmt.Printf("Error: %s\n", err)
//}
//由于fmt.Printf会自动调用String()方法，所以错误信息"Error: math - square root of negative number"会打印出来。通常错误信息都有像“Error”这样的前缀，所以你的错误信息不要以大写字母开头。
//在大部分情况下自定义错误结构很有意义的，可以包含除了（低层级的）错误信息以外的其他有用信息，例如，正在进行的操作（打开文件等），全路径或名字。
type PathError struct {
	Op string//open\unlink etc
	Path string//associated file.
	Err error//returned by the system call
}

func (p *PathError) String() string {
	return p.Op + " " + p.Path + ": " + p.Err.Error()
}
//如果有不同错误条件可能会发生，那么对于实际的错误使用类型断言或类型判断（type-switch）是很有用的，并且可以根据错误场景做一些补救和恢复操作。
//err != nil
//if e, ok := err.(*os.PathError); ok {
//	//remedy situation
//}
//或
//switch err := err.(type) {
//	case ParseError:
//		PrintParseError(err)
//	case PathError:
//		PrintPathError(err)
//	...
//	default:
//		fmt.Printf("not a special error, just%s\n", err)
//}
//作为第二个例子考虑用json包的情况下。当json.Decode在解析JSON文档发生语法错误时，指定返回一个SyntaxError类型的错误。
//type SyntaxError struct {
//	msg string//description of error
//	//error occurred after reading Offset bytes, from which line and column can be obtained
//	Offset int64
//}
//在调用代码中你可以像这样用类型断言测试错误是不是上面的类型
//if serr, ok := err.(*json.SyntaxError); ok {
//	line, col := findLine(f, serr.Offset)
//	return fmt.Errorf("%s:%d:%d: %v",f.Name(), line, col, err)
//}
//包也可以用额外的方法(methods)定义特定的错误，比如net.Error：
//package net
//type Error interface {
//	Timeout bool //Is the error a timeout?
//	Temporary() bool //Is the error temporary?
//}
//在15.1我们将看到如何使用它.
//正如你所看到的一样，所有的例子都遵循一种命名规范：
//错误类型以Error结尾，错误变量以err或Err开头。
//syscall是低阶外部包，用来提供系统基本调用的原始接口。它们返回整数的错误码；类型syscall.Errno实现了Error接口。
//大部分syscall函数都返回一个结果和可能的错误，比如：
//r.err := syscall.Open(name, mode, perm)
//if err != nil {
//	fmt.Printf(err.Error())
//}
//os包也提供了一套像os.EINAL这样的标准错误，它们基于syscall错误：
//var (
//	EPERM	Error = Errno(syscall.EPERM)
//...
//)


//13.1.2用fmt创建错误对象
//通常你想要返回包含错误参数的更有信息量的字符串，例如：可以用fmt.Errorf()来实现：它和fmt.Printf()完全一样，接收有一个或多个格式占位符的格式化字符串和相应数量的占位变量。和打印信息不同的是它用信息生成错误对象。
//比如前面平方根例子中使用：
//if f < 0 {
//	fmt.Errorf("math - square root of negative number %g", f)
//}
//第二个例子：从命令行读取输入时，如果加了help标志，我们可以用有用的信息产生一个错误：
//if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
//	err = fmt.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
//  return
//}