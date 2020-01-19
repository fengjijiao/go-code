package main
//函数内互相调用时变量作用范围变化
var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"//此处声明并初始化了局部变量，作用范围仅在f1内,若将:去除则改变全局变量a而非声明局部变量
	print(a)
	f2()
}

func f2() {
	print(a)
}
