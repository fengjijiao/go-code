package main

import (
	"time"
)

func main() {
	//import . "./pack1"
	//当使用.来作为包的别名时，可以不通过包名来使用其中的项目。test :+ ReturnStr()
	//import _ "./pack2"
	//pack2包只导入其副作用，也就是说只执行它的init函数并初始化其中的全局变量。
	//
	//练习9.3 创建一个程序能够和用户说“good day”或“good night”。不同的问候应该放到greetings包中。
	//在同一个包中创建一个ISAM函数返回一个布尔值用来判断当前时间是AM还是PM，同样创建IsAfternoon和IsEvening函数。
	//使用main_greetings作出合适的问候（提示：使用time包）。
	//if IsEvening() {
	//	fmt.Println("good night")
	//}
	//if IsAfternoon() {
	//	fmt.Println("good day")
	//}

	//练习9.4 判断前100个整数是不是偶数，包内同时包含测试的功能。

	//练习9.5 ...
}

func ISAM() bool {
	h := time.Now().Hour()
	if h < 12 {
		return true
	}
	return false
}

func IsAfternoon() bool {
	h := time.Now().Hour()
	if h > 12 && h < 18 {
		return true
	}
	return false
}

func IsEvening() bool {
	h := time.Now().Hour()
	if h > 18 || h < 5 {
		return true
	}
	return false
}
