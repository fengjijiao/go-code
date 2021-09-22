package main

import "fmt"

//Go语言不是一种“传统”的面向对象编程语言：它里面没有类和继承的概念。
//但Go语言里有非常灵活的接口概念，通过它可以实现很多面向对象的特性。接口提供了一种方式来说明对象的行为：如果谁能搞定这件事，它就可以用在这。
//接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
//通过以下格式定义接口。
//type Namer interface {
//	Method1(param_list) return_type
//	Method2(param_list) return_type
//}
//上面的Namer是一个接口类型。
//（按照约定，只包含一个方法的）接口的名字由方法名加[e]r后缀组成，例如Printer、Reader、Writer、Logger、Converter等等。还有一些不常用的方式（当后缀er不合适时），比如Recoverable，此时接口名以able结尾，或者以I开头（像.NET或JAVA中那样）。
//Go语言中的接口都很简短，通常它们会含有0个、最多3个方法。
//不像大多数面向对象编程语言，在Go语言中可以有值，一个接口类型的变量或一个接口值：var ai Namer，ai是一个多字（multiword）数据结构，它的值是nil。它本质上是一个指针，虽然不完全是一回事。指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误。

//...

//类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
//实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
//一个类型可以实现多个接口。
//接口类型可以包含一个实例的引用，该实例的类型实现了此接口（接口是动态类型）。
//即使接口在类型之后才定义，二者处于不同的包中，被单独编译：只要类型实现了接口中的方法，他就实现了此接口。
//所有这些特性使接口具有很大的灵活性。
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue()float32 {
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("value of asset is %f\n", asset.getValue())
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areIntf Shaper
	areIntf = sq1
	fmt.Printf("area: %f\n", areIntf.Area())
	//area: 25.000000
	//在main方法中创建了一个Square的实例。在主程序以外定义了一个接收者类型是Square方法的Area()，用来计算正方形面积：结构体Square实现了接口Shaper。
	//所以可以将一个Square类型的变量赋值给一个接口类型的变量：areIntf = sq1。
	//现在接口变量包含一个指向Square变量的引用，通过它可以调用Square上的方法Area()。当然也可以直接在Square的实例上调用此方法，但是在接口实例上调用此方法更让人兴奋，它使此方法更具有一般性。接口类型变量里包含了接收者实例的值和指向对应方法表的指针。
	//这是多态的Go版本，多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法，或者说：同一种类型在不同的实例上似乎表现出不同的行为。
	//如果Square没有实现Area()方法，编译器将会给出清晰的错误信息：
	//cannot use sq1(type *Square) as type Shaper in assignment: *Square does not implement Shaper(missing Area method)
	r := Rectangle{5,3}
	q := &Square{5}
	shapes := []Shaper{r,q}
	for _, shape := range shapes {
		fmt.Printf("detail: %v\n", shape)
		fmt.Printf("area: %v\n", shape.Area())
	}
	/**
	detail: {5 3}
	area: 15
	detail: &{5}
	area: 25
	*/
	//在调用shapes[n].Area()的时候，只知道shapes[n]是一个Shaper对象，最后它摇身一变成为了一个Square或Rectangle对象，并且表现出相应的行为。
	//也许从现在开始你将看到通过接口如何产生更干净、更简单、及更具有拓展性的代码。
	//下面是一个更具体的例子：有两个类型stockPosition和car，它们都有一个getValue()方法，我们可以定义一个具有此方法的接口valuable。接着定义一个使用valuable类型作为参数的函数showValue()，所有实现了valuable接口的类型都可以用这个函数。
	var o valuable = stockPosition{"goo", 57.9,5}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
	/**
	value of asset is 289.500000
	value of asset is 66500.000000
	*/
	
}
