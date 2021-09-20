package main

import "fmt"

var barVal = map[string]int{
	"alpha": 34,
	"bravo": 56,
	"charlie": 23,
	"delta": 87,
	"echo": 56,
	"foxtrot": 12,
	"golf": 34,
	"hotel": 16,
	"indio": 87,
	"juliet": 65,
	"kili": 43,
	"lima": 98,
}

func main() {
	//调换key和value。如果map的值类型可以作为key且所有的value是唯一的，那么通过下面的方法可以简单的做到键值对调。
	invMap := make(map[int]string, len(barVal))
	for k,v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted: ")
	for k, v := range invMap {
		fmt.Printf("key: %v, value: %v /", k, v)
	}
	//如果原始value值不唯一那么这么做肯定会出错；为了保证不出错，当遇到不唯一的key时应当立即停止，这样可能会导致没有包含原map的所有键值对！一种解决方法就是仔细检查唯一性并且使用多值map，比如使用map[int][]string类型。
	//练习8.2 构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；先打印所有的饮料，然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。

}
