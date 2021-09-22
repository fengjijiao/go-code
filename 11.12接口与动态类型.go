package main

import "fmt"

//在经典的面向对象语言（像C++，Java和C#）中数据和方法被封装为类的概念：类包含它们两者，并且不能剥离。
//Go没有类：数据（结构体或更为一般的类型）和方法是一种松耦合的正交关系。
//Go中的接口跟Java/C#类似：都是必须提供一个指定方法集的实现。但是更加灵活通用：任何提供了此接口方法实现代码的类型都隐式地实现了该接口，而不用显式的声明。
//和其他语言相比，Go是唯一结合了接口值，静态类型检查（是否该类型实现了某接口），运行时动态转换的语言，并且不需要显式地声明类型是否满足某接口。该特性允许我们在不改变已有代码的情况下定义和使用新接口。
//接收一个（或多个）接口类型作为参数的函数，其实参可以是任何实现了该接口的类型。实现了某个接口的类型可以被传给任何以此接口为参数的函数。
//类似与Python和Ruby这类动态语言中的动态类型(duck typing)；这意味着对象可以根据提供的方法被处理（例如，作为参数传递给函数），而忽略它们的实际类型：它们能做什么比它们是什么更重要。

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 0; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	//
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}

func main() {
	b := new(Bird)
	DuckDance(b)
	/**
	I am quacking!
	I am walking!
	I am quacking!
	I am walking!
	I am quacking!
	I am walking!
	I am quacking!
	I am walking!
	*/
	//如果对cat调用函数DuckDance()，Go会提示编译错误，Python和Ruby会以运行时错误结束。


	//11.12.2 动态方法调用
}