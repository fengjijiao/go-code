package main


type struct1 struct {
	field1 int
	field2 float64
}

func main() {
	//创建
	//type struct1 struct {
	//	field1 int
	//	field2 float64
	//}
	ms := new(struct1)
	//初始化
	ms := &struct1{10,15.5}
	//当结构体的命名以大写字母开头时，该结构体在包外可见。
	//通常情况下，为每个结构体定义一个构建函数，并且推荐使用构建函数初始化结构体(10.2)
	ms := Newstruct1(10,15.5)
}

func Newstruct1(n int, f float64) *struct1 {
	return &struct1{n, f}
}
