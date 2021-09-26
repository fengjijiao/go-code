package main

//思考下面这个client-server配置：客户端无限循环执行从某个来源（可能是来自网络）接收的数据；数据使用一个Buffer类型的缓冲区读取。为了避免过多的分配和释放buffers，可以保留一个用缓冲通道表示的空闲列表：var freeList = make(chan *Buffer, 100)
//这个可以重复使用的缓冲队列与服务端共享。当客户端接收数据时，会尝试先从freeList获取一个buffer；如果freeList这个通道是空的，就分配一个新的buffer。当这个buffer被加载完，他就会通过serverChan发送给服务器端。
//var serverChan = make(chan *Buffer)
//下面是客户端的算法：
func client() {
	for {
		var b *Buffer
		select {
		case b= <-freeList:
			//获取到一个
		default:
			b = new(Buffer)
		}
		loadInto(b)//从网络中获取下一条
		serverChan <- b//发送给服务器端
	}
}
//服务端循环接收每一个客户端的消息，处理它，并尝试将buffer返回给共享的buffer列表：
func server() {
	for {
		b := <- serverChan//等待客户端发送一个buffer过来
		process(b)
		select {
		case freeList<-b:
			//插入freeList
		default:
			//freeList已满，会将buffer丢弃
		}
	}
}
//但是当freeList已满时它不能工作，这种情况下的缓冲区是:掉落到地上（因此命名为漏桶算法）被垃圾回收器回收。