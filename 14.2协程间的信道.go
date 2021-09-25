package main

import (
	"fmt"
	"time"
)

//Go有一个特殊的类型，通道(channel)，像是通道，可以通过它们发送类型化的数据在协程之间通信，可以避开所有内存共享导致的坑；通道的通信方式保证了同步性。数据通过通道：同一时间只有一个协程可以访问数据：所以不会出现数据竞争，设计如此。数据的归属（可以读写数据的能力）被传递。
//通道服务于通信的两个目的：值的交换，同步的，保证了两个计算（协程）任何时候都是可知状态。

//通常使用这样的格式来声明通道：var identifier chan datatype
//未初始化的通道的值是nil
//所以通道只传输一种类型的数据，比如chan int或者chan string，所有的类型都可以用于通道，空接口interface{}也可以。甚至可以（有时非常有用）创建通道的通道。

//通道实际上是类型化的消息的队列：使数据得以传输。他是先进先出（FIFO）的结构所以可以保证发送给它们的元素的顺序（有些人知道，通道可以比作Unix Shells中的双向管道(two-way pipe)）。通道也是引用类型，所以我们使用make()函数来给它分配内存。这里先声明了一个字符串通道ch1，然后创建了它（实例化）：
var ch1 chan string

//ch1 = make(chan string)
//这里我们构建一个int通道：chanOfChans := make(chan int)
//或者函数通道: funcChan := chan func() , 14.17
//所以通道是对象的第一类型：可以存储在变量中，作为函数的参数传递，从函数返回以及通道发送它们自身。另外它们是类型化的，允许类型检查，比如尝试使用整数通道发送一个指针。

//14.2.2 通信操作符<-
//信息按照箭头的方向流动。
//流向通道（发送）
//ch <- int1 表示：用通道ch发送变量int1（双目运算，中缀 = 发送）
//从通道流出（接收），三种方式
//int2 = <- ch 表示变量int2从通道ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；假设int2已经声明过了，如果没有的话可以写成 int2 := <- ch。
//<- ch 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：
//if <- ch != 1000 {
//	//...
//}

func main2() {
	ch1 = make(chan string)
	go func() {
		for {
			if <-ch1 != "" {
				fmt.Println("ch1 receive non-empty string.")
			}
		}
	}()
	go func() {
		ch1 <- ""
		ch1 <- "ok"
	}()
	for {
	}
	/**
	ch1 receive non-empty string.
	exit status 0xc000013a
	*/
}

//操作符<-也被用来发送和接收，Go尽管不必要，为了可读性，通道的命名通常以ch开头或包含chan。通道的发送和接收操作都是自动的：它们通常一气呵成。

func main3() {
	ch1 = make(chan string)
	go sendData(ch1)
	go getData(ch1)
	for {

	}
	/**
	W
	A
	B
	C
	Q
	exit status 0xc000013a
	*/
}

func sendData(ch chan string) {
	ch <- "W"
	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "Q"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s\n", input)
	}
}

//注意：不要使用打印状态来表明通道的发送和接收顺序：由于打印状态和通道实际发送读写的时间延迟会导致和真实发送的顺序不同。

//14.2.3通信阻塞
//默认情况下，通信是同步的且无缓冲的：在有接收者接收数据之前，发送不会结束。可以想象一个无缓冲的通道在没有空间来保存数据的时候：必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。所以通道的发送/接收操作在对方准备好之前是阻塞的:
//1)对于同一个通道，发送操作（协程或函数中的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再通过通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待ch再次变为可用状态：就是通道值被接收时（可以传入变量）。
//2)对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
//尽管这看上去是非常严格的约束，实际在大部分情况下工作的很不错。
//以下程序验证了以上理论，一个协程在无限循环中给通道发送整数数据。不过因为没有接收者，只输出了一个数字0.
func main4() {
	ch2 := make(chan int)
	go pump(ch2)
	fmt.Println(<-ch2)
	/**
	0
	*/
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

//14.2.4 通过一个（或多个）通道交换数据进行协程同步
//通信是一种同步形式：通过通道，两个协程在通信（协程会和）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。
//甚至可以在通道两端互相阻塞对方，形成了叫做死锁的状态。Go运行时会检查并panic，停止程序。死锁几乎完全是由糟糕的设计导致的。
//无缓冲通道会被阻塞。设计无阻塞通道的程序可以避免这种情况，或者使用带缓冲的通道。

//练习14.2
func f1(in chan int) {
	fmt.Println(<-in)
}

func main5() {
	out := make(chan int)
	out <- 2
	go f1(out)
}

//解释以上程序为何会导致死锁？
//out通道首先发送2，等待接收者接收，但是接收者f1()此时还未初始化。

//14.2.5 同步通道 - 使用带缓冲的通道
//一个无缓冲通道只能包含一个元素，有时显得很局限。我们给通道提供了一个缓存，可以在扩展的make命令中设置它的容量，如下：
//buf := 100
//ch3 := make(chan string, buf)
//buf是通道可以同时容纳的元素（这里是string）个数
//在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，从而通道读取数据也不会阻塞，直到缓冲空了。
//缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要它们拥有同样的元素类型。内置的cap函数可以返回缓冲区容量。
//如果容量大于0 ，通信就是异步了：缓冲满载（发送）或者变空（接收）之前通信不会阻塞，元素会按照发送顺序被接收。如果容量是0或未设置，通信仅在收发双方准备好的情况下才可以成功。
//同步： ch := make(chan type, value)
//value ==0 ->synchronous, unbuffered(阻塞)
//value > 0 -> asynchronous, buffered(非阻塞)
//取决于value元素
//若使用通道的缓冲，你的程序将在“请求”激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性（scalable）。要在首要位置使用无缓冲通道来设计算法，只在不确定的情况下使用缓冲。
func main6() {
	out := make(chan int, 2)
	out <- 2
	go f2(out)
	out <- 3
	for {

	}
	/**
	2
	3
	exit status 0xc000013a
	*/
}

func f2(in chan int) {
	for {
		fmt.Println(<-in)
	}
}

//14.2.6 协程中用通道输出结果
//为了知道计算何时完成，可以通过信道回报。
//ch := make(chan int)
//go sum(bigArr, ch)
//sum := <- ch
//也可以使用通道来达到这个同步的目的，这个很有效的用法在传统计算机中称为信号量(semaphore)。或者换个方式：通过通道发送信号告知处理已经完成（在协程中）。
//在其他协程运行时让main程序等待协程完成，就是所谓的信号量模式。

//14.2.7信号量模式
//下面的片段阐明：协程通过在通道ch中放置一个值来处理结束的信号。main协程等待<- ch 直到从中获取到值。
//从ch获取返回结果
//func compute(ch chan int) {
//	ch <- someComputation()
//}
//
//func main() {
//	ch := make(chan int)
//	go compute(ch)
//	doSomethhingElseForAWhile()
//	result: <- ch
//}
//这个信号也可以是其他的，不返回结果，比如下面这个协程中的匿名函数(lambda)协程
//ch := make(chan int)
//go func() {
//	ch <- 1
//}
//doSomethhingElseForAWhile()
//<- ch
//或者等待两个协程完成
//done := make(chan bool)
//doSort := func(s []int) {
//	sort(s)
//	done <- true
//}
//i := pivot(s)
//go doSort(s[:i])
//go doSort(s[i:])
//<-done
//<-done
//下面的代码，用完整的信号量模式对长度为N的float64切片进行了N个doSomething()计算并同时完成，通道sem分配了相同的长度（且包含空接口类型的元素），待所有的计算都完成后发送信号（通过放入值）。在循环中从通道sem不停的接收数据来等待所有的协程完成。
//type Empty interface {}
//var empty Empty
////...
//data := make([]float64, N)
//res := make([]float64, N)
//sem := make(chan Empty, N)
////...
//for i, xi := range data {
//	go func (i int, xi float64) {
//		res[i] = doSomething(i, xi)
//		sem <- empty
//    }(i, xi)
//}
//for i := 0;i<N;i++{<-sem}

//14.2.8实现并行的for循环
//for循环的每一个迭代是并行完成的。
//for i, v := range data {
//	go func(i int, v float64) {
//		doSomething(i, v)
//		//...
//    }(i,v)
//}
//在for循环中并行计算迭代可能带来很好的性能提升。不过所有的迭代都必须是独立完成的。有些语言比如Fortress或者其他并行框架以不同的结构实现了这种方式，在Go中用协程实现起来非常容易：

//14.2.9用带缓冲通道实现一个信号量
//信号量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写问题，比如没有实现信号量的sync包，使用带缓冲的通道可以轻松实现：
//1.带缓冲通道的容量和要同步的资源容量相同
//2.通道的长度（当前存放的元素个数）与当前资源被使用的数量相同
//3.容量减去通道的长度就是未处理的资源个数（标准信号量的整数值）
//不用管通道中存放的是什么，只关注长度；因此我们创建了一个长度可变但容量为0（字节）的通道：
type Empty interface{}
type semaphore chan Empty

//将可用资源的数量N来初始化信号量
//sem := make(semaphore, N)//sem满了之后，因为缓冲区没有了空间所有不能再次写入数据，会被阻塞
//然后直接对信号量进行操作：
//请求N个资源
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

//释放N个资源
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

//可以用来实现一个互斥的例子
/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}
func (s semaphore) UnLock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) { //等待n个信号
	s.P(n)
}
func (s semaphore) Signal() { //发送一个信号
	s.V(1)
}

func main7() {
	N := 2
	sem := make(semaphore, N)
	res1 := make(chan int)
	res2 := make(chan int)
	intl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go calc(intl[:len(intl)/N], sem, res1)
	go calc(intl[len(intl)/N:], sem, res2)
	sem.Wait(N)
	res := <-res1 + <-res2
	fmt.Printf("res: %d\n", res)
	/**
	res: 45
	*/
}

func calc(data []int, sem semaphore, res chan int) {
	r := 0
	for _, datum := range data {
		r += datum
	}
	res <- r
	sem.Signal()
}

//通道工厂模式
//编程中常见的另外一种模式如下：不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）；函数内有个匿名函数被协程调用。
func main8() {
	stream := pump2()
	go suck2(stream)
	time.Sleep(1e9)
}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck2(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

//14.2.10给通道使用for循环
//for 循环的range语句可以用在通道ch上，便可以从通道中获取值，像这样：
//for v := range ch {
//	fmt.Printf("the value is %v\n", v)
//}
//它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很显然，另外一个协程必须写入ch（不然代码就阻塞在for循环了），而且必须在写入完成后才关闭。suck函数可以这样写，且在协程中调用这个动作，程序变成了这样：
func main9() {
	suck3(pump3())
	time.Sleep(1e9)
}

func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck3(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

//习惯用法：通道迭代模式
//这个模式用到了生产者-消费者模式，通常，需要从包含了地址索引字段items的容器给通道填入元素。为容器的类型定义一个方法Iter()，返回一个只读的通道items。
//func (c *container) Iter() <- chain items {//只读通道<- chain items
//	ch := make(chan item)
//	go func() {
//		for i := 0 ;i<c.Len();i++ {
//			ch <- c.items[i]
//        }
//	}()
//	return ch
//}

//在协程里，一个for循环迭代容器c中的元素（对于树或图的算法，这种简单的for循环可以替换为深度优先搜索）。
//调用这个方法的代码可以这样迭代容器：
//for x := range container.Iter(){ ... }

//可以运行在自己的协程中，所以上边的迭代用到了一个通道和两个协程（可能运行在两个线程上）。就有了一个特殊的生产者-消费者模式。如果程序在协程给通道写完值之前结束，协程不会被回收；设计如此。这种行为看起来是错误的，但是通道是一种线程安全的通信。在这种情况下，协程尝试写入一个通道，而这个通道永远不会被读取，这可能是个bug而并非期望它被静默回收。

//习惯用法：生成者-消费者模式
//假设你有Produce()函数来产生Consume函数需要的值。它们都可以运行在独立的协程中，生产者在通道中放入给消费者读取的值。整个处理过程可以替换为无限循环。
//for {
//	Consume(Produce())
//}

//14.2.11 通道的方向
//通道类型可以用注解来表示它只发送或者只接收：
//var send_only chan<- int
//var recv_only <-chan int
//只接收的通道（<-chan）无法关闭，因为关闭通道是发送者用来表示不再给通道发送值了，所以对只接收通道是没有意义的。通道创建的时候都是双向的，但也可以分配有方向的通道变量。
//var c = make(chan int)
//go source(c)
//go sink(c)
//func source(ch chan<- int) {
//	for {
//		ch <- 1
//	}
//}
//func sink(ch <-chan int) {
//	for {
//		<-ch
//	}
//}

//习惯用法：管道和选择器模式
//更具体的例子还有协程处理它从通道接收的数据并发送给输出通道
//sendChan := make(chan int)
//reciveChan := make(chan string)
//go processChannel(sendChan, reciveChan)
//
//func processChannel(in <-chan int, out chan<- string) {
//	for inValue := range in {
//		result := ...//processing inVlaue
//		out <- result
//	}
//}
//通过使用方向注解来限制协程对通道的操作。

//下面有一个很ok的例子，打印了输出的素数，使用选择器作为它的算法。每个prime都有一个选择器。
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main10() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

//协程filter(in, out chan int, prime int) 拷贝整数到输出通道，丢弃掉可以被prime整除的数字。然后每个prime又开启了一个新的协程，生成器和选择器并发请求。s

//第二个版本引入了上边的习惯用法：函数sieve、generate和filter都是工厂；它们创建通道并返回，而且使用了协程的lambda函数。main函数现在短小清晰：它调用sieve()返回了包含素数的通道，然后通过fmt.Println(<-primes)打印出来。
func generate2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {//2...
			ch <- i
		}
	}()
	return ch
}

func filter2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()//chan(2...n)
		for {
			prime := <-ch//2...n
			ch = filter2(ch, prime)//in = chan(2..n)?, prime =2...n
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
