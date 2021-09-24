package main

import "testing"

//编写测试代码时，一个较好的办法是把测试代码的输入数据和期望的结果写在一起组成一个数据表：表中的每条记录都是一个含有输入和期望值的完整测试用例，有时还可以结合像测试名字这样额外信息来让测试输出更多的信息。
//实际测试时简单迭代表中的每条记录，并执行必要的测试。
//可以抽象为下面这段代码
var tests = []struct{//test table
	in string
	out string
} {
	{"in1", "exp1"},
	{"in2", "exp2"},
	{"in3", "exp3"},
}
func TestFunction(t *testing.T) {
	for i, tt := range tests {
		s := FuncToBeTested(tt.in)
		if s != tt.out {
			t.Errorf("%d. %q => %q\n,wanted: %q", i,tt.in,s,tt.out)
		}
	}
}
//如果大部分函数都可以写成这种形式，那么写一个帮助函数verify对实际测试会很有帮助：
func verify(t *testing.T, testnum int, testcase, input, output, expected string) {
	if expected != output {
		t.Errorf("%d. %s with input = %s:\noutput %s != %s", testnum, testcase, input, output, expected)
	}
}
//TestFunction则变为
func TestFunction(t *testing.T) {
	for i, tt := range tests {
		s := FuncToBeTested(tt.in)
		verify(t, i, "FuncToBeTested: ", tt.in, s, tt.out)
	}
}