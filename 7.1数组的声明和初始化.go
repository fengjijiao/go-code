package main

import (
	"fmt"
	"time"
)

func main() {
	a := [...]string{"a","b","c","d","e"}
	for i := range a {//输出索引
		fmt.Printf("%d\n", i)
	}
	for _, v := range a {//输出值
		fmt.Printf("%s\n", v)
	}
	//修改数组
	a[2] = "f"
	for _, v := range a {//输出值
		fmt.Printf("%s\n", v)
	}
	//在函数中修改原数组
	func(arr *[5]string) {
		arr[3] = "e"
	}(&a)//传递引用即可在函数中修改原数组
	fmt.Printf("%v\n", a)
	//练习7.1???证明当数组赋值时，发生了数组内存拷贝
	var b []int
	fmt.Printf("b%p\n", &b)
	b = []int{2,3,4,5,6}
	fmt.Printf("b%p\n", &b)
	b[2] = 9
	fmt.Printf("b%p\n", &b)
	//练习7.2，写一个循环并用下标给数组赋值（从0~15）并且将数组打印在屏幕上。
	var c [16]int
	for i := 0; i < 16; i++ {
		c[i] = i*2
		//fmt.Printf("c%p\n", &c)
	}
	fmt.Printf("%v\n", c)
	//练习7.3，通过数组计算斐波那契数，打印前50个
	var d [50]uint64
	start := time.Now()
	for i := 0; i < 50; i++ {
		if i<=1 {
			d[i] = 1
		}else {
			d[i] = d[i-1] + d[i-2]
		}
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("%dms\n", delta)
	fmt.Printf("%v\n", d)
	//数组常量初始化
	arrayKeyValue := [5]string{3:"ok",4:"finished!"}//0-2默认为空
	for _, v := range arrayKeyValue {
		fmt.Printf("%s\n", v)
	}
	arrayKeyValue2 := [10]int{1,2,3}
	for _, v := range arrayKeyValue2 {//其余默认为0
		fmt.Printf("%d\n", v)
	}
	//多维数组
	const (
		WIDTH = 1920
		HEIGHT = 1080
	)
	type pixel int
	var screen [WIDTH][HEIGHT]pixel
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	//将数组传递给函数
	//把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
	//1.传递数组的指针
	//2.传递数组的切片
	//方法一
	e := [3]float64{7.0,9.0,7.6}
	x := Sum(&e)
	fmt.Printf("the sum of the array is: %f\n", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}
