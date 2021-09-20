package main

import (
	"bytes"
	"fmt"
	"sync"
)

type Info struct {
	mu sync.Mutex
	s string
}

type SyncedBuffer struct {
	lock sync.Mutex
	buffer bytes.Buffer
}

func main() {
	//在一些复杂的程序中，通常通过不同线程执行不同应用来实现应用程序的并发。当不同线程要使用同一个变量时，经常会出现一个问题：无法预知变量被不同线程修改的顺序（这通常被称为资源竞争，指不同线程对同一变量使用的竞争）显然这无法让人忍受，那么我们该如何解决这个问题呢？
	//经典的做法是一次只能让一个线程对共享变量进行操作。当变量被一个线程改变时（临界区），我们为它上锁，直到这个线程执行完成并解锁后，其他线程才能访问它。
	//特别是我们之前章节学习的map类型是不存在锁的机制来实现这种效果的（处于对性能的考虑），所有map类型是非线程安全的。当并行访问一个共享的map类型的数据，map数据将会出错。
	//在Go语言中这种锁的机制是通过sync包中Mutex来实现的。sync来源于“synchronized”一词，这意味着线程将有序的对同一变量进行访问。
	//sync.Mutex是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
	//假设info是一个需要上锁的放在共享内存中的变量。通过包含Mutex来实现的一个典型例子如下：
	//如果一个函数想要改变这个变量可以这样写
	var i0 = &Info{s: "bbb"}
	fmt.Printf("%s\n", i0.s)
	Update(i0)
	fmt.Printf("%s\n", i0.s)
	//还有一个非常有用的例子是通过Mutex来实现一个可以上锁的共享缓冲器：type SyncedBuffer struct
	//在sync包中还有一个RWMutex锁，他能通过RLock()来允许同一时间多个进程对变量进行读操作，但是只能一个进程进行写操作。如果使用Lock()将和普通的Mutex作用相同。
	//包中还有一个方便的Once类型变量的方法once.Do(call)，这个方法确保被调用函数只能被调用一次。
	//相对简单的情况下，通过使用sync包可以解决同一时间只能一个线程访问变量或map类型数据的问题。如果这种方式导致程序明显变慢或者引起其他问题，我们要重新思考来通过goroutines和channels来解决问题,这是在go语言中所提倡用来实现并发的技术。
}

func Update(info *Info) {
	info.mu.Lock()
	info.s = "aaa"
	info.mu.Unlock()
}
