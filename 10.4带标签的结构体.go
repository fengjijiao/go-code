package main

import (
	"fmt"
	"reflect"
)

//结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）:它是一个附属于字段的字符串，可以是文档或其他的重要标记。标签的内容不可以在一般的编程中使用，只有包reflect能获取到它。
//在11中将深入探讨reflect包，它可以在运行时自省类型、属性和方法，比如在一个变量上调用reflect.TypeOf()可以获取变量的正确类型，如果变量是一个结构体类型，就可以通过Field来索引结构体的字段，然后就可以使用Tag属性了。
type TagType struct {
	field1 bool "an imo"
	field2 string "are you"
	field3 int "yes"
}
func main() {
	tt := TagType{true, "ok", 123}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, i int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(i)
	fmt.Printf("tag: %v\n", ixField.Tag)
	/**
	tag: an imo
	tag: are you
	tag: yes
	*/
}