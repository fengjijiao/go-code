package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "版权所有 (C) Microsoft Corporation。保留所有权利。\n"
	res := strings.Map(func(r rune) rune {
		if r > 255 {//ascii 0~255 or unicode.Is(unicode.Han, r)判断中文
			return '?'
		}
		return r
	}, str)
	fmt.Print(res)
	//str2 := str
	//res2 := []rune(str)
	//j := 0
	//k := 0
	//for {
	//	 i := strings.IndexFunc(str2, func(r rune) bool {
	//		return unicode.Is(unicode.Han, r)
	//	})
	//	if i == -1 {
	//		break
	//	}
	//	k++
	//	j = j + i
	//	fmt.Printf("i: %d, j: %d\n", i, j)
	//	str2 = str2[i+2:]
	//	res2[j] = '?'
	//}
	//fmt.Print(string(res2))
}
