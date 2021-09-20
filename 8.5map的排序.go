package main

import (
	"fmt"
	"sort"
)

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
	//map默认是无序的，不管是按照key还是按照value默认都是不排序的。
	//如果你想为map排序，需要将key(或者)value拷贝到一个切片，再对切片排序（使用sort包），然后可以使用切片的for-range方法打印出所有的key和value.
	fmt.Println("unsorted: ")
	for k,v := range barVal {
		fmt.Printf("key: %s, value: %d /", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted: ")
	for _, k := range keys {
		fmt.Printf("key: %s, value: %d /", k, barVal[k])
	}
	//如果想要排序的列表最好使用结构体切片，这样会更有效。
	type name struct {
		key string
		value int
	}
}