package main

import "fmt"

func main() {
	//8.1.1
	//map是引用类型，可以使用如下声明：
	//var map1 map[keytype]valuetype
	//var map1 map[string]int
	//[keytype]和valuetype之间允许有空格，但是gofmt移除了空格
	//在声明的时候不需要知道map的长度，map是可以动态增长的。
	//未初始化的map的值是nil。
	//key可以是任意可以用==或！=操作符比较的类型，比如string、int、float。所以切片和结构体不能作为key(含有数组切片的结构体不能作为key，只包含内建类型的struct是可以作为key的)，但是指针和接口类型可以。
	//如果要用结构体作为key可以提供Key()和Hash()方法，这样可以通过结构体的域计算出唯一的数字或者字符串的key。
	//value可以是任意类型的；通过使用空接口类型，我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言(.(type))。
	//map传递给函数的代价很小；在32位机器上占4个字节，64位机器上占8个字节，无论实际上存储了多少数据。通过key在map中寻找值是很快的，比线性查找快得多，但是仍然比从数组或切片的索引中直接读取要慢100倍；所以如果你很在乎性能的话还是建议用切片来解决问题。
	//map也可以用函数作为自己的值，这样就可以用来做分支结构：key用来选择要执行的函数。
	//如果key1是map1的key，那么map1[key1]就是对应key1的值，就如同数组的索引符号一样（数组可以视为一种简单形式的map，key是从0开始的整数）。
	//key1对应的值可以通过赋值符号来设置为val1：map1[key1]=val1。
	//令 v:=map1[key1]可以将key1对应的值赋值给v；如果map中没有key1存在，那么v将被赋值为map1的值类型的空值。
	//常用的len(map1)方法获得map中的pair数目，这个数目是可以伸缩的，因为map-pairs在运行时可以动态的添加和删除。
	//map可以使用{key1:val1,key2:val2}的描述方法来初始化，就像数组和结构体一样。
	//map是引用类型的；内存用make方法来分配。
	//map的初始化：var map1 = make(map[keytype]valuetype)
	//make(map[string]float32) = map[string]float32{}
	//map2 = map1，这里map2也是map1的引用，对map2的修改也会影响到map1的值。
	//
	//
	//不要使用new，永远使用make来构造map
	//注意：如果你错误的使用new()分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	//map3 := new(map[string]float32)
	//接下来当我们调用：map3["key1"] = 4.5时，编译器会报错。
	//invalid operation: map3["key1"](index of type *map[string]float32)。
	//为了说明值可以是任意类型的，这里给出了一个使用func()int作为值的map。
	mf := map[int]func() int {
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Print(mf)
	//整形都被映射到了函数地址。

	//8.1.2 map容量
	//和数组不同，map可以根据新增的key-value对动态的伸缩，因此它不存在固定长度或者最大限制。但是你也可以选择标明map的初始容量capacity，就像这样：make(map[keytype]valuetype, cap)。
	//例如map2 := make(map[string]float32, 100)
	//当map增长到容量上限时，如果再增加新的key-value对，map的大小会自动加1.所以出于性能考虑，对于大的map或者会快速扩张的map，即使只是大概知道容量，也最好先标明。

	//8.1.3 用切片作为map的值
	//既然一个key只能对应一个value，而value又是一个原始类型，那么如果一个key要对应多个值该怎么办？例如，当我们要处理unix机器上的所有进程，以父进程(pid为整数)作为key，所有子进程（以所有子进程的pid组成的切片）作为value。通过将value定义为[]int类型或者其他类型的切片，就可以优雅的解决这个问题。
	//mp1 := make(map[int][]int)
	//mp2 := make(map[int]*[]int)
}
