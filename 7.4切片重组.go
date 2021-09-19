package main

import "fmt"

func main() {
	//我们已经知道切片创建的时候通常比相关数组小，例如：
	//slice1 :+ make([]type, start_length, capacity)
	//其中start_length作为切片的初始长度而capacity作为相关数据的长度。
	//这么做的好处是我们的切片在达到容量上限后可以扩容。改变切片长度的过程称之为切片重组reslicing，做法如下：slice1 := slice1[0:end]，其中end是新的末尾索引（即长度）。
	//将切片扩展1位可以这么做：
	//s1 = s1[0:len(s1)+1]
	//切片可以反复扩展直至占据整个相关数组。
	slice1 := make([]int, 0, 50)
	for i:= 0; i<cap(slice1);i++ {
		slice1 = slice1[0:i+1]
		slice1[i] = i
		fmt.Printf("the length of slice is %d\n", len(slice1))
	}
	fmt.Printf("%v\n", slice1)
	//另一个例子
	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	var a = ar[5:7]
	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))//2,4
	//将a重新分片
	a = a[0:4]
	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))//4,5 cap(a)依然是5
	//问题7.7
	//1)如果s是一个切片，那么s[n:n]的长度是多少？
	var s = ar[5:7]
	fmt.Printf("len: %d\n", len(s[3:3]))//0
	//2)s[n:n+1]的长度又是多少？
	var s1 = ar[5:7]
	fmt.Printf("len: %d\n", len(s1[3:4]))//1
}