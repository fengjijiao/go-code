package main

import (
	"fmt"
)

func main() {
	//移位运算
	//
	//按位与&
	//按位或|
	//按位异或^
	//位清除&^:将指定位置上的值设置为0
	//
	//一元运算
	//按位补足^
	//位左移n位<<
	//bitP<<n
	type ByteSize float64
	const (
		_ = iota
		KB ByteSize = 1<<(10*iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	//位右移n位>>
	//bitP>>n
}
