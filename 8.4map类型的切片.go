package main

import "fmt"

func main() {
	//假设我们想获取一个map类型的切片，我们必须使用两次make()函数，第一次分配切片，第二次分配切片中的每个map元素。
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("value of items: %v\n", items)
	//不能使用以下方式初始化
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)//the item only is a copy of slice element.
		item[1] = 2//the value will be lost on the next iteration.
	}
	fmt.Printf("items2: %v\n", items2)
	//需要注意的是，应当像A那样通过索引使用切片的map元素。在B版本中获得的项只是map值的一个拷贝而已，真正的map元素没有得到初始化。
}