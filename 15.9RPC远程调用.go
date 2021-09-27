package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

//Go程序可以通过net/rpc包互相通信，所以这是另一个客户端-服务端模式的应用。它提供了通过网络连接进行函数调用的便捷方法。只有程序运行在不同的机器上它才有用。rpc包建立在gob（12.11）上，将其编码/解码，自动转换为可以通过网路调用的方法。
//服务器注册一个对象，通过对象的类型名称暴露这个服务：注册后就可以通过网络或者其他远程客户端的IO连接它的导出方法。
//这个包使用了http协议、tcp协议和用于数据传输的gob包。服务器可以注册多个不同类型的对象（服务），但是相同的类型注册多个对象的时候会出错。
//一个简单的示例：定义一个Args类型，并在它上面创建一个Multiply方法，最好封装在一个单独的包中；这个方法必须返回一个可能的错误：
//package rpc_objects

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.M * args.N
	return nil
}

//服务器创建一行用于计算的对象，并且将它通过rpc.Register(object)注册，调用HandleHTTP(),并在一个地址上使用net,Listen开始监听。你也可以通过名称注册对象，如：rpc.RegisterName("Calculator", calc)
//对每一个进入到listener的请求，都是由协程去启动一个http.Serve(listener, nil),为每一个传入的HTTP连接创建一个新的服务线程。我们必须保证在一个特定的时间内服务器是唤醒状态的，例如: time.Sleep(1000e9)//1000s

////rpc_server.go

func runServer(done chan bool) {
	calc := new(Args)
	err := rpc.Register(calc)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	go func() {
		err := http.Serve(listener, nil)
		if err != nil {
			panic(err)
		}
	}()
	//time.Sleep(1000e9)
	<-done
}

//客户端必须知道服务器定义的对象类型和它的方法。它调用rpc.DialHTTP()去创建连接的客户端，当客户端被创建时，它可以通过client.Call("Type.Method", args, &reply)去调用远程的方法，其中Type与Method是调用的远程服务器被定义的类型和方法，args是一个类型的初始化对象，reply是一个变量，使用前必须要先声明它，它用来存储调用方法的返回结果。
////rpc_client.go
const serverAddress = "127.0.0.1"

func runClient(done chan bool) {
	client, err := rpc.DialHTTP("tcp", serverAddress+":8000")
	if err != nil {
		panic(err)
	}
	//synchronous call
	args := &Args{9, 10}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)
	//asynchronous call
	args2 := &Args{5, 7}
	call2 := client.Go("Args.Multiply", args2, &reply, nil)
	//replyCall := <- call2.Done
	<-call2.Done
	fmt.Printf("Args: %d * %d = %d\n", args2.N, args2.M, reply)
	done <- true
}

//main
func main() {
	done := make(chan bool)
	go runServer(done)
	runClient(done)
	/**
	Args: 9 * 10 = 90
	Args: 5 * 7 = 35
	*/
}
