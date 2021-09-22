package main

import "fmt"

//这是类型断言中的一个特例：假定v是一个值，然后我们想这样测试它是否实现了Stringer接口，可以这样做：
type Stringer interface {
	String() string
}

func main() {
	var v Stringer
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String())
	}
	//Print函数就是如此检测类型是否可以打印自身的。
	//接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。接口彻底将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质。
	//编写参数是接口变量的函数，这使得它们更具有一般性。
	//使用接口使代码更具有普适性。
}
