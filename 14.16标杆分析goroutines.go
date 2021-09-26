package main

import (
	"fmt"
	"testing"
)

//在13.7中，提到了Go函数中的性能基准测试原则。在此我们将它应用于一个具体的范例之中：使用一个goroutine填充整数，然后再读取。测试中函数将被调用N次（N=1000000）。基准测试中，BenchMarkResult有一个String()方法用于输出结果。数值N由gotest决定，该值只有足够大才能判断出基准测试结果可靠合理。
//基准测试同样适用于普通函数。
//如果想排除一部分代码或者更具体的测算时间，你可以适当适用testing.B.StopTimer()和testing.B.StartTimer()来关闭或者开启计时器。
//只有所有测试全部通过，基准测试才会运行。

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {}
}

func main() {
	fmt.Println("sync ", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered ", testing.Benchmark(BenchmarkChannelBuffered).String())
	/**
	sync   4382287         267.3 ns/op
	buffered  16469331              70.36 ns/op
	*/
}