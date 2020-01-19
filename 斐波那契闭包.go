package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
    x1, x2 := 0, 1
    sum := 0
    fmt.Println("init:  ","x1:",x1,"x2:",x2,"sum:",sum)
    //以上代码相当于初始化
    return func() int {
        sum = x1 + x2
        x1 = x2
        x2 = sum
	fmt.Println("calc:  ","x1:",x1,"x2:",x2,"sum:",sum)
        return x1
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Print(f(),", ")
    }
    fmt.Println("...")
}
