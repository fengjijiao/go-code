package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	//7.6.1从字符串生成字符切片
	//假设s是一个字符串（本质上是一个字节数组），那么就可以直接通过 c := []byte(s)来获取一个字节数组的切片c。另外，你还可以通过copy函数来达到相同的目的：copy(dst []byte, src string)
	//同样的，可以使用for-range来获得每个元素
	s := "hello"
	for i2, i3 := range s {
		fmt.Printf("%d: %c\n", i2, i3)
	}
	//我们知道，Unicode字符会占用两个字节，有些甚至需要3个或4个字节来表示。如果发现错误的UTF8字符，则该字符会被设置为U+FFFD并且索引向前移动一个字节。和字符串转换一样，同样可以使用 c:= []int32(s)语法，这样切片中的每个int都会包含对应的UniCode代码，因为字符串中的每次字符都会对应一个整数。
	//类似的，也可以将字符串转换为元素类型为rune的切片：r := []rune(s)。
	//可以通过代码 len([]int32(s)) 来获取字符串中字符的数量，但使用 utf8.RuneCountInString(s)效率会更高一些。

	//还可以将字符串追加到某一字符数组的尾部：
	var b []byte = []byte{0x00, 0x33, 0x33}
	var s1 string = "hello"
	b = append(b, s1...)
	//7.6.2获取字符串的某一部分
	//使用 substr := str[start:end] 可以从字符串str获取到从索引start开始到end - 1位置的子字符串。同样的，str[start:]则表示获取从start开始到len(str) - 1位置的子字符串。而str[:end]表示获取从0开始到end - 1的子字符串。

	//7.6.3字符串和切片的内存结构
	//在内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数。因为指针对用户来说是完全不可见的，因此我们可以依旧把字符串看做是一个值类型，也就是一个字符数组。
	//...

	//7.6.4修改字符串中的某个字符
	//Go语言中字符串是不可变的，也就是说str[index]这样的表达式是不可以被放在等号左边的。如果尝试运行 str[i] = 'D' 会得到错误 cannot assign to str[i]。
	//因此，必须先将字符串转换成字符数组，然后在通过修改数组中元素的值来达到修改字符串的目的，最后将字符数组转换回字符串格式。
	//例如，将字符串"hello"转换为“cell0”
	s2 := "hello"
	c0 := []byte(s2)
	c0[0] = 'c'
	s3 := string(c0)
	fmt.Printf("res: %s\n", s3) //res: cello
	//所以，可以通过操作切片来完成对字符串的操作。

	//7.6.5字节数组对比函数
	//下面的Compare函数会返回两个字节数组字典顺序的整数对比结果，即0 if a == b, -1 if a < b, 1 if a > b。

	//7.6.6搜索及排序切片和数组
	//标准库提供了sort包来实现常见的搜索和排序操作。可以使用sort包中的函数func Ints(a []int)来实现对int类型的切片排序。例如 sort.Ints(arri)，其中变量arri就是需要被升序排列的数组或切片。为了检查某个数组是否已被排序，可以通过函数IntsAreSorted(a []int) bool来检查，如果返回true则表示已经被排序。
	//类似的，可以使用函数func Float64s(a []float64)来排序float64的元素，或使用func Strings(a []stirng)排序字符串元素。

	//想要在数组或切片中搜索一个元素，该数组或切片必须先被排序（因为标准库的搜索算法使用的是二分法）。然后，就可以使用函数func SearchInts(a []int, n int)int进行搜索，并返回对应结果的索引值。
	//当然还可以搜索float64和string
	//func SearchFloat64s(a []float64, x float64) int
	//func SearchStrings(a []string, x string) int

	//7.6.7 append函数常见操作
	//1.将切片b的元素追加到切片a之后：a = append(a, b...)
	//2.复制切片a的元素到新的切片b上：
	//b = make([]T, len(a))
	//copy(b, a)
	//3.删除位于索引i的元素：
	//a = append(a[:i], a[i+1:]...)
	//4.切除切片a从索引i到j位置的元素：
	//a = append(a[:i], a[j:]...)
	//5.为切片a扩展j个元素长度：
	//a = append(a, make([]T, j))
	//6.在索引i的位置插入元素x：
	// a = append(a[:i], append([]T{x}, a[i:]...)...)
	//7.在索引i的位置插入长度为j的新切片
	//a=append(a[:i], append(make([]T, j), a[i:]...)...)
	//8.在索引为i的位置插入切片b的所有元素
	//a=append(a[:i], append(b, a[i:]...)...)
	//9.取出位于切片a最末尾的元素x
	a := []int{1, 2, 3, 4, 5, 6}
	x, a := a[len(a)-1], a[:len(a)-1]
	fmt.Printf("x: %d, a: %v\n", x, a) //x: 6, a: [1 2 3 4 5]
	//10.将元素x追加到切片a：
	//a=append(a, x)
	//因此。可以使用切片和append操作来表示任意可变长度的序列。
	//从数学的角度来看，切片相当于向量，如果需要的话可以定义一个向量作为切片的别名来进行操作。
	//如果需要更完整的方案，可以学习一下 Eleanor McHugh编写的几个包：slice、chain和lists。

	//7.6.8切片和垃圾回收
	//切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。
	//示例 函数FindDigits 将一个文件加载到内存，然后搜索其中所有的数字并返回一个切片。
	//...
	//这段代码可以顺利运行，但返回的[]byte指向的底层是整个文件的数据。只要该返回的切片不被释放，垃圾回收器就不能释放整个文件所占用的内存。换句话说，一点点有用的数据却占用了整个文件的内存。
	//想要避免这个问题，可以通过拷贝我们需要的部分到一个新的切片中：
	//FindDigits2
	//事实上，上面这个函数只能找到第一个匹配正则表达式的字符串。想要找到所有数字，可以这样做：
	//FindDigits3

	//练习7.12 编写一个函数，要求其接受两个参数，原始字符串str和分割索引i，然后返回两个分割后的字符串。
	s4 := "hello world!"
	s5, s6 := SplitSlice(s4, 6)
	fmt.Printf("%s@%s\n", s5, s6) //hello @world!
	//练习7.13 假设有字符串str，那么str[len(str)/2:] + str[:len(str)/2]的结果是什么？
	s7 := s4[len(s4)/2:] + s4[:len(s4)/2] //这样两个s4[...]并列执行
	fmt.Printf("%s\n", s7)                //world!hello
	//练习7.14 编写一个程序，要求能够反转字符串，即将“Google”转换成“elgooG”（提示：使用[]byte类型的切片）
	fmt.Printf("reversed: %s\n", ReverseString("Google"))
	//如果使用两个切片来实现翻转，请再尝试使用一个切片（提示：使用交换法）。
	//如果想要反转Unicode编码的字符串，请使用[]int32类型的切片。
	//练习7.15 编写一个程序，要求能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组。
	s8 := "oookkkaaaabbbbdddeee"
	fmt.Printf("source: %s\nresult: %s\n", s8, string(ArrDiff(s8)))
	/**
	source: oookkkaaaabbbbdddeee
	result: okabde
	*/
	//编写一个程序，使用冒泡排序的方法排序一个包含整数的切片。
	a3 := []int{8, 5, 3, 9, 10, 1, 6}
	fmt.Printf("%v\n", BubbleSortArr(a3))
	//在函数式编程中，一个map-function是指能够接受一个函数原型和一个列表，并使用列表中的值依此执行函数原型，公式为：map(F(), (e1,e2,...,en))=(F(e1),F(e2)...F(en))
	//编写一个函数mapFunc要求接受以下2个参数：
	//1.一个将整数乘以10的函数
	//2.一个整数列表
	//最后返回保存结果的整数列表。
	fmt.Printf("%v\n", MapFunc(func(i2 int) int {
		return i2 * 10
	}, []int{1, 2, 3, 4, 5, 6}))
}

func MapFunc(f func(int) int, l0 []int) (res []int) {
	for _, i3 := range l0 {
		res = append(res, f(i3))
	}
	return
}

func BubbleSortArr(slice []int) []int {
	l := len(slice)
	var temp int
	for range slice {
		for i4 := range slice {
			if i4+1 < l && slice[i4+1] < slice[i4] {
				temp = slice[i4]
				slice[i4] = slice[i4+1]
				slice[i4+1] = temp
			}
		}
	}
	return slice
}

func ArrDiff(s0 string) (res []byte) {
	a0 := []byte(s0)
	//l0 := len(a0)-1
	j := 0
	res = append(res, a0[0])
	for i := 0; (i + 1) < len(a0); i++ {
		if a0[j] != a0[i+1] {
			res = append(res, a0[i+1])
			j = i + 1
		}
	}
	return
}

func ReverseString(str string) (res string) {
	b := []byte(str)
	l := len(b) - 1
	for i2 := range b {
		fmt.Printf("%d---%d\n", i2, l/2)
		temp := b[l-i2]
		b[l-i2] = b[i2]
		b[i2] = temp
		if i2 == l/2 {
			break
		}
	}
	res = string(b)
	return
}

func SplitSlice(str string, i2 int) (s0, s1 string) {
	b := []byte(str)
	s0 = string(b[:i2])
	s1 = string(b[i2:])
	return
}

func FindDigits3(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes)) //!!important
	c := make([]byte, 0)
	for _, bytes := range b { //!!important
		c = append(c, bytes...)
	}
	return c
}

func FindDigits2(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}
