package main

import (
	"bytes"
	"fmt"
)

/**
切片是对数组一个连续切片的引用（该数组我们称为相关数组，通常是匿名的），所以切片是一个引用类型。
切片提供一个相关数组的动态窗口。
和数组不同的是，切片的长度可以在运行时修改，最小为0最大为相关数组的长度，切片是一个长度可变的数组。
cap()获取切片最长可以达到多少，他等于切片从第一个元素开始，到相关数组的末尾元素个数。对于切片s，以下不等式永远成立：0<=len(s)<=cap(s)

多个切片如果表示同一个数组的片段，它们可以共享数据；因此一个切片和相关数据的其他切片是共享存储的。
声明切片的格式为：var identifier []type（不需要说明长度）
一个切片在未初始化之前默认为nil，长度为0
切片初始化的标准格式为：var slice1 []type = arr1[start:end]
这表示slice1是数组arr1从start索引到end-1索引之间的元素构成的子集（切分数组，start:end被称为slice表达式）。所以slice1[0]就等于arr1[start]。这可以在arr1被填充前就定义好。
var slice1 []type = arr1[:] == slice1 = &arr1
一个由数字1、2、3组成的切片可以这么生成，s := [3]int{1,2,3}[:]或s := []int{1,2,3}

一个切片可以这样扩展到它的大小上限： s = s[:cap(s)]

 */
func main() {
	arr := []int{1,2,3,4,5,6,7,8,9}
	var slice []int = arr[2:6]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
	fmt.Print("\n")
	slice[1] = 9
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
	fmt.Print("\n")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Print("\n")
	/*
	3 4 5 6
	3 9 5 6
	1 2 3 9 5 6 7 8 9
	*/
	/**
	new()和make()的区别
	看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。
	1.new(T)为每个新的类型T分配一片内存，初始化为0并且返回类型为*T的内存地址；这种方法返回一个指向类型为T，值为0的地址的指针，它适用于值类型如数组和结构体（见第10章）；它相当于&T{}
	2.make(T)返回一个类型为T的初始值，它只适用于3种内建的引用类型：切片、map和channel。
	换言之，new函数分配内存，make函数初始化。
	 */
	var p0 *[]int = new([]int)
	fmt.Printf("%d, len: %d, cap: %d\n", *p0, len(*p0), cap(*p0))//*p == nil
	p1 := new([]int)
	fmt.Printf("%d, len: %d, cap: %d\n", *p1, len(*p1), cap(*p1))
	/**
	[], len: 0, cap: 0
	[], len: 0, cap: 0
	*/
	//以上两种方式的实用性都不高。
	var v0 []int = make([]int, 10, 50)
	fmt.Printf("%d, len: %d, cap: %d\n", v0, len(v0), cap(v0))
	v1 := make([]int, 10, 50)
	fmt.Printf("%d, len: %d, cap: %d\n", v1, len(v1), cap(v1))
	/**
	[0 0 0 0 0 0 0 0 0 0], len: 10, cap: 50
	[0 0 0 0 0 0 0 0 0 0], len: 10, cap: 50
	*/
	//这样分配一个有50个int值的数组，并且创建了一个长度为10，容量为50的切片v，该切片指向数组的前10个元素
	s0 := make([]byte, 5)
	fmt.Printf("%d, len: %d, cap: %d\n", s0, len(s0), cap(s0))
	s1 := s0[2:4]
	fmt.Printf("%d, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	/**
	[0 0 0 0 0], len: 5, cap: 5
	[0 0], len: 2, cap: 3
	*/
	s2 := []byte{'p','o','e','m'}
	fmt.Printf("s2: %v\n", s2)
	s3 := s2[2:]
	fmt.Printf("s3: %v\n", s3)
	s3[1] = 't'
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	//多维切片
	/**
	和数组一样，切片通常也是一维的，但是也可以由一维组合成高维。通过分片的分片（或者切片的数组），长度可以任意动态变化，所以Go语言中的多维切片可以任意切分。而且，内层的切片必须单独分配（通过make函数）。
	 */
	//byte包
	/**
	类型[]byte的切片十分常见，Go语言有一个bytes包专门用来解决这种类型的操作方法。
	bytes包和字符串十分类似。而且他还包含一个十分有用的类型Buffer.
	Buffer提供Read和Write方法，因此读写未知长度的bytes最好使用buffer.
	buffer可以这样定义：var buffer bytes.Buffer。
	或者使用new获取一个指针：var r *bytes.Buffer = new(bytes.Buffer)。
	或者通过函数：func NewBuffer(buf []byte) *Buffer，创建一个Buffer对象并且用buf初始化好；NewBuffer最好用在从buf读取的时候使用。
	 */
	/**
	使用buffer串联字符串
	类似于java上的StringBuilder类
	创建一个buffer，通过buffer.WriteString(s)方法将字符串s追加到后面，最后再通过buffer.String()方法转换为string。
	 */
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		}else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")
	/**
	这种实现方式比+=更加节省内存和CPU,尤其是要串联的字符串数目特别多的时候。
	 */
	//练习1，给定切片sl，将一个[]byte数组追加到sl后面，写一个函数Append(slice, data []byte)[]byte，该函数在sl不能存储更多数据时自动扩容
	fmt.Print("len(data)<cap(slice)\n")
	r2 := []byte{0x02,0x04,0x06,0x08,0x10,0x12,0x14,0x16,0x18}
	r0 := []byte{0x00, 0x01, 0x02}
	fmt.Printf("r0: %v\n", r0)
	r1 := r2[2:4]
	fmt.Printf("r1: %v\n", r1)
	r1 = Append(r1, r0)
	fmt.Printf("r1: %v\n", r1)
	fmt.Print("len(data)>cap(slice)\n")
	r2 = []byte{0x02,0x04,0x06,0x08,0x10,0x12,0x14,0x16,0x18}
	r0 = []byte{0x00, 0x01, 0x02, 0x03,0x04,0x05,0x06,0x07}
	fmt.Printf("r0: %v\n", r0)
	r1 = r2[2:4]
	fmt.Printf("r1: %v\n", r1)
	r1 = Append(r1, r0)
	fmt.Printf("r1: %v\n", r1)
}

var i int = 0
func getNextString() (string, bool) {
	i++
	if i < 20 {
		return "ok", true
	}
	return "", false
}

func Append(slice []byte, data []byte) []byte {
	scap := cap(slice)
	slen := len(slice)
	dlen := len(data)
	if scap - slen >= dlen {
		slice = slice[:slen+dlen]
		for i2, datum := range data {
			slice[slen+i2] = datum
		}
		return slice
	}else {
		narr := make([]byte, slen + dlen)
		copy(narr, slice)
		for i2, datum := range data {
			narr[slen+i2] = datum
		}
		return narr
	}
}