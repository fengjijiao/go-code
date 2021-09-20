package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{"ok": 1, "!ok": -2, "err": -1}
	for key, value := range map1 {
		fmt.Printf("%s: %d  ", key, value)
	}
	fmt.Println()
	for _, value := range map1 {
		fmt.Printf("%d  ", value)
	}
	fmt.Println()
	for key := range map1 {
		fmt.Printf("%s   ", key)
	}
	fmt.Println()
	map2 := make(map[int]float32)
	map2[1] = 1.0
	map2[2] = 2.0
	map2[3] = 3.0
	map2[4] = 4.0
	map2[5] = 7.0
	map2[6] = 6.0
	for key, value := range map2 {
		fmt.Printf("%d: %f\n", key, value)
	}
	/**
	3: 3.000000
	4: 4.000000
	5: 7.000000
	6: 6.000000
	1: 1.000000
	2: 2.000000
	*/
	//注意：map不是按照key的顺序排列的，也不是按照value的顺序排列的。
	//问题8.1:下面这段代码的输出是什么？
	capitals := map[string] string {"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Printf("map item: capital of %s is %s \n", key, capitals[key])
	}
	/**
	map item: capital of France is Paris
	map item: capital of Italy is Rome
	map item: capital of Japan is Tokyo
	*/
	//练习 8.1
	//创建一个map来保存每周7天的名字，并将他们打印出来并测试是否存在Tuesday和Hollyday。
}
