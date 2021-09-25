package main

import "fmt"

//通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。
//只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道，接收者永远不会需要。

//如何在发送完成时发送一个信号？如何检测到通道是否关闭或阻塞？
//1.可以通过函数close(ch)来完成：这个将通道标记为无法通过发送操作符<-接受更多的值；给已经关闭的通道发送或者再次关闭都会导致运行时的panic。在创建一个通道后使用defer语句是个不错的办法（类似这种情况）：
//ch := make(chan float64)
//defer close(ch)
//2.可以使用逗号，ok操作符：用来检测通道是否被关闭。
//如何来检测可以收到没有被阻塞（或通道没有被关闭）？
//v, ok := <-ch//ok is true if v received value
//通常和if语句一起使用
//if v,ok := <-ch;ok {
//	process(v)
//}
//或者在for循环中接收的时候，当关闭或者阻塞的时候使用break:
//v, ok := <-ch
//if !ok {
//	break
//}
//process(v)


//非阻塞通道的读取，使用了select。
func sendData(ch chan string) {
	ch <- "W"
	ch <- "A"
	ch <- "Q"
	ch <- "I"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
	/**
	W A Q I
	*/
}

//改变了以下代码：
//1.现在只有sendData()是协程,getData()和main()在同一个线程中。
//go sendData(ch)
//	getData(ch)
//2.在sendData()函数的最后，关闭了通道：
//close(ch)
//3.在for循环的getData()中，在每次接收通道的数据之前都使用if !open来检测
//使用for-range语句来读取通道是更好的办法，因为这会自动检测通道是否关闭：
//for input := range ch {
//	process(input)
//}
