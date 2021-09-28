package main

import "fmt"

func main() {
	//1.如何检测一个值v是否实现了接口Stringer
	if v, ok := v.(Stringer); ok {
		//实现了
	}
	//2.如何使用接口实现一个类型分离函数
	//func classifier(items ...interface{})
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is a int\n", i)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d's type is a unknown\n", i)
		}
	}
}
