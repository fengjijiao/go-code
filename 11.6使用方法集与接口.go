package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i < end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	var lst List
	if LongEnough(lst) {
		fmt.Println("lst is long enough")
	}
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Println("plst is long enough")
	}
	/**
	plst is long enough
	*/
	//讨论
	//在lst上调用CountInto时会导致一个编译器错误，因为CountInto需要一个Appender，而它的方法Append只定义在指针上。在lst上调用LongEnough是可以的，因为Len定义在值上。
	//在plst上调用CountInto是可以的，因为CountInto需要一个Appender，而它的方法Append定义在指针上。在plst上调用LongEnough也是可以的，因为指针会被自动解引用。

	//总结
	//在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型P直接可以辨识的：
	//1.指针方法可以通过指针调用
	//2.值方法可以通过值调用
	//3.接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	//4.接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
	//将一个值赋值给一个接口（intf = intfImpl）时，编译器会确保所有可能的接口方法都可以在此值上被调用(确保intfImpl有intf的实现)，因此不正确的赋值在编译期间就会失败。
	//

	//Go语言规范定义了接口方法集的调用规则：
	//1.类型*T的可调用方法集包含接收者为*T或T的所有方法集。
	//2.类型T的可调用方法集包含接收者为T的所有方法。
	//3.类型T的可调用方法中不包含接收者为*T的方法。
}
