package main

import (
	"fmt"
	"io"
)

//在经典的面向对象语言（像C++，Java和C#）中数据和方法被封装为类的概念：类包含它们两者，并且不能剥离。
//Go没有类：数据（结构体或更为一般的类型）和方法是一种松耦合的正交关系。
//Go中的接口跟Java/C#类似：都是必须提供一个指定方法集的实现。但是更加灵活通用：任何提供了此接口方法实现代码的类型都隐式地实现了该接口，而不用显式的声明。
//和其他语言相比，Go是唯一结合了接口值，静态类型检查（是否该类型实现了某接口），运行时动态转换的语言，并且不需要显式地声明类型是否满足某接口。该特性允许我们在不改变已有代码的情况下定义和使用新接口。
//接收一个（或多个）接口类型作为参数的函数，其实参可以是任何实现了该接口的类型。实现了某个接口的类型可以被传给任何以此接口为参数的函数。
//类似与Python和Ruby这类动态语言中的动态类型(duck typing)；这意味着对象可以根据提供的方法被处理（例如，作为参数传递给函数），而忽略它们的实际类型：它们能做什么比它们是什么更重要。

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 0; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	//
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}

func main() {
	b := new(Bird)
	DuckDance(b)
	/**
	I am quacking!
	I am walking!
	I am quacking!
	I am walking!
	I am quacking!
	I am walking!
	I am quacking!
	I am walking!
	*/
	//如果对cat调用函数DuckDance()，Go会提示编译错误，Python和Ruby会以运行时错误结束。

	//11.12.2 动态方法调用
	//像Python，Ruby这类语言，动态类型是延迟绑定的（在运行时进行）：方法只是用参数和变量简单地调用，然后在运行时才解析（它们很可能有像responds_to这样的方法来检查对象是否可以响应某个方法，但是这也意味着更大的编码量和更多的测试工作）
	//Go的实现与此相反，通常需要编译器静态检查的支持：当变量被赋值给一个接口类型的变量时，编译器会检查其是否实现了该接口的所有函数。如果方法调用作用于像interface{}这样的泛型上，你可以通过类型断言来检查变量是否实现了相应的接口。
	//例如：你用不同的类型表示XML输出流中的不同实体。然后我们为XML定义一个如下的写接口（甚至可以把它定义为私有接口）：
	//type xmlWriter interface {
	//	WriteXML(w io.Writer) error
	//}
	//现在我们可以实现适用于该流类型的任何变量的StreamXML函数，并用类型断言检查传入的变量是否实现了该接口；如果没有，我们就调用内建的encodeToXML来完成相应的工作：
	//func StreamXML(v interface{}, w io.Writer) error {
	//func encodeToXML(v interface{}, w io.Writer) error {
	//
	//Go在这里用了和gob相同的机制：定义了两个接口GobEncoder和GobDecoder。这样就允许类型自己实现从流编解码的具体方式；如果没有实现就使用标准的反射方式。
	//因此Go提供了动态语言的优点，却没有其他动态语言在运行时可能发生错误的缺点。
	//对于动态语言非常重要的单元测试来说，这样即可以减少单元测试的部分需求，又可以发挥相当大的作用。
	//GO的接口提高了代码的分离度，改善了代码的复用性，使得代码开发过程中的设计模式更容易实现。用Go接口还能实现依赖注入模式。

	//11.12.3接口的提取
	//提取接口是非常有用的设计模式，可以减少需要的类型和方法数量，而且不需要像传统的基于类的面向对象语言那样维护整个的类层次结构。
	//Go接口可以让开发者找出自己写的程序中的类型。假设有一些拥有共同行为的对象，并且开发者想要抽象出这些行为，这时就可以创建一个接口来使用。
	//假设我们需要一个新的接口TopologicalGenus，用来给shape排序（这里简单地实现为返回int）。我们需要做的是给想要满足接口的类型实现Rank()方法：
	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	for i2, _ := range shapes {
		fmt.Printf("shape detail: %v\n", shapes[i2])
		fmt.Printf("area of this shape: %f\n", shapes[i2].Area())
	}
	topgen := []TopologicalGenus{r, q}
	for i2, _ := range topgen {
		fmt.Printf("shape detail: %v\n", topgen[i2])
		fmt.Printf("topological genus of this shape: %f\n", topgen[i2].Rank())
	}
	/**
	shape detail: {5 3}
	area of this shape: 15.000000
	shape detail: &{5}
	area of this shape: 25.000000
	shape detail: {5 3}
	topological genus of this shape: %!f(int=2)
	shape detail: &{5}
	topological genus of this shape: %!f(int=1)
	*/
	//所有你不用提前设计出所有的接口：整个设计可以持续演进，而不用废弃之前的设定。类型要实现某个接口，它本身不用改变，只需要在这个类型上是实现新的方法。

	//11.12.4显式的指明类型实现了某个接口
	//如果你希望满足某个接口的类型显式的声明它们实现了这个接口，你可以向接口的方法中添加一个具有描述性名称的方法，例如：
	//type Fooer interface {
	//	Foo()
	//	ImplementsFooer()
	//}
	//类型Bar必须实现ImplementsFooer方法来满足Footer接口，以清楚的记录这个事实。
	//type Bar struct {
	//}
	//func (b Bar) ImplementsFooer() {}
	//func (b Bar) Foo(){}
	//大部分代码并不使用这样的约束，因为它限制了接口的实用性。
	//但是有些时候，这样的约束在大量相似的接口中被用来解决歧义。

	//11.12.5空接口和函数重载
	//函数重载是不被允许的。在Go语言中函数重载可以用可变参数...T作为函数最后一个参数来实现。如果我们把T换成空接口，那么可以知道任何类型的变量都是满足T(空接口)类型的，这样就允许我们传递任何数量任何类型的参数给函数，即重载的实际含义。
	//函数fmt.Printf就是这样做的
	//fmt.Printf(format string, ...interface{}) (n int, errno error)
	//这个函数会通过枚举slice类型的实参动态确定所有参数的类型。并查看每个类型是否实现了String()方法，如果是就用于产生输出信息。

	//11.12.6接口的继承
	//当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，这个类型就可以使用（另一个类型）所有的接口方法。
	//例如
	//type Task struct {
	//	Command string
	//	*log.Logger
	//}
	//这个类型的工厂方法像这样：
	//func NewTask(command string, logger *log.Logger) *Task {
	//	return &Task{command, logger}
	//}
	//当log.Logger实现了Log()之后，Task的实例task就可以调用该方法：
	//task.Log()
	//类型可以通过继承多个接口来提供像多重继承一样的特性:
	//type ReaderWriter struct {
	//	*io.Reader
	//	*io.Writer
	//}
	//上面概述的原理被应用于整个Go包，多态用得越多，代码就相对越少（12.8）。这被认为是Go编程中的重要的最佳实践。
	//有用的接口可以在开发的过程中被归纳出来。添加新接口非常容易，因为已有的类型不用变动（仅仅需要实现新接口的方法）。已有的函数可以扩展为使用接口类型的约束性参数：通常只有函数签名需要改变。对比基于类的OO类型的语言在这种情况下则需要适应整个类层此结构的变化。


	//练习11.11


}

type obj interface{}
type StringArr []string
type IntArr []int
func mapFunc(datas []obj) {
	var res string = ""
	for i2, _ := range datas {
		switch datas[i2].(type) {
		case string:
			res += datas[i2].(string)
			fmt.Println("int value * 2")
		case int:
			datas[i2] = datas[i2].(int) * 2
			fmt.Println("int value * 2")
		default:
			fmt.Println("unknown type")
		}
	}
	if res != "" {
		println(res)
	}
}

type Shaper interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Rank() int {
	return 1
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.length
}

func (r Rectangle) Rank() int {
	return 2
}

type xmlWriter interface {
	WriteXML(w io.Writer) error
}

func StreamXML(v interface{}, w io.Writer) error {
	if xw, ok := v.(xmlWriter); ok {
		return xw.WriteXML(w)
	}
	return encodeToXML(v, w)
}

func encodeToXML(v interface{}, w io.Writer) error {
	//...
	return nil
}