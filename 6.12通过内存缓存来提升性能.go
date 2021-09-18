package main

import (
	"fmt"
	"time"
)

/**
当进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。通过内存中缓存和重复利用相同的计算结果，称为内存缓存。最明显的例子就是生成斐波那契数列。
要计算数列中第n个数字，需要先得到之前两个数的值，但很明显绝大多数情况下前两个数的值都是已经计算过的。即每个更后面的数都是基于之前计算结果的重复计算。
而我们要做的就是将n个数的值存在数组中索引为n的位置，然后在数组中查找是否已经计算过，如果没有找到，则在进行计算。
 */
const LIM = 41
var fibs [LIM]uint64
func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("calculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res  uint64) {
	//memoization: check if fibonacci(n) is already known in array.
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	fibs[n] = res
	return
}
