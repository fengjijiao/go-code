package main

import "fmt"

func main() {
	//1.如何修改字符串中的一个字符
	str := "hello"
	c := []byte(str)
	c[0]='c'
	s2:=string(c)//cello
	//2.如何获取字符串的子串
	subset := str[m:n]
	//3.如何使用for或者for-range遍历一个字符串
	//gives only the bytes
	for i := 0;i<len(str);i++ {
		fmt.Printf("%d: %c\n", i, str[i])
	}
	//gives only the Unicode characters
	for ix, ch := range str {
		//...
	}
	//4.如何获取一个字符串的字节数：len(str)
	//如何获取一个字符串的字符数：最快速utf8.RuneCountInString(str)、len([]int(str))
	//5.如何连接字符串
	//最快速 with a bytes.Buffer(7.2)
	//Strings.Join()(7.4)
	//使用+=
	//6.如何解析命令行参数：使用os或者flag包
}
