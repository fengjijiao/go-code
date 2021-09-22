package main

import (
	"fmt"
	"math"
)

//一个接口变量varI中可以包含任何类型的值，必须有一种方式来检测它的动态类型，即运行时在变量中存储的值的实际类型。在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。通常我们可以用类型断言来测试在某时刻varI是否包含类型T的值。
//v := varI.(T)//unchecked type assertion
//varI必须是一个接口变量，否则编译器会报错: invalid type assertion: varI.(T) (non-interface type (type of varI) on left)
//类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言：
//if v,ok :+ varI.(T); ok {//checked type assertion
//	Process(v)
//	return
//}
////varI is not of type T
//如果转换合法，v是varI转换到类型T的值，ok会是true；否则v是类型T的零值，ok是false，也没有运行时错误发生。
//应该总是使用上面的方式来进行类型断言。
//多数情况下，我们可能只是想在if中测试一下ok的值，此时使用以下的方法会是最方便的：
//if _,ok := varI.(T); ok {
//	//...
//}
type Square struct {
	side float32
}

func (s Square) Area() float32 {
	return s.side * s.side
}

func (s Circle) Area() float32 {
	return s.radius * s.radius * math.Pi
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	//Is Square the type of areIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("the type of areaIntf is: %T\n", t)
	}
	if t, ok := areaIntf.(*Circle); ok {
		fmt.Printf("the type of areaIntf is: %T\n", t)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle.")
	}
	/**
	the type of areaIntf is: *main.Square
	areaIntf does not contain a variable of type Circle.
	*/
	//备注：如果忽略areaIntf.(*Square)中的*号，会导致编译错误：impossible type assertion: Square does not implement Shaper(Area method has pointer receiver).
}
