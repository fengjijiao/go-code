package main

import "fmt"

//11.9.1概念
//空接口或最小接口 不包含任何方法，他对实现不做任何要求：
type Any interface{}
//任何其他类型都实现了空接口（它不仅仅像Java/C#中的Object引用类型），any或Any是空接口一个很好的别名或缩写。
//空接口类似Java/C#中所有类的基类：Object类，二者的目标也很接近。
//可以给一个空接口类型的变量var val interface{}赋任何类型的值。
var i = 5
var str = "ABC"

type Person struct {
	name string
	age int
}


type Element interface{}

type Vector struct {
	a []Element
}

func (v *Vector) At(i int) Element {
	return v.a[i]
}

func (v *Vector) Set(i int, e Element) {
	v.a[i] = e
}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.age = 55
	pers1.name = "Rob"
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	/**
	val has the value: 5
	val has the value: ABC
	val has the value: &{Rob 55}
	Type pointer to Person *main.Person
	*/
	//在上面的例子中，接口变量val被依次赋予一个int，string和Person实例的值，然后使用type-switch来测试它的实际类型。每个interface{}变量在内存中占据两个字长：一个用来存储它包含的类型，另一个用来存储它包含的数据或指向数据的指针。
	//testFunc := func(any interface{}) {
	//	switch v := any.(type) {

	//11.9.2构建通用类型或包含不同类型变量的数组
	//int数组\float数组\string数组能被搜索和排序，那么对于其他类型的数组呢，是不是我们必须得自己编程实现它们呢？
	//现在我们知道该怎么做了，就是通过使用空接口。让我们给空接口定一个别名类型Element： type Element interface{}
	//然后定义一个容器类型的结构体Vector，它包含一个Element类型元素的切片。
	//...
	//Vector里能放任何类型的变量，因为任何类型都实现了空接口，实际上Vector里面放的每个元素可以是不同类型的变量。我们为它定义一个At()方法用于返回第i个元素：
	//...
	//再定一个Set()方法用于设置第i个元素的值。
	//...
	//Vector中存储的所有元素都是Element类型，要得到它们的原始类型（unboxing:拆箱）需要用到类型断言。TODO:the compiler rejects assertions guaranteed to fail，类型断言总书在运行时才会执行，因此它会产生运行时错误。


	//练习11.10
	d0 := IntArray{1,2,3,4,-8,5,6,9,5,11,99,22}
	fmt.Printf("minimum value is %d.\n", Min(d0))
	//minimum value is -8.

	//11.9.3复制数据切片至空接口切片
	//假设你有一个myType类型的数据切片，你想将切片中的数据复制到一个空接口切片中，类似：
	//var dataSlice []myType = FuncReturnSlice()
	//var interfaceSlice []interface{} = dataSlice
	//可惜不能这么做，编译时会报错：cannot use dataSlice(type []myType) as type []interface{} in assignment.
	//原因是它们两在内存中的布局是不一样的。
	//必须使用for-range语句来一个一个显式的复制：
	//var dataSlice []myType = FuncReturnSlice()
	//var interfaceSlice []interface{} = make([]interface, len(dataSlice))
	//for i, d := range dataSlice{
	//	interfaceSlice[i] = d
	//}


	//11.9.4通用类型的节点数据结构
	//在10.1中我们遇到了诸如列表和树这样的数据结构，在它们的定义中使用了一种叫节点的递归结构体类型，节点包含一个某种类型的数据字段。现在可以使用空接口作为数据字段的类型，这样我们就能写出通用的代码。
	//下面是一个二叉树的部分代码：通用定义、用于创建空节点的NewNode方法，及设置数据的SetData方法。
	root := NewNode(nil,nil)
	root.SetData("root node")
	//make child nodes.
	a := NewNode(nil, nil)
	a.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.le = a
	root.ri = b
	fmt.Printf("%v\n", root)
	/**
	&{0xc00004c400 root node 0xc00004c420}
	*/


	//接口到接口
	//一个接口的值可以赋值给另一个接口变量，只要底层类型实现了必要的方法。这个转换是在运行时进行检测的，转换失败会导致一个运行时错误：这是Go语言动态的一面，可以拿它和Ruby和Python这些动态语言相比较。
	//假定：
	//var ai AbsInterface//declares method Abs()
	//type SqrInterface interface {
	//	Sqr() float32
	//}
	//var si SqrInterface
	//pp := new(Point)//say *Point implements Abs,Sqr
	//var empty interface{}
	//那么下面的语句和类型断言是合法的：
	//empty = pp//ererything satisfies empty
	//ai = empty.(AbsInterface)//underlying value pp implements Abs()
	////runtime failure otherwise
	////si = ai.(SqrInterface)//*Point has Sqr() even though AbsInterface does not
	//empty = si//*Point implements empty set
	////Note: statically checkable so type assertion not necessary.
	//下面是函数调用的一个例子
	//type myPrintInterface interface {
	//	print()
	//}
	//func f3(x myInterface) {
	//	x.(myPrintInterface).print()
	//}
	//x转换为myPrintInterface类型是完全动态的：只要x的底层类型（动态类型）定义了print方法这个调用就可以正常运行。
}

type Miner interface {
	min() int
	max() int
}

type IntArray []int

func (i2 IntArray) min() (res int) {
	res = 999999
	for _, i4 := range i2 {
		if i4 < res {
			res = i4
		}
	}
	return
}

func (i2 IntArray) max() (res int) {
	res = -999999
	for _, i4 := range i2 {
		if i4 > res {
			res = i4
		}
	}
	return
}

func Min(m Miner) int {
	return m.min()
}



type Node struct {
	le *Node
	data interface{}
	ri *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}