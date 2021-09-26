package main

import "fmt"

//生成器是指当被调用时返回一个序列中下一值的函数，例如:
//generateInteger() => 0
//generateInteger() => 1
//generateInteger() => 2
//...
//生成器每次返回的是序列中下一值而非整个序列；这种特性也被称之为惰性求值：只在你需要时进行求值，同时保留相关变量资源（内存和CPU）：这是一项在需要时对表达式进行求值的技术。例如：生成一个无限数量的偶数序列：要产生这样一个序列并且在一个一个的使用可能会很困难，而且内存会溢出！但是一个含有通道和go协程的函数能轻易实现这个需求。
var resume chan int
func interges() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count*2
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func main2() {
	resume = interges()
	fmt.Println(generateInteger())//0
	fmt.Println(generateInteger())//2
	fmt.Println(generateInteger())//4
	fmt.Println(generateInteger())//6
	fmt.Println(generateInteger())//8
	fmt.Println(generateInteger())//10
}

//有一个细微的区别是从通道读取的值可能会是稍早前产生的，并不是在程序被调用时生成的。如果确实需要这样的行为，就得实现一个请求响应机制。当生成器生成数据的过程是计算密集型且各个结果的顺序并不重要时，那么就可以将生成器放入go协程实现并行化。但是得小心，使用大量的Go协程的开销可能会超过带来的性能增益。
//这些原则可以概括为：通过巧妙地使用空接口、闭包和高阶函数，我们能实现一个通用的惰性生成器的工厂函数BuildLazyEvaluator（这个应该放在一个工具包中实现）。工厂函数需要一个函数和初始状态作为输入参数，返回一个无参、返回值是生成序列的函数。传入的函数需要计算出下一个返回值以及下一个状态参数。在工厂函数中，创建一个通道和无限循环的go协程。返回值被放到了该通道中，返回函数稍后被调用时从该通道中取得该返回值。每当取得一个值时，下一个值即被计算。在下面的例子中，定义了一个evenFunc函数，其是一个惰性生成函数：在main函数中，我们创建了前10个偶数，每个都是通过调用even()函数取得下一个值的。为此，我们需要在BuildLazyIntEvaluator函数中具体化我们的生成函数，然后我们能够基于此做定义。

type Any interface {}
type EvalFunc func(Any) (Any, Any)//生成新值

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)//return: 上一个值，新值, eg: actState = 0, retVal = 0, actState = 2 (even)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <- retValChan
	}
	go loopFunc()
	return retFunc//返回一个取值函数
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {//返回一个取值函数
		return ef().(int)//Any to int
	}
}

func main3() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)//上一个值
		ns := os + 2//新值
		return os, ns
	}
	even := BuildLazyIntEvaluator(evenFunc, 0)
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even %v\n", i ,even())
	}
	/**
	0th even 0
	1th even 2
	2th even 4
	3th even 6
	4th even 8
	5th even 10
	6th even 12
	7th even 14
	8th even 16
	9th even 18
	*/
}

//练习14.12
var resume2 chan uint64

func fibs() chan uint64 {
	yield := make(chan uint64)
	go func() {
		var nv, ov uint64= 1, 1
		for {
			yield <- ov
			ov, nv = nv, ov + nv
		}
	}()
	return yield
}

func generateFib() uint64 {
	return <-resume2
}

func main() {
	resume2 = fibs()
	fmt.Println(generateFib())
	fmt.Println(generateFib())
	fmt.Println(generateFib())
	fmt.Println(generateFib())
	fmt.Println(generateFib())
	fmt.Println(generateFib())
	/**
	1
	1
	2
	3
	5
	8
	*/
}