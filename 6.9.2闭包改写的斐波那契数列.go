package main

import "fmt"

func main() {
	//递归调用方式
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	//匿名函数方式
	fibo := fibonacciFunc()
	for i := 0; i < 10; i++ {
		result = fibo()
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacciFunc() func() int {
	back1, back2 := 1, 1//后1和后2
	return func() int {
		temp := back1//记录后1的值
		back1, back2 = back2, back1 + back2//将原后2的值写入后1，同时保存新后2的值
		return temp
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n -1) + fibonacci(n-2)
	}
	return
}
