package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	//syscall
	//重启
	//syscall...
	//container - /list-ring-heap:实现对集合的操作
	//list:双链表
	//ring:环形链表
	//下面的代码演示如何遍历一个链表
	//for e:= l.Front(); e!=nil;e=e.Next(){
	//	//e.Value
	//}
	//练习9.1使用contianer/list包实现一个双向链表，将101、102、103放入其中并打印出来。
	l := list.New()
	l.Init()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)
	for e:=l.Front();e!=nil;e=e.Next() {
		fmt.Printf("%d\n", e.Value)
	}
	//练习9.2通过使用unsafe包中的方法来测试你电脑上一个整型变量占用多少个字节。
	var a int
	fmt.Printf("size is: %d\n", unsafe.Sizeof(a))//8
}
