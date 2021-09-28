package main

func main() {
	//创建：
	map1 := make(map[keytype]valuetype)
	//初始化：
	map1 := map[string]int{"one":1, "two":2}
	//1.如何使用for或者for-range遍历一个映射
	for key, value := range map1 {
		//...
	}
	//2.如何在一个映射中检测键key1是否存在：
	val1, isPresent = map1[key1]
	//返回值：键key1对应的值或0，true/false
	//3.如何在映射中删除一个键
	delete(map1, key1)
}

