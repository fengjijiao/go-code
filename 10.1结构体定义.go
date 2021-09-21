package main

import "fmt"

/**
结构体定义的一般方式如下
type identifier struct {
	field1 type1
	field2 type2
	field3 type3
}
type T struct {a,b int}也是合法的语法，它更适用于简单的结构体。
结构体中的字段都有名字，例如field1等，如果字段在代码中从来也不会被用到，那么可以命名为_。
结构体的字段可以是任意类型，甚至可以是结构体本身(10.5)，也可以是函数或者接口(11)。
可以声明结构体类型的一个变量，然后像下面这样给它的字段赋值。
var s T
s.a = 5
s.b = 6
数组可以看作是一种结构体类型，不过它使用下标而不是具名的字段。

使用new
使用new函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T) ==> var t *T = &T{} ，如果需要可以把这条语句放在不同的行（比如定义是包范围的，但是分配却没有必要在开始就做）。
var t *T
t=new(T)
写这条语句的惯用方法是：t := new(T)，变量t是一个指向T的指针，此时结构体字段的值是它们所属类型的零值。
声明 var t T也会给t分配内存，并零值化内存，但是这个时候t是类型T。在这两种方式中，t通常被称作类型T的一个实例（instance）或对象（object）。
*/
type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "ok"
	fmt.Printf("int: %d\n", ms.i1)
	fmt.Printf("float: %f\n", ms.f1)
	fmt.Printf("str: %s\n", ms.str)
	fmt.Println(ms)
	/**
	int: 10
	float: 15.500000
	str: ok
	&{10 15.5 ok}
	*/
	//就像前面在面向对象语言所作的那样，可以使用点号符给字段赋值：structname.fieldname = value。
	//同样地，使用点号符可以获取结构体字段的值：structname.fieldname。
	//在Go语言中这叫选择器（selector）。无论变量是结构体类型还是结构体类型指针，都使用同样的选择器符（selector-notation）来引出结构体的字段。
	//type struct2 struct {
	//	i int
	//}
	//var v struct2
	//var p *struct2
	//v.i
	//p.i
	//初始化一个结构体实例（一个结构体字面量：struct-literal）的更简短和惯用的方式如下：
	//ms := &struct1{10,15.5,"ok"}
	//或者
	//var ms struct1
	//ms = struct1{10,15.5,"ok"}
	//混合字面量语法(composite literal syntax) &struct1{a,b,c}是一种简写，底层仍然会调用new()，这里值的顺序必须按照字段顺序来写。在下面的例子中能看出可以通过在值的前面放上字段名来初始化字段的方式。
	//表达式new(Type)和&Type{}是等价的.
	//可以直接通过指针，像p.i=12这样给结构体字段赋值，没有像c++中那样需要使用->操作符，Go会自动做这样的转换。
	//也可以通过解指针的方式来设置值：(*p).i=12。

	//结构体的内存布局
	//Go语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他结构体，这在性能上带来了很大的优势。不像java中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，这点和Go语言中的指针很像。

	//结构体转换
	//Go中的类型转换遵循严格的规则。当为结构体定义了一个alias类型时，此结构体类型和它的alias类型都有相同的底层类型，它们可以互相转换a1 := a(b1);b1 := b(a1)，同时需要注意其中非法赋值或转换引起的编译错误。

	//练习10.1 定义结构体Address和VCard，后者包含一个人的名字、地址编号、出生日期和图像，试着选择正确的数据类型。构建一个自己的vcard并打印出来。
	//提示：VCard必须包含住址，它应该以值类型还是指针类型放在VCard中呢？
	//第二种会好点，因为它占用内存少。包含一个名字和两个指向地址的指针的Address结构体可以使用%v打印

	type Address struct {
		Value string
	}

	type VCard struct {
		Name string
		Addr *Address
		Addr2 *Address
	}

	v := new(VCard)
	v.Addr = &Address{
		Value: "funny",
	}
	v.Addr2 = &Address{
		Value: "hanny",
	}
	v.Name = "follow"

	fmt.Printf("%v\n", v)//&{follow 0xc00010e120 0xc00010e130}

	//练习10.2
	type Person struct {
		Name string
	}
	fn1 := func(per *Person) {
		per.Name = "123"
	}
	fn2 := func(per Person) {
		per.Name = "123"
	}
	per1 := &Person{
		Name: "ok",
	}
	per2 := &Person{
		Name: "ok",
	}
	fn1(per1)
	fn2(*per2)
	fmt.Printf("per1: %v\n", per1)
	fmt.Printf("per2: %v\n", per2)
	/**
	per1: &{123}
	per2: &{ok}
	*/

	//练习10.4...
}
