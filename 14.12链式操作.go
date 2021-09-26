package main

import "flag"

//下面的示例演示了启动大量的协程是多么的容易。它发送在main函数的for循环中。在循环之后，向rightmode通道插入0,在不到1.5s的时间执行了100000个协程，并将结果100000打印。
//这个程序还演示了如何通过命令行参数定义一个协程的数量，并通过给flag.Int解析，例如chaining -n=7000，可以生成7000个协程.

var ngoroutine = flag.Int("n", 10000, "how mang goroutines")

func f(left, right chan int) {//将right通道收到的值+1后发送到left通道
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost//nil, nil(inited)
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)//leftmost: nil(inited), nil(inited)
		go f(left, right)//right = nil, ngoroutine个协程全阻塞
	}
	right <- 0//right = 0，其中一个协程进行处理
	//start the chaining
	x := <-leftmost
	println(x)
}
/**
left right
nil, leftmost(nil)
leftmost(nil), nil(1)
nil(1), nil(2)
nil(2), nil(3)
nil(3), nil(4)
nil(4), nil(5)
nil(5), nil(6)
nil(6), nil(7)
nil(7), nil(8)
nil(8), nil(9)
...
nil(n-1),nil(n)
//right<-0:实际上表示将0值送入nil(n)这个通道，送入之后nil(n-1)的值也将被求出(0+1),即可得到之后上一步的nil(n-1)的值，而后从f函数可以得知nil(n-2)的值...最终将会反推到leftmost的值（10000）
 */