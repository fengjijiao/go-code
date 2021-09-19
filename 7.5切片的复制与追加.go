package main

import "fmt"

func main() {
	//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。下面的代码描述了从拷贝切片的copy函数和向切片追加新元素的append函数。
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) //n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	//func append(s []T, x ...T) []T 其中append方法将0个或多个具有相同类型s的元素追加到切片后面并且返回新的切片；追加的元素必须和原切片的元素同类型。
	//如果s的容量不足以储存新增元素，append会分配新的切片来保证已有切片元素和新增元素的储存。因此，返回的切片可能已经指向一个不同的相关数组了。append方法总是返回成功，除非系统内存耗尽了。
	//如果你想将切片y追加到切片x的后面，只要将第二个参数扩展成一个列表即可：x = append(x, y...)
	//注意：append在大多数情况下很好用，但是如果你想完全掌控整个追加过程，你可以实现一个这样的AppendByte方法。

	//func copy(dst, src []T) int   copy方法将类型为T的切片从源地址src拷贝到目的地址dst，覆盖dst的相关元素，并且返回拷贝的元素个数。
	//源地址和目标地址可能会重叠。拷贝个数是src和dst的长度最小值。如果src是字符串那么元素类型就是byte。如果你还想继续使用src，在拷贝后执行src=dst。

	//练习7.9 给定 slice s[]int 和一个int类型的因子factor，扩展s使其长度为len(s) * factor
	s := make([]int, 3, 10)
	factor := 4
	fmt.Printf("len(s): %d\n", len(s))
	if len(s)*factor <= cap(s) {
		s = s[:len(s)*factor]
	} else {
		ns := make([]int, len(s)*factor)
		s = ns
	}
	fmt.Printf("len(s): %d\n", len(s))
	//练习7.10 用顺序函数过滤容器：s是前10个整型的切片。构造一个函数Filter，第一个参数是s，第二个参数是一个fn func(int) bool，返回满足函数fn的元素切片。通过fn测试当整型值是偶数时的情况。
	a0 := []int{1,2,3,4,5,6,7,8,9,10}
	a1 := Filter(a0, func(i2 int) bool {
		if i2 % 2 == 0 {
			return true
		}
		return false
	})
	fmt.Printf("%v\n", a1)
	//练习7.11 写一个函数InsertStringSlice将切片插入到另一个切片的指定位置。
	a2 := []int{1,2,3,4}
	a3 := []int{5,6,7,8}
	fmt.Printf("%v\n", InsertStringSlice(a2, a3, 2))
	//练习7.12 写一个函数RemoveStringSlice将从start到end索引的元素从切片中移除。
	a4 := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Printf("%v\n", RemoveStringSlice(a4, 3,5))
}

func RemoveStringSlice(src []int, start int, end int) (dst []int) {
	if end <= start {
		dst = src
		return
	}
	dst = append(dst, src[:start]...)
	dst = append(dst, src[end:]...)
	return
}

func InsertStringSlice(dst []int, src []int, pos int) (ns []int) {
	ns = append(ns, dst[:pos]...)
	ns = append(ns, src...)
	ns = append(ns, dst[pos:]...)
	return
}

func Filter(s []int, fn func(int) bool) []int {
	res := make([]int, 0, 10)
	i := 0
	for _, s1 := range s {
		if fn(s1) {
			res = res[:len(res)+1]
			res[i] = s1
			i++
		}
	}
	return res
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]
	copy(slice[m:n], data)
	return slice
}
