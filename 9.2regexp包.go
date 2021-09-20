package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//在下面的程序里，我们将在字符串中对正则表达式进行匹配。
	//如果是简单模式，使用Match方法便可：
	//ok, _ := regexp.Match(pat, []byte(searchIn))
	//变量ok将返回true/false，我们也可以使用MatchString:
	//ok,_ := regexp.MatchString(pat, searchIn)
	//更多的方法中，必须先将正则通过Compile方法返回一个Regexp对象。然后我们将掌握一些匹配，查找，替换相关的功能。
	searchIn := "0hel9lo wor91.93321ld!9"
	pat := "[0-9]+.[0-9]+"
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	fmt.Printf("%s\n", searchIn)
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("match found!")
	}
	//将匹配到的部分替换为"###.#"
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "###.#")
	fmt.Println(str)
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Printf("%s\n", str2)
	//Compile函数也可能返回一个错误，我们在使用时忽略对错误的判断是因为我们确信自己正则表达式是有效的。当用户输入或从数据中获取正则表达式时，我们有必要去校验它的正确性。另外我们也可以使用MustCompile方法，他像Compile方法一样校验正则的有效性，但当正则不合法时程序将panic。
}