package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

//从不同的并发执行的协程中获取值可以通过关键字select来完成，它和switch控制语句非常相似也被称作通信开关；它的行为像是“你准备好了吗”的轮询机制；select监听进入通道的数据，也可以是用通道发送值的时候。
//select {
//case u := <- ch1:
//	//...
//case v := <- ch2:
//	//...
//default://no value ready to be received
//	//...
//}
//default语句是可选的；fallthrough行为，和普通的switch相似，是不允许的。在任何一个case中执行break或者return，select就结束了。
//select做的就是：选择处理列出的多个通信情况中的一个。
//1.如果都阻塞了，会等待直到其中一个可以处理
//2.如果多个可以处理，随机选择一个
//3.如果没有通道操作可以处理并且写了default语句，他就会执行default。

//在select中使用发送操作并且有default可以确保发送不被阻塞！如果没有case，select就会一直阻塞。

//select语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过break语句使循环退出。
func pump1(ch chan int) {
	for i := 0;; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0;; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <- ch1:
			fmt.Printf("received on channel1: %d\n", v)
		case v := <- ch2:
			fmt.Printf("received on channel2: %d\n", v)
		}
	}
}

func main2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e9)
}

//练习14.7
//a)
func tel(ch chan int) {
	for i :=0 ;; i++ {
		ch <- i
	}
}

func main3() {
	ch := make(chan int)
	go tel(ch)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		}
	}
}

//b)
//协程初始化(发送)在接收之后
func tel2(ch chan int, quit chan bool) {
	defer func() {
		quit <- true
	}()
	for i :=0 ;; i++ {
		ch <- i
		if i > 1000 {
			break
		}
	}
}

func main4() {
	ch := make(chan int)
	quit := make(chan bool)
	go tel2(ch, quit)
	for {
		select {
		case <-quit:
			close(ch)
			close(quit)
			fmt.Println("closed: ")
			os.Exit(1)
		case v := <-ch:
			fmt.Println(v)
		default:
			fmt.Println("default")
		}
	}
}

//练习14.8
func fib(ch chan int) {
	var n1, n2 int = 1, 1
	var n3 int
	defer close(ch)//使用select后必须关闭否则会导致死锁
	ch <- 1
	ch <- 1
	for i := 0; i < 25; i++ {
		n3 = n1 + n2
		ch <- n3
		n1, n2 = n2, n3
	}
}

func main5() {
	ch := make(chan int)
	go fib(ch)
	for {
		select {
		case v := <-ch:
			if v == 0 {
				os.Exit(1)
			}
			fmt.Printf("%d \n", v)
		}
	}
	/**
	1
	1
	2
	3
	5
	8
	13
	21
	34
	55
	89
	144
	233
	377
	610
	987
	1597
	2584
	4181
	6765
	10946
	17711
	28657
	46368
	75025
	121393
	196418
	exit status 1
	*/
}


//练习14.9,做一个随机位生成器，程序可以提供无限的随机0或1的序列

func main6() {
	ch := make(chan int)
	go func() {
		for {
			fmt.Print(<-ch, " ")
		}
	}()
	for {
		select {//select在case条件都满足时，随机执行一个
		case ch <- 0:
		case ch <- 1:
		}
	}
}

//练习14.10
//写一个可交互的控制台程序，要求用户输入二维平面极坐标上的点（半径和角度）。计算对应的笛卡尔坐标系上的点(x,y)。
//实际上做这种不提倡使用协程和通道，但是如果是运算量很大很耗时，这种方案设计就非常合适。
type Cartesian struct {
	x, y float64
}

type Polar struct {
	radius, angle float64
}

func conv(ch1 chan Polar, ch2 chan Cartesian) {
	var res Cartesian
	defer close(ch1)
	defer close(ch2)
	data := <-ch1
	res.x, res.y = data.radius * math.Cos(data.angle/180*math.Pi), data.radius * math.Sin(data.angle/180*math.Pi/2)
	ch2 <- res
}

func main7() {
	var data Polar
	var res Cartesian
	ch1 := make(chan Polar)
	ch2 := make(chan Cartesian)
	_, err := fmt.Scanf("%f %f", &data.radius, &data.angle)
	if err != nil {
		panic(err)
	}
	fmt.Printf("your input: (r%.2f a%.2f)\n", data.radius, data.angle)
	go conv(ch1, ch2)
	ch1 <- data
	res = <- ch2
	fmt.Printf("%.2f %.2f\n", res.x, res.y)
	/**
	12 60.6
	your input: (r12.00 a60.60)
	5.89 6.05
	*/
}


//练习14.11
//1-1/3+1/5-1/7+1/9...=pi/4
func calc(ch1 chan float64, n int) {
	defer close(ch1)
	var res float64
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			res = float64(1)/float64(1+2*i)
		}else {
			res = - float64(1)/float64(1+2*i)
		}
		ch1 <- res
	}
}

func main8() {
	var res float64
	ch1 := make(chan float64)
	go calc(ch1, 500000)

	for {
		v, ok := <- ch1
		if !ok {
			break
		}
		res += v
	}
	fmt.Printf("pi approximately equal %f\n", res*4)
	/**
	pi approximately equal 3.141591
	*/
}



//习惯用法：后台服务模式
//服务通常是用后台协程中的无限循环实现的，在循环中使用select获取并处理通道中的数据：
//backend goroutine
//func backend() {
//	for {
//		select {
//		case cmd := <- ch1:
//		case cmd := <- ch2:
//		case cmd := <- chStop:
//			//stop
//		}
//	}
//}
//在程序的其他地方给通道ch1,ch2发送数据，比如通道stop用来清除结束服务程序。
//另一种方式（但是不太灵活）就是（客户端）在chRequest上提交请求，后台协程循环这个通道，使用switch根据请求的行为来分别处理：
//func backend() {
//	for req := range chRequest {
//		switch req.Subject() {
//		case A1:
//		case A2:
//		default:
//		}
//	}
//}