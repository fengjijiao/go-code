package main

import "time"

//出于性能考虑的建议：
//实践经验表明，如果你使用并行运算获得高于串行运算的效率：在协程内部已经完成的大部分工作，其开销比创建协程和协程间通信还高。

//1.出于性能考虑建议使用带缓存的通道
//使用带缓存的通道可以很轻易的成倍提高它的吞吐量，某些场景其性能可以提高至10倍甚至更多。通过调整通道的容量，甚至可以尝试着更进一步的优化其性能。

//2.限制一个通道的数量数量并将它们封装成一个数组：
//如果使用通道传递大量单独的数据，那么通道将变成性能瓶颈。然而，将数据块打包封装成数组，在接收端解压数据时，性能可能提高至10倍。
//创建：ch := make(chan type, buf)

//3.如何使用for或者for-range遍历一个通道
for v:= range ch {
	//do something with v
}
//4.如何检测通道ch是否关闭
//read channel until it closes or error-condition
for {
	if input, open := <-ch; !open{
		break
	}
	fmt.Printf("%s", input)
}
//或者使用3自动检测
//5.如何通过一个通道让主程序等待直到协程完成（信号量模式）
ch := make(chan int)
go func() {
	ch <- 1
}()
doSomethingElseForWhile()
<-ch
//6.通道的工厂模板：以下函数是一个通道工厂，启动一个匿名函数作为协程以生成通道
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0;;i++ {
			ch <-i
		}
	}()
	return ch
}
//7.通道迭代器模板
//8.如何限制并发处理请求的数量（14.11）
//9.如何在多核CPU上实现并行计算（14.13）
//10.如何终止一个协程：runtime.Goexit()
//11.简单的超时模板
timeout := make(chan bool, 1)
go func() {
	time.Sleep(1e9)
	timeout <- true
}()
select {
	case <-ch:
		//a read from ch has ocurred
	case <-timeout:
		//the read from ch has timed out
}
//12.如何使用输入通道和输出通道替代锁：
func Worker(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <-t
	}
}
//13.如何在同步调用运行时间过长时将之丢弃(14.5第二个变体)
//14.如何在通道中使用计时器和定时器（14.5）
//15.典型的服务器后端模型(14.4)
