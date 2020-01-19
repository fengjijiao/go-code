package main

func main() {
	var a int
	var b int32
	a = 15
	b = a + a // 编译错误 int8+int8不能自动转换类型为int32
	b = b + 5
}
