package main

import (
	"fmt"
	"sort"
)

//一个很好的例子来自标准库的sort包，要对一组数组或字符串排序，只需要实现三个方法：反映元素个数的Len()方法、比较第i和j个元素的Less(i, j)方法以及交换第i和j个元素的Swap(i, j) 方法。
//Sort函数接收一个接口类型参数：Sorter，它声明了这些方法：
//type Sorter interface {
//	Len() int
//	Less(i, j int) bool
//	Swap(i, j int)
//}
//现在我们要想对一个int数组排序，所有必须要做的事情就是：为数组定一个类型并在它上面实现Sorter接口的方法
type IntArray []int
func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//冒泡排序
//func Sort(data Sorter) {
//	for pass := 1; pass<data.Len();pass++ {
//		for i := 0;i<data.Len()-pass;i++ {
//			if data.Less(i+1, i) {
//				data.Swap(i, i+1)
//			}
//		}
//	}
//}
//
//func IsSorted(data Sorter) bool {
//	n := data.Len()
//	for i := n-1; i > 0; i-- {
//		if data.Less(i, i-1) {
//			return false
//		}
//	}
//	return true
//}
func main() {
	var data IntArray
	data = []int {9,6,4,1,11,3,7,0}
	sort.Sort(data)
	fmt.Println(data)
	//[0 1 3 4 6 7 9 11]
}
