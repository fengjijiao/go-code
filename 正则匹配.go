package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var searchIn string = "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"//匹配浮点数
	ok, _ := regexp.Match(pat, []byte(searchIn))
	//或者
	ok2, _ := regexp.MatchString(pat, searchIn)
	fmt.Println(ok,ok2)//true,true
	if aa:= true; aa {
		fmt.Println("ok")//ok
	}
	//与上段类似
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
        	fmt.Println("Match Found!")//Match Found!
    	}
	re, _ := regexp.Compile(pat)//检测你所写的正则表达式是否有问题
	fmt.Println(re)//[0-9]+.[0-9]+
	//将匹配到的部分替换为"##.#"
    	str := re.ReplaceAllString(searchIn, "##.#")
    	fmt.Println(str)//John: ##.# William: ##.# Steve: ##.#
	//参数为函数时
	f := func(s string) string{
        	v, _ := strconv.ParseFloat(s, 32)//str --> float32
		fmt.Printf("%T, %v\n", v, v)//float64, 2578.340087890625 ...
        	return strconv.FormatFloat(v * 2, 'f', 2, 32)//v*2:将匹配到的部分(字符)化为浮点数*2;将浮点数转为字符串
    	}
	str2 := re.ReplaceAllStringFunc(searchIn, f)
    	fmt.Println(str2)//John: 5156.68 William: 9134.46 Steve: 11264.36
	// 将浮点数转为字符串
	// bitSize 表示来源类型（32：float32、64：float64）
	// fmt 表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	// prec 控制精度（排除指数部分）:对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
}