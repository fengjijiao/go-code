package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	//有些时候通过编程的方式去进行计算是不精确的。如果你使用Go语言中的float64类型进行浮点运算，返回结果将精确到15位，足以满足大多数任务。当对超过int64或uint64类型这样的大数进行计算时，如果对精度没有要求，float32或float64可以胜任，但如果对精度有严格要求的时候，我们不能使用浮点数，在内存中它们只能被近似的表示。
	//对于整数的高精度计算Go语言中提供了big包。其中包含了math包：有用来表示大整数的big.Int和表示大有理数big.Rat类型（可以表示为2/5或3.1416这样的分数，而不是无理数或Π）。这些类型可以实现任意位类型的数字，只要内存足够大。缺点是更大的内存和处理开销使它们使用起来要比内置的数字类型慢很多。
	//大的整型数字是通过big.NewInt(n)来构造的，其中n为int64类型整数。而大有理数是通过big.NewRat(N,D)方法构造的。N分子和D分母都是int64型整数。因为Go语言不支持运算符的重载，所以所有大数字类型都有像是Add()和Mul()这样的方法。它们的作用于作为receiver的整数和有理数，大多数情况下它们修改receiver并以receiver作为返回结果。因为没有必要创建big.Int类型的临时变量来存放中间结果，所以这样的运算可通过内存链式存储。
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)//(ip=(im*in)+im)/io,  ip.Add(ip, im) ==> ip = ip + im
	fmt.Printf("big int: %v\n", ip)//43492122561469640008497075573153004
	i0 := big.NewInt(4)
	i1 := big.NewInt(2)
	i2 := big.NewInt(5)
	i1.Mul(i0, i2)//i1=i0*i2=20
	fmt.Printf("i0: %v  i1: %v  i2: %v\n", i0, i1, i2)
	i1.Mul(i0, i2).Add(i0, i2)//i1=i0+i2=9
	fmt.Printf("i0: %v  i1: %v  i2: %v\n", i0, i1, i2)
	i1.Mul(i0, i2).Add(i1, i2)//i1=(i0*i2)+i2=25
	fmt.Printf("i0: %v  i1: %v  i2: %v\n", i0, i1, i2)

	//
	rm := big.NewRat(math.MaxInt64, 1956)//分数a/b
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19,56)
	rp := big.NewRat(1111,2222)
	rq := big.NewRat(1,1)
	rq.Mul(rm,rn).Add(rq, ro).Mul(rq,rp)
	fmt.Printf("big rat: %v\n", rq)
}
