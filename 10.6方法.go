package main

import (
	"fmt"
	"math"
)

type Int uint64
type IntVector []int

//Go方法的作用是在接收者(receiver)上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
//接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是int\bool\string\或数组的别名类型。但是接收者不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现。如果这样做会引发一个编译错误：invaild receiver type...
//最后接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。
//一个类型加上它的方法等价于面向对象中的一个类。一个重要区别是：Go中允许类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是一个包的。
//因为方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法。但是如果基于接收者类型，是有重载的：具有同样名字的方法可以在2个或多个不同的接收者类型上存在，比如在同一个包里这么做是允许的：
//func (a *aStruct) Add(n int) int
//func (b *bStruct) Add(n int) int
//别名类型不能有它原始类型上已经定义过的方法
//定义方法的一般格式如下：
//func (recv receiver_type) methodName(parameter_list) (return_value_list){...}
//在方法名之前，func关键字之后的括号中指定receiver。
//如果recv是receiver的实例，Method1是它的方法名，那么方法的调用遵循传统的object.name选择器符号: recv.Method1()。
//如果recv是一个指针，Go会自动解引用。
//如果方法不需要使用recv的值，可以用_替换它，比如：
//func (_ receiver_type) methodName(parameter_list) (return_value_list) {...}
//recv就像面向对象语言中的this或self，但是Go中并没有这两个关键字。随个人喜好，你可以使用this或self作为receiver的名字。


//类型和作用在它上面的方法必须在同一个包里定义，这也是为什么不能在int\string或类似这些的类型上定义方法。试图在int类型上定义方法会得到一个编译错误：
//cannot define new methods on non-local type int
//
//但是有一个间接的方式：可以先定义该类（比如: int 或float）的别名类型，然后在为别名类型定义方法。或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
//type myTime struct { time.Time }
//func (t myTime) first3Chars() string {
//    return t.Time.String()[0:3]
//}
//初始化
//m := myTime{time.Now()}


//10.6.2函数和方法的区别
//函数将变量作为参数：Function1(recv)
//方法在变量上被调用：recv.Method1()
//在接受者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。
//receiver_type叫做（接收者）基本类型，这个类型必须在和方法同样的一个包中被声明。
//在Go中（接收者）类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接收者来建立。
//方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

//10.6.3指针或值作为接收者
//鉴于性能的原因，recv最常见的是一个指向receiver_type的指针（因为我们不想要一个实例的拷贝，如果按值调用的化就会是这样），特别是在receiver类型是结构体时，就更是如此了。
//如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法。否则，就在普通的值类型上定义方法。
//下面的例子，change()接受一个指向B的指针，并改变它内部的成员；write()接受通过拷贝接受B的值并只输出B的内容。
//注意: go为我们做了探测工作，我们自己并没有指出是否在指针上调用方法，Go替我们做了这些事情。b1是值而b2是指针，方法都支持运行了。
type B struct {
	thing string
}

func (b *B) change() { b.thing = "ok" }

func (b B) write() string { b.thing = "no";return fmt.Sprint(b) }

func main() {
	var b1 B //b1是值
	b1.change()
	fmt.Printf("%s\n", b1.write())
	fmt.Printf("%v\n", b1)

	b2 := new(B) //b2是指针
	b2.change()
	fmt.Printf("%s\n", b2.write())
	fmt.Printf("%v\n", b2)
	/**
	{no}
	{ok}
	{no}
	&{ok}
	*/
	//试着在write()中改变接收者的值；将看到它可以正常编译，但是开始的b没有被改变。
	//我们知道方法将指针作为接收者不是必须的，如下的例子，我们只是需要Point3的值来做计算：
	//...

	//这样做稍微有点昂贵，因为Point3的作为值传递给方法的，因此传递的是它的拷贝，这在Go中是合法的。也可以在指向这个类型的指针上调用此方法（会自动解引用）。
	//假设p3定义为一个指针： p3 := &Point{3,4,5}
	//可以使用p3.Abs()代替(*p3).Abs()。
	//在值和指针上调用方法：
	//可以有连接到类型的方法，也可以有连接到类型指针的方法。
	//但是这没关系：对于类型T，如果在*T上存在方法Meth()，并且t是这个类型的变量，那么t.Meth()会被自动转换为(&t).Meth()
	//指针方法和值方法都可以在指针或非指针上被调用，如下面的程序所示，类型List在值上有一个方法Len()，在指针上有一个方法Append()，但是可以看到两个方法可以在两种类型的变量上被调用。
	//值
	var l0 List
	l0.Append(1)
	fmt.Printf("%v , len: %d\n", l0, l0.Len())
	//指针
	l1 := new(List)
	l1.Append(2)
	fmt.Printf("%v , len: %d\n", l1, l1.Len())
	/**
	[1] , len: 1
	&[2] , len: 1
	*/

	//10.6.4方法和未导出字段
	//类型Person中存在一个未导出字段firstName，该如何在另一个包中修改或只是读取一个Person的firstname？
	//这可以通过面向对象语言一个众所周知的技术来完成：提供getter和setter方法。对于setter方法使用Set前缀，对于getter方法只使用成员名。
	//...

	//并发访问对象
	//对象的字段（属性）不应该由2个或2个以上的不同线程在同一时间去改变。如果在程序发生这种情况，为了安全并发访问，可以使用sync中的方法。
	//在14.17节中我们会通过goroutines和channels探索另一种方式。

	//10.6.5内嵌类型的方法和继承
	//当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型继承了这些方法：将父类型放在子类型中来实现亚型。这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果，也累是Ruby中的混入（mixin）。
	//
	//示例，假定有一个Engine接口类型，一个Car结构体类型，它包含一个Engine类型的匿名字段。
	//type Engine interface {
	//	Start()
	//	Stop()
	//}
	//type Car struct {
	//	Engine
	//}
	//我们可以构建如下的代码：
	//func (c *Car) GoToWorkIn() {
	//	c.Start()
	//	//drive to work
	//	c.Stop()
	//}
	n := &NamedPoint{Point{3,4}, "ok"}
	fmt.Println(n.Abs())//5
	//内嵌将一个已存在类型的字段和方法注入到了另一个类型里：匿名字段上的方法"晋升"成为了外层类型的方法。当然类型可以有只作用于本身实例而不作用于内嵌“父”类型上的方法。
	//可以覆写方法（像字段一样）：和内嵌类型方法具有同名的外层类型的方法会覆写内嵌类型对应的方法。
	//在NamedPoint上覆写Abs()
	//...
	//现在fmt.Println(n.Abs())会打印500。
	//因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承，就像：type Child struct { Father; Mother }。
	//结构体内嵌和自己同一个包中的结构体时，可以彼此访问对方所有的字段和方法。

	//10.8
	//创建一个上面Car和Engine可运行的例子，并且给Car类型一个wheelCount字段和一个numberOfWheels()方法。
	//创建一个Mercedes类型，它内嵌Car，并新建Mercedes的一个实例，然后调用它的方法。
	//然后仅在Mercedes类型上创建方法sayHiToMerkel()并调用它。
	var mercedes *Mercedes = new(Mercedes)
	mercedes.Start()
	fmt.Println("running")
	mercedes.Stop()

	mercedes.sayHiToMerkel()
	/**
	start!!!
	running
	stop!!!
	hello merkel!
	*/

	//10.6.6如何在类型中嵌入功能
	//主要有两种方式来实现在类型中嵌入功能：
	//A. 聚合（或组合）：包含一个所需功能类型的具名字段。
	//B. 内嵌：内嵌（匿名地）所需功能类型。
	//为了使这些概念具体化，假设有一个Customer类型，我们想让它通过Log类型来包含日志功能，Log类型只是简单地包含一个累积的消息（担任它可以是复杂的）。如果想让特定类型都具有日志功能，你可以实现一个这样的Log类型，然后将它作为特定类型的一个字段，并提供Log()，它返回这个日志的引用。
	//方式A可以通过如下方式实现（使用了10.7的String()功能）
	//
	c := new(Customer)
	c.Name = "Obama"
	c.log = new(Log)
	c.log.msg = "init"
	//c = &Customer{"Obama", &Log{"init"}}
	c.Log().Add("step 2")
	fmt.Println(c.Log())
	/**
	init
	step 2
	*/
	//方式B
	c2 := &Customer2{"Obama", Log{"init"}}
	c2.Add("step 2")
	c2.Add("step 3")
	fmt.Println(c2)
	/**
	Obama
	Log: {init
	step 2
	step 3}
	*/
	//内嵌的类型不需要指针，Customer(2)也不需要Add方法，它使用Log的Add方法，Customer2有自己的String方法，并且在它里面调用了Log的String方法。

	//10.6.7 多重继承
	//多重继承指的是类型获得多个父类型行为的能力，它在传统的面向对象语言中通常是不被实现的（C++和Python例外）。因为在类继承层次中，多重继承会给编译器引入额外的复杂度。但在Go中，通过在类型中嵌入所有必要的父类型，可以很简单的实现多重继承。
	//作为一个例子，假设有一个类型CameraPhone，通过它可以Call()，也可以TakeAPicture()，但是第一个方法属于类型Phone，第二个方法属于类型Camera。
	//只要嵌入这两个类型就可以解决问题
	cp := new(CameraPhone)
	fmt.Println(cp.TakeAPicture())
	fmt.Println(cp.Call())
	/**
	Click
	ring ring
	*/

	//练习10.10
	em := new(Employee)
	em.SetId(99)
	fmt.Printf("id: %d\n", em.Id())
	//练习10.11
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
	/**
	voodoo magic
	base magic
	base magic
	*/

	//10.6.8通用方法和方法命名
	//Open(),Close(),Read(),Write(),Sort()

	//总结
	//在Go中，类型就是类（数据和关联的方法）。Go拥有类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。
	//在Go中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫组件编程。

	//goop包中包含更多面向对象的能力。支持多重继承和类型独立分派，通过它可以实现你喜欢的其他编程语言里的一些结构。
}

type Base2 struct {}

func (Base2) Magic() {
	fmt.Println("base magic")
}

func (self Base2) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base2
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

type Employee struct {
	Person2
	salary float32
}

type Person2 struct {
	Base
	FirstName string
	LastName string
}

type Base struct {
	id int
}

func (b *Base) Id() int  {
	return b.id
}

func (b *Base) SetId(newId int) {
	b.id = newId
}

type Camera struct {}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct {}

func (p *Phone) Call() string {
	return "ring ring"
}

type CameraPhone struct {
	Camera
	Phone
}

type Customer2 struct {
	Name string
	Log
}

func (c *Customer2) String() string {
	return c.Name + "\nLog: "+fmt.Sprintln(c.Log)
}

type Log struct {
	msg string
}

func (l *Log) Add(s string)  {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

type Customer struct {
	Name string
	log *Log
}

type Engine interface {
	Start()
	Stop()
}

type FiEngine struct {
	Engine
}

func (f *FiEngine) Start() {
	fmt.Println("start!!!")
}

func (f *FiEngine) Stop() {
	fmt.Println("stop!!!")
}

type Car struct {
	FiEngine
	wheelCount int
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("hello merkel!")
}

type Point struct {
	x,y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}

type Person struct {
	firstName string
	lastName string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

type List []int
func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Point3 struct {x,y,z float64}
func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

//非结构体类型上方法的使用
//1
func (i2 Int) add(n Int) Int  {
	return i2 + n
}
//2
func (v IntVector) Sum() (s int) {
	for _, i3 := range v {
		s += i3
	}
	return
}
