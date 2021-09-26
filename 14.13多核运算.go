package main

import "runtime"

//假设我们的CPU核数是NCPU个：const NCPU = 4//4代表4核处理器，我们将计算划分为NCPU部分，每部分与其他部分并行运行。
const NCPU = 4
func DoAll() {
	sem := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}
	for i := 0; i < NCPU; i++ {
		<-sem
	}
	//全部完成
}

func DoPart(sem chan int) {
	//进行计算
	//...
	sem <- 1
}

func main() {
	runtime.GOMAXPROCS(NCPU)
	DoAll()
}

//函数DoAll()生成一个通道sem,在此基础上完成每一个并行计算；在for循环中启动NCPU个协程，每一个协程执行全部工作的1/NCPU。通过sem发送每一个协程中DoPart（）完成的信号。
//DoAll()中用一个for循环来等待所有（NCPU个）协程完成计算：通道sem的行为就像一个semaphore（信号量）(14.2.7)；

//在此前的运行模式下，你还必须设置GOMAXPROCS为NCPU.