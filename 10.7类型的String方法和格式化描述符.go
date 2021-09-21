package main

import (
	"fmt"
	"strconv"
)

//当定义了一个有很多方法的类型时，十之八九你会使用String()方法来定制类型的字符串形式的输出，换句话说：一种可阅读性和可打印的输出。如果类型定义了String()方法，它会被用在fmt.Printf()中生成默认的输出：等同意使用格式化描述符%v产生的输出。
//还有fmt.Print()和fmt.Println()也会自动使用String()方法。

type TwoInts struct {
	a int
	b int
}

type T struct {
	a int
	b float32
	c string
}

func (i2 T) String() string {
	return strconv.Itoa(i2.a)+" / "+fmt.Sprintf("%f", i2.b) + " / " + "\""+i2.c+"\""
}

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%f", c)+"℃"
}

var days = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}
type Day int

const(
	M1 = iota
	M2
	M3
	M4
	M5
	M6
	M7
)

func (d Day) String() string {
	return days[d]
}
func main() {
	t := new(TwoInts)
	t.a = 12
	t.b = 9
	fmt.Printf("%v\n", t)
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	fmt.Printf("%#v\n", t)
	/**
	(12/9)
	(12/9)
	*main.TwoInts
	&main.TwoInts{a:12, b:9}
	*/
	//不要在String()方法中调用涉及String()方法的方法，否则会无限迭代调用
	//练习10.12
	t1 := &T{7,-2.35,"abc\tdef"}
	fmt.Printf("%v\n", t1)
	//7 / -2.350000 / "abc    def"

	//练习10.13
	var t2 Celsius = 9.9
	fmt.Println(t2)//9.900000℃

	//练习10.14
	var t3 Day = 2
	fmt.Println(t3)//星期三

	//练习10.15
	//...
}

func (t *TwoInts) String() string {
	return "("+ strconv.Itoa(t.a) + "/" + strconv.Itoa(t.b)+")"
}
