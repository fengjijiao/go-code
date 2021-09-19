package main

import "fmt"

func main() {
	//这种构建方法可以应用于数组和切片
	slice1 := []int{-2,-1,0,1,2}
	for ix, value := range slice1 {
		fmt.Printf("%d：%d\n", ix, value)
	}
	/**
	第一个返回值ix是数组或切片的索引，第二个是在该索引位置的值；他们都是仅在for循环内部可见的局部变量。value只是slice1某个索引位置的值的一个拷贝，不能用来修改slice1该索引的值。
	 */
	//如果只需要索引，可以忽略第二个变量
	for ix := range slice1 {
		fmt.Printf("%d: %d\n", ix, slice1[ix])
	}
	//如果需要修改slice1[ix]的值可以使用这个版本
	//多维切片下的for-range
	/**
	通过计算行数和矩阵值可以很方便的写出for循环来，例如：
	 */
	//for row := range screen {
	//	for column := range screen[row] {
	//		screen[row][column] = 1
	//	}
	//}
	//问题7.5 假设我们有如下数组：items := [...]int{10,20,30,40,50}
	//a) 如果我们写了如下的for循环，那么执行完for循环后items的值是多少？
	//for _,item := range items {
	//    item *= 2
	//}
	items :=[...]int{10,20,30,40,50}
	for _, item := range items {
		item *= 2
	}
	fmt.Printf("%v\n", items)
	//b) 如果a)无法正常工作，写一个for循环让值可以double。
	for ix := range items {
		items[ix] *= 2
	}
	fmt.Printf("%v\n", items)
	//问题7.6通过使用省略号操作符...来实现累加。
	fmt.Printf("sum: %d\n", Adder(9, 10, 11))
	//练习7.7
	//a)写一个Sum函数，传入参数为一个float32数组成的数组arrF,返回该数组的所有数字和。
	arrF := [10]float32{0,1.1,2.2,3.3,4.4,5.5,6.6,7.7,8.8,9.9}
	fmt.Printf("sum: %f\n", Sum(arrF))
	//如果把数组修改为切片的话要做怎样的修改？
	sliceF := arrF[2:5]
	fmt.Printf("sum: %f\n", SumSlice(sliceF))
	//如果用切片形式方法实现不同长度数组的和呢？
	sliceF2 := arrF[3:6]
	fmt.Printf("sum: %f\n", SumSliceArr(sliceF, sliceF2))
	//写一个SumAndAverage方法，返回int和float32类型的未命名变量的和与平均值。
	r0, r1 := SumAndAverage(9, 9.9)
	fmt.Printf("sum: %f, average: %f\n", r0, r1)
	//练习7.8
	//写一个minSlice方法，传入一个int的切片并且返回最小值，再写一个maxSlice方法返回最大值。
	arr2 := [...]int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	slice2 := arr2[3:]
	fmt.Printf("min: %d\n", minSlice(slice2))
	fmt.Printf("max: %d\n", maxSlice(slice2))
}

func minSlice(p0 []int) int {
	min := 9999999
	for _, i3 := range p0 {
		if i3 < min {
			min = i3
		}
	}
	return min
}

func maxSlice(p0 []int) int {
	max := -9999999
	for _, i3 := range p0 {
		if i3 > max {
			max = i3
		}
	}
	return max
}

func SumAndAverage(p0 int, p1 float32) (float32, float32) {
	p02 := float32(p0)
	return p02 + p1, (p02 + p1) / 2
}

func SumSliceArr(arrF ...[]float32) float32 {
	var sum float32 = 0
	for _, f := range arrF {
		for _, i := range f {
			sum += i
		}
	}
	return sum
}

func SumSlice(arrF []float32) float32 {
	var sum float32 = 0
	for _, f := range arrF {
		sum += f
	}
	return sum
}

func Sum(arrF [10]float32) float32 {
	var sum float32 = 0
	for _, f := range arrF {
		sum += f
	}
	return sum
}

func Adder(params ...int) int {
	sum := 0
	for _, param := range params {
		sum += param
	}
	return sum
}
