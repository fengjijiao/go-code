package main

//可以很轻松的实现一个带缓冲的通道（14.2.5）,它的容量是并发请求的最大数目。下面的示例没有做任何事情，它包含了下列技巧：不超过MAXREQS的请求将被处理并且是同时处理，因为当通道sem的缓冲区全被占用时，函数handle被阻塞，直到缓冲区的请求被执行完成并且从sem中删除之前，不能执行其他操作。sem就像一个semaphore(信号量)，表示一个在一定条件的程序中的一个标志变量的技术术语：由此得名。
const (
	AvailableMemory = 10<<20//10MB
	AverageMemoryPerRequest = 10<<10//10KB
	MAXREQS= AvailableMemory / AverageMemoryPerRequest
)
var sem = make(chan int, MAXREQS)
type Request struct {
	a, b int
	replyc chan int
}

func process(r *Request) {
	//do something
	//可能需要很长时间并使用大量内存或CPU
}

func handle(r *Request) {
	process(r)
	//信号完成：开始启用下一个请求
	//将sem的缓冲区释放一个位置
	<-sem
}

func Server(queue chan *Request){
	for {
		sem <- 1 //当通道已满时，将会在此等候空闲
		request := <- queue
		go handle(request)
	}
}

func main() {
	queue := make(chan *Request)
	go Server(queue)
}
//通过这种方式，程序中的协程通过使用缓冲通道（这个通道作为一个semaphore被使用）来调整资源的使用，实现了对内存等有限资源的优化。