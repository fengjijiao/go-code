package main

import "log"

//一个用到recover的程序（13.3）停掉了服务器内部一个失败的协程而不影响其他协程工作。
func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)//start the goroutine for that work
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("work failed with %s in %v\n", err, work)
		}
	}()
	do(work)
}
//上面的代码如果do(work)发生panic，错误会被记录且协程会退出并释放，而其他协程不受影响。
//因为recover总是返回nil，除非直接在defer修饰的函数中调用，defer修饰的代码可以调用那些自身可以panic和recover避免失败的库例程（库函数）。举例，safelyDo()中defer修饰的函数可能会在调用recover之前就调用了一个logging函数，panicking状态不会影响logging代码的运行。因为加入了恢复模式，函数do(以及它调用的任何东西)可以通过调用panic来摆脱不好的情况。但是恢复是在panicking的协程内部的：不能被另外一个协程恢复。
