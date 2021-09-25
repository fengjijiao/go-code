package main

import (
	"fmt"
	"time"
)

//time包中有一些有趣的功能可以和通道组合使用。
//其中就包含了time.Ticker结构体，这个对象以指定的时间间隔重复的向通道c发送时间值：
//type Ticker struct {
//	C <- chan Time //the channel on which the ticks are delivered.
//	//contains filtered or unexported fields
//}
//时间间隔的单位是ns（纳秒，int64），在工厂函数time.NewTicker中以Duration类型的参数传入 func NewTicker(dur Duration) *Ticker。
//在协程周期性的执行一些事情（打印状态日志，输出，计算等待）的时候非常有用。
//调用Stop()使计时器停止，在defer语句中调用。这些都很好的适用select语句。
//updateInterval := 2*time.Second
//ticker := time.NewTicker(updateInterval)
//defer ticker.Stop()
////...
//select {
//case u:= <-ch1:
//	//...
//case v := <-ch2:
//	//...
//case <-ticker.C://每2s将会从该通道（ticker.C）发送时间戳
//	logState(status)
//default:
//	//...
//}

func main2() {
	d := time.NewTicker(2 * time.Second)
	//ch1 := make(chan bool)
	//go func() {
	//	time.Sleep(7 * time.Second)
	//	ch1 <- true
	//}()
	for {
		select {
		//case <- ch1:
		//	fmt.Println("completed!")
		//	return
		case tm := <-d.C://每2s从通道收到一个ticker数据（时间戳）
			fmt.Printf("the current time is: %v\n", tm)
		}
	}
	/**
	the current time is: 2021-09-25 19:04:31.1563411 +0800 CST m=+2.002415201
	the current time is: 2021-09-25 19:04:33.1583395 +0800 CST m=+4.004413601
	the current time is: 2021-09-25 19:04:35.1602311 +0800 CST m=+6.006305201
	completed!
	*/
}

//time.Tick()函数声明为Tick(d Duration) <- chan Time,当你想返回一个通道而不必关闭它的时候这个函数非常有用：它以d为周期给返回的通道发送时间，d是纳秒。如果需要像下边的代码一样，限制处理频率（函数 client.Call()是一个RPC调用,15.9）
//rate_per_sec := 10
//var dur Duration = 1e9 / rate_per_sec
//chRate := timeTick(dur)//a tick every 1/10th of a second
//for req := range requests {
//	<- chRate//rate limit Service.Method RPC calls
//	go client.Call("Service.Method", req, ...)
//}
//这样只会按照特定频率处理请求：chRate阻塞了更高的频率。每秒处理的频率可以根据机器负载（和/或）资源的情况而增加或减少。
//问题14.1扩展上面的代码，思考如何承载周期请求数的暴增（提示：使用带缓冲通道和计时器对象）。
//定时器Timer的结构体看上去和计时器Ticker结构体确实很像（构造为NewTimer(d Duration)），但是它只能发送一次时间，在Duration d之后。
//还有time.After(d)函数，声明如下：
//func After(d Duration) <-chan Time
//在Duration d之后，当前时间被发送到返回的通道；所以它和NewTimer(d).C是等价的；它类似Tick(),但是After()只发送一次时间。下边有个很具体的示例,很好的阐明了select中default的作用。

func main() {
	tick := time.Tick(1e8)//100ms
	boom := time.After(5e8)//500ms
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("   .")
			time.Sleep(5e7)
		}
	}
	/**
	tick.
	   .
	tick.
	   .
	   .
	tick.
	   .
	tick.
	BOOM!
	*/
}



//习惯用法：简单超时模式
//要从通道ch中接收数据，但是最多等待1s。先创建一个信号通道，然后启动一个lambda协程，协程在给通道发送数据之前是休眠的：
//timeout := make(chan bool, 1)
//go func() {
//	time.Sleep(1e9)//1s
//	timeout <- true
//}()
//然后使用select语句接收ch或者timeout的数据：如果ch在1s内没有收到数据，就选择到了time分支并放弃了ch的读取。
//select {
//case <- ch:
//	//a read from ch has occured
//case <- timeout:
//	//the read from ch has timed out
//	break
//}



//第二种形式：取消耗时很长的同步调用
//也可以使用time.After()函数替换timeout-channel。可以在select种通过time.After()发送的超时信号来停止协程的执行。以下代码，在timeoutNs纳秒后执行select的timeout分支后，执行client.Call的协程也随之结束，不会给通道ch返回值：
//ch := make(chan error, 1)
//go func() {
//	ch <- client.Call("Service.Method", args, &reply)
//}()
//select {
//case resp := <-ch
//    //use resp and reply
//case <-time.After(timeoutNs):
//	//call time out
//	break
//}
//注意缓冲大小设置为1是必要的，可以避免协程死锁以及确保超时的通道可以被垃圾回收。此外，需要注意在有多个case符合条件时，select对case的选择是伪随机的，如果上面的代码稍作修改如下，则select语句可能不会在定时器超时信号到来时立即选中time.After(timeoutNs)对应的case，因此协程可能不会严格的按照定时器设置的时间结束。
//ch := make(chan int, 1)
//go func(){
//	for{
//		ch <- 1
//	}
//}()
//L:
//	for {
//		select {
//		case <-ch:
//			//do something
//		case <-time.After(timeoutNs):
//			//call time out
//		break L
//		}
//	}


//第三种形式：假设程序从多个复制的数据库同时读取。只需要一个答案，需要接收首先到达的答案，Query函数获取数据库的连接切片并请求。并行请求每一个数据库并返回收到的第一个响应：
//func Query(conns []conn, query string) Result {
//	ch := make(chan Request, 1)
//	for _, conn := range conns{
//		go func(c Conn) {
//			select {
//			case ch <- c.DoQuery(query):
//			default:
//			}
//		}(conn)
//	}
//	return <- ch
//}
//再次声明，结果通道ch必须是带缓冲的：以保证第一个发送进来的数据有地方可存，确保放入的首个数据总会成功，所以第一个到达的值会被获取而与执行的顺序无关。正在执行的协程可以总是可以使用runtime.Goexit()来使用。

//在应用中缓冲数据：
//应用程序中用到了来自数据库（或者常见的数据存储）的数据时，经常会把数据缓存到内存中，因为从数据库中获取数据的操作代价很高；如果数据库中的值不发生变化就没有问题。但是如果值有变化，我们需要一个机制来周期性的从数据库重新读取这些值：缓存的值就不可用（过期）了，而且我们也不希望用户看到陈旧的数据。