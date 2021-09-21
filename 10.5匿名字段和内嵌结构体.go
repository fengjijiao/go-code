package main

import "fmt"

//结构体中可以包含一个或多个匿名（或内嵌）字段，即使这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体。
//可以粗略的将这个和面向对象语言中的继承概念相比较，随后将会看到他被用来模拟类的继承的行为。Go语言中的继承是通过内嵌或组合来实现的，所以可以说，在Go语言中，相比较于继承，组合更受青睐。

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int//匿名字段
	innerS//anonymous field
}

func main() {
	//使用new
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.6
	outer.int = 60
	outer.in1 = 9
	outer.in2 = 99

	fmt.Printf("outer.b: %d\n", outer.b)
	fmt.Printf("outer.c: %f\n", outer.c)
	fmt.Printf("outer.int: %d\n", outer.int)
	fmt.Printf("outer.in1: %d\n", outer.in1)
	fmt.Printf("outer.in2: %d\n", outer.in2)

	//使用结构体字面量
	outer2 := outerS{6,7.5,60,innerS{9,99}}
	fmt.Printf("outer2: %v\n", outer2)

	/**
	outer.b: 6
	outer.c: 7.600000
	outer.int: 60
	outer.in1: 9
	outer.in2: 99
	outer2: {6 7.5 60 {9 99}}
	*/

	//通过类型outer.int的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：在一个结构体中对于每一种数据类只能有一个匿名字段。

	//10.5.2内嵌结构体
	//同样的结构体也是一种数据类型，所以它可以作为匿名字段来使用，如同上面的例子中的那样。外层结构体通过outer.in1直接进入内层结构体的字段，内嵌结构体甚至可以来自其他包。内层结构体被简单的插入或者内嵌进外层结构体。
	//这个简单的继承机制提供了一种方式，使得可以从另一个或一些类型继承部分或全部实现。

	type A struct {
		ax, ay int
	}
	type B struct {
		A
		bx, by float32
	}
	b := B{A{1,2},3.0,4.0}
	fmt.Println(b.ax,b.ay,b.bx,b.by)
	fmt.Println(b.A)
	/**
	1 2 3 4
	{1 2}
	*/

	//练习10.5 创建一个结构体，他有一个具名的float字段，2个匿名字段，类型分别为int,string。通过结构体字面量新建一个结构体实例并打印它的内容。
	type TestA struct {
		a float32
		int
		string
	}
	a0 := TestA{10.10,9,"ok"}
	fmt.Println(a0)
	/**
	{10.1 9 ok}
	*/

	//10.5.3命名冲突
	//当两个字段拥有相同的名字时（可能是继承来的名字）时该怎么办呢？
	//1.外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式。
	//2.如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没办法来解决这种问题引起的二义性，必须由程序员自己修正。
	//例子：
	//type A2 struct{a int}
	//type B2 struct {a, b int}
	//type C2 struct{A2;B2}
	//var c C2
	//使用c.a是错误的，到底是使用c.A2.a还是c.B2.a呢？会导致编译器错误：ambiguous NOT reference c.a disambiguate with either c.A2.a or c.B2.a。
	//type D2 struct {B2;b float32}
	//var d D2
	//使用d.b是没有问题的，他是float32，而不是B2的b。如果想要内层的b可以通过d.B2.b得到。
}
