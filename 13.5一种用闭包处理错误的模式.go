package main

import "fmt"

//每当函数返回时，我们应该检查是否有错误发生：但是这会导致重复乏味的代码。
//结合defer/panic/recover机制和闭包可以得到一个我们马上就要讨论的更加优雅的模式。不过这个模式只有当所有的函数都是同一种签名时可用，这样就有相当大的限制。一个很好的使用它的例子是web应用，所有的处理函数都是下面这样。
//func handler1(w http.ResponseWriter, r *http.Request) { ... }
//假设所有函数都有这样的签名
//func f(a type1, b type2)
//参数的类型和数量是不相关的
//我们给这个类型一个名称
//fType1 := func f(a type1, b type2)
//在我们的模式中使用了两个帮助函数
//1)check: 这是用来检查是否有错误和panic发生的函数
//func check(err error) { if err != nil { panic(err) } }
//2)errorhandler: 这是一个包装函数。接收一个fType1类型的函数fn并返回一个调用fn的函数。里面就包含有defer/recover机制。
//func errorHandler(fn fType1) fType1 {
//	return func(a type1, b type2) {
//		defer func() {
//			if err, ok := recover().(error); ok {
//				log.Printf("run time panic: %v", err)
//			}
//		}()
//		fn(a, b)
//	}
//}
//当错误发生时会recover并打印在日志中；除了简单的打印，应用也可以用template包为用户生成自定义输出。
//check()函数会在所有的被调用函数中调用，像这样：
//func f1(a type1, b type2) {
//	...
//	f, _, err := //call function/method
//	check(err)
//	t, err := //call function/method
//	check(err)
//	_, err2 := //call function/method
//	check(err)
//	...
//}
//通过这种机制，所有的错误都会被recover，并且调用函数后的错误检查代码也被简化为调用check(err)即可。在这种模式下，不同的错误处理必须对应不同的函数类型；它们（错误处理）可能被隐藏在错误处理包内部。可选的更加通用的方式是用一个空接口类型的切片作为参数和返回值。
//将在15.5中的web应用中使用这种模式。
//练习13.1
func badCall(a int, b int) (err error) {
	defer func() {
		if err = recover().(error); err != nil {
			fmt.Println("error from srcCall")
		}
	}()
	srcCall(a, b)
	err = nil
	return
}

func srcCall(a int, b int) int {
	c := a / b
	fmt.Printf("res: %d\n", c)
	return c
}
func main() {
	err := badCall(9, 0)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}