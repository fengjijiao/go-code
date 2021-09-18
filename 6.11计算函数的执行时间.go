package main

import (
	"fmt"
	"time"
)

/**
有时候，能够知道一个计算执行消耗的时间是非常有意义的，尤其是在对比和基准测试中。最简单的一个办法是在计算开始之前设置一个起始时间，在在计算结束后时的结束时间，最后求出它们的差值，就是这个计算所消耗的时间。
想要实现这样的做法，可以使用time包中Now()和Sub函数。
 */
func main() {
	start := time.Now()
	for i := 0; i < 30; i++ {
		//some code
		fmt.Printf("%d\n", i)
		//some code
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("this calculation took this amount of time: %s\n", delta)
	/**
	this calculation took this amount of time: 1.6445ms
	*/
}
