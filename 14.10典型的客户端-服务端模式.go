package main

//Client-Server类的应用是协程和频道的大显身手的闪光点。
//客户端可以是任何一种运行在任何设备上的，且需要来自服务端信息的一种程序，所以它需要发送请求。服务端接收请求，做一些处理，然后给客户端发送响应信息。在通常情况下，就是多个客户端（很多请求）对一个（或几个）服务端。
//在Go中，服务端通常会在一个协程里操作一个对客户端的响应，所以协程和客户端请求是一一对应的。一种典型的做法就是客户端请求本身包含了一个频道，服务端可以用它来发送响应。
//例如，一个请求结构体类似如下形式，内嵌一个回复channel:
//type Request struct {
//	a, b int;
//	replayc chan int;
//}
//或者更通常如下
//type Reply struct {
//	//...
//}
//type Request struct {
//	arg1, arg2, arg3 some_type
//	replyc chan *Reply
//}

//服务端可以在一个goroutine里面为每个请求都分配一个run()函数，这个函数会把binOp类型的操作作用于整数，然后通过回复channel发送结果：
//type binOp func(a, b int) int
//func run(op binOp, req *Request) {
//	req.replyc <- op(req.a, req.b)
//}
//服务端通过死循环来从chan *Request接收请求，为了避免长时间运行而导致阻塞，可以为每个请求都开一个goroutine来处理：
//func server(op binOp, service chan *Request) {
//	for {
//		req := <-service;//requests arrive here
//		//为请求开一个goroutine
//		go run(op, req);
//		//不用等待op结束
//	}
//}

//使用startServer函数来启动服务的自有的协程
//func startServer(op binOp) chan *Request {
//	reqChan := make(chan *Request)
//	go server(op, reqChan)
//	return reqChan
//}

//startServer()将会在main()主线程里被调用。


//在下面的例子中，我们发送100个请求，并在所有请求发送完毕后，再逐个检查其返回的结果
//func main() {
//	adder := startServer(func(a,b int) int {return a+b})
//	const N =100
//	var reqs [N]Request
//	for i := 0; i < N; i++ {
//		req := &reqs[i]
//		req.a = i
//		req.b = i+N
//		req.replyc = make(chan int)
//		adder <- req
//	}
//
//	for i := N-1; i >- 0; i-- {
//		if <-reqs[i].replyc != N+2*i {
//			fmt.Println("fail at", i)
//		}else{
//			fmt.Println("request ", i, "is ok!")
//		}
//	}
//
//	fmt.Println("done")
//}
//执行100000个Goroutines程序，甚至可以看到它在几秒内完成。这说明了Goroutines是有多么的轻量；如果我们启动相同数量的实际线程，程序很快就会崩溃。



//14.10.2拆解：通过发信号通知关闭服务器

//在以前的版本中，服务器在主返回时并不会被干净的关闭；他被强行停止。为了改善这一点，我们可以像服务器提供第2个退出通道
//func startServer(op binOp) (service chan *Request, quit chan bool) {
//	service = make(chan *Request)
//	quit = make(chan bool)
//	go server(op, service, quit)
//	return service, quit
//}
//server函数使用select在服务通道和退出通道之间进行选择：
//func server(op binOp, service chan *Request, quit chan bool) {
//	for {
//		select {
//		case req := <- service:
//			go run(op, req)
//		case <-quit:
//			return
//		}
//	}
//}
//当true进入退出通道时，服务器返回并终止
//主要改变下面这一行
//addr, quit := startServer(func(a, b int) int {return a+ b})
//在结尾处，我们放置该行
//quit <- true