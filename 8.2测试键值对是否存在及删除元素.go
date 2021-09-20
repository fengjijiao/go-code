package main

import "fmt"

func main() {
	map1 := map[string]float32{"ok": 3.3}
	val, ok := map1["ok"]
	fmt.Printf("value: %f, exist: %t\n", val, ok)
	val2, ok2 := map1["ok2"]
	fmt.Printf("value2: %f, exist2: %t\n", val2, ok2)
	if _, ok := map1["ok"]; ok {
		fmt.Printf("value: %f\n", map1["ok"])
	}
	//从map1中删除key1:
	//直接delete(map1, key1)就可以。
	//如果key1不存在，该操作不会产生错误。
}
