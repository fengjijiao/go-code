package main

import (
	"fmt"
	"testing"
)

//even_main.go
func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("is the integer %d even? %v\n", i, Even(i))
	}
}
//even/even.go
func Even(i int) bool {
	return i % 2 == 0
}
func Odd(i int) bool {
	return i % 2 != 0
}
//even/even_test.go
func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log("10 must be even!")
		t.Fail()
	}
	if !Even(7) {
		t.Log("7 is not even!")
		t.Fail()
	}
}
func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log("11 must be odd!")
		t.Fail()
	}
	if !Odd(10) {
		t.Log("10 is not odd!")
		t.Fail()
	}
}
//由于测试需要具体的输入用例且不可能测试到所有的用例（非常像一个无穷的数），所以我们必须对要使用的测试用例思考再三。
//至少应该包括
//1.正常的用例
//2.反面的用例（错误的输入，如用负数或字母代替数字，没有输入等）
//3.边界检查用例（如果参数的取值范围是0到1000,检查0和1000的情况）
//