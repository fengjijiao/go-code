package main

import (
    "fmt"
    "math"
    "math/big"
)

func main() {
    // Here are some calculations with bigInts:
    //fmt.Println(math.MaxInt64)//constant 9223372036854775807 overflows int
    im := big.NewInt(math.MaxInt64)
    fmt.Println(im)//9223372036854775807
    in := im
    io := big.NewInt(1956)
    fmt.Println(io)//1956
    ip := big.NewInt(1)
    fmt.Println(ip)//1
    ip.Mul(im, in).Add(ip, im).Div(ip, io)
    fmt.Printf("Big Int: %v\n", ip)
    // Here are some calculations with bigInts:
    rm := big.NewRat(math.MaxInt64, 1956)
    rn := big.NewRat(-1956, math.MaxInt64)
    ro := big.NewRat(19, 56)
    rp := big.NewRat(1111, 2222)
    rq := big.NewRat(1, 1)
    rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
    //c.Add(a,b)计算a+b的值并存储在c中
    fmt.Printf("Big Rat: %v\n", rq)
}

/* Output:
Big Int: 43492122561469640008497075573153004
Big Rat: -37/112
*/