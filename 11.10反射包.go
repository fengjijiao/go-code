package main

import (
	"fmt"
	"reflect"
)

//11.10.1方法和类型的反射
//反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。反射可以在运行时检查类型和变量，例如它的大小、方法和动态的调用这些方法。这对于没有源代码的包尤为有用。
//这是一个强大的工具，除非真的很有必要，否则应当避免使用或小心使用。
//变量的最基本信息就是类型和值：反射包的Type用来表示一个Go类型，反射包的Value为Go值提供了反射接口。
//两个简单的函数，reflect.TypeOf和reflect.ValueOf，返回被检查对象的类型和值。例如，x被定义为var x float64 = 3.4，那么reflect.TypeOf(x)返回float64，reflect.ValueOf(x)返回<float64 Value>
//实际上，反射是通过检查一个接口的值，变量首先被转换成空接口。这从下面的两个函数签名能够很明显的看出来。
//func TypeOf(i interface{}) Type
//func ValueOf(i interface{}) Value
//接口的值包含一个type和value。
//反射可以从接口值反射到对象，也可以从对象反射回接口值。
//reflect.Type和reflect.Value都有许多方法用于检查和操作它们。一个重要的例子是Value有一个Type方法返回reflect.Value的Type。另一个是Type和Value都有Kind方法返回一个常量来表示类型：Uint、Float64、Slice等等。同样Value有叫做Int和Float的方法可以获取存储在内部的值。(跟int和float64一样)
//...
//const(
//	Invalid Kind = iota
//	Bool
//	Int
// ...
//)
//对于float64类型的变量x，如果v:=reflect.ValueOf(x)，那么v.Kind()返回reflect.Float64，所以下面的表达式是true, v.Kind() == reflect.Float64
//Kind总是返回底层类型：
//type MyInt int
//var m MyInt = 5
//v := reflect.ValueOf(m)
//方法v.Kind()返回reflect.Int。
//变量v的Interface()方法可以得到还原（接口）值，所以可以这样打印v的值:fmt.Println(v.Interface())

func main() {
	type FI float64
	var x FI = 3.4
	fmt.Printf("type: %v\n", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())//底层类型
	fmt.Println("value: ", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(FI)
	//y := v.Interface().(float64)
	fmt.Println(y)
	/**
	type: main.FI
	value:  3.4
	type:  main.FI
	kind:  float64
	value:  3.4
	3.4
	value is 3.40e+00
	3.4
	*/
	//11.10.2通过反射修改（设置）值
	//假设我们要把x的值改为3.1415。Value有一些方法可以完成这个任务，但是必须小心使用：v.SetFloat(3.1415)。
	//这将产生一个错误：reflect.Value.SetFloat using unaddressable value。
	//问题的原因是v不是可设置的（这里并不是说值不可寻址）。是否可设置是Value的一个属性，并且不是所有的反射值都有这个属性：可以使用CanSet()方法测试是否可设置。
	//在例子中我们看到v.CanSet()返回false：settability of v: false.
	//当v := reflect.ValueOf(x)函数通过传递一个x拷贝创建了v，那么v的改变并不能更改原始的x。想要v的更改能作用到x，那就必须传递x的地址 v = reflect.ValueOf(&x)。
	//通过Type()我们看到v现在的类型是*float64并且仍然是不可设置的。
	//要想让其可设置我们需要使用Elem()函数，这间接的使用指针：v = v.Elem()
	//现在 v.CanSet()返回true并且v.SetFloat(3.1415)设置成功了。
	var x1 float64 = 3.9
	v1 := reflect.ValueOf(x1)
	fmt.Printf("settability of v: %t\n", v1.CanSet())
	v2 := reflect.ValueOf(&x)
	fmt.Printf("type of v: %v\n", v2.Type())
	fmt.Printf("settability of v: %t\n", v2.CanSet())
	v3 := v2.Elem()
	fmt.Printf("the elem of v is: %v\n", v3)
	fmt.Printf("settability of v: %t\n", v3.CanSet())
	v3.SetFloat(3.1415)
	fmt.Println(v3.Interface())
	fmt.Println(v3)
	/**
	settability of v: false
	type of v: *main.FI
	settability of v: false
	the elem of v is: 3.4
	settability of v: true
	3.1415
	3.1415
	*/
	//反射中有些内容是需要地址去改变它的状态的。
	//有时需要反射一个结构体类型。NumField()方法返回结构体内的字段数量；通过一个for循环用索引取得每个字段的值Field(i)。
	//我们同样能够调用签名在结构体上的方法，例如，使用索引n来调用：Method(n).Call(nil)。
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println(results)
	/**
	main.NotKnownType
	struct
	Field 0: ada
	Field 1: Go
	Field 2: Ob
	[ada-Go-Ob]
	*/
	//但是如果尝试更改一个值，会得到一个错误：
	//panic: reflect.Value.SetString using value obtained using unexported field
	//这是因为结构体中只有被导出字段（首字母大写）才是可设置的；
	t := T{23, "hello"}
	s := reflect.ValueOf(&t).Elem()
	typOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d：%s %s = %v\n", i, typOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(57)
	s.Field(1).SetString("ok")
	fmt.Println(t)
	/**
	0：A int = 23
	1：B string = hello
	{57 ok}
	*/
}

type T struct {
	A int
	B string
}

type NotKnownType struct {
	s1, s2, s3 string
}

func (n NotKnownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotKnownType{"ada", "Go", "Ob"}