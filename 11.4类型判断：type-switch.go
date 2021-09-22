package main

import "fmt"

func main() {
	//接口变量的类型可以用一种特殊形式的switch来检测：type-switch
	//switch t := areaIntf.(type) {
	//case *Square:
	//	fmt.Printf("Square type %T with value %v\n", t,t)
	//case *Circle:
	//	fmt.Printf("Circle type %T with value %v\n", t,t)
	//case nil:
	//	fmt.Printf("nil value: nothing to check?\n")
	//default:
	//	fmt.Printf("unexpected type %T\n", t)
	//}
	//变量t得到了areaIntf的值和类型，所有case语句中列举的类型(nil除外)都必须实现对应的接口（在上例中即Shaper），如果被检测类型没有在case语句列举的类型中，就会执行默认语句。
	//可以用type-switch运行进行时类型分析，但是在type-switch不允许有fallthrough。
	//如果仅仅是测试变量的类型，不用它的值，那么可以不需要赋值语句。
	//switch areaIntf.(type) {
	//case *Square:
	//	// TODO
	//case *Circle:
	//	//TODO
	//default:
	//	//TODO
	//}
	//下面的代码片段展示了一个类型分类函数classifier，他有一个可变长度参数，可以是任意类型的数组，他会根据数组原始的实际类型执行不同的动作。
	//在处理来自外部的、类型未知的数据时，比如解析诸如JSON或XML编码的数据，类型测试和转换会非常有用。

}

func classifier(items ...interface{}) {
	for i2, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("param %#d is a bool\n", i2)
		case float64:
			fmt.Printf("param %#d is a float64\n", i2)
		case int, int64:
			fmt.Printf("param %#d is a int\n", i2)
		case nil:
			fmt.Printf("param %#d is a nil\n", i2)
		case string:
			fmt.Printf("param %#d is a string\n", i2)
		default:
			fmt.Printf("param %#d is unknown\n", i2)
		}
	}
}
