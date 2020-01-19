package main
import (
	"fmt"
	"flag"
	"math"
)
var a = flag.Int("a",1,"被加数a")
var b = flag.Int("b",2,"加数b")
var c float64
func init() {
	c = 666 * math.Atan(1)
}
func main() {
	fmt.Println("我是从init初始化的c:",c)
	flag.Parse()
	fmt.Println("被加数：", *a)
	fmt.Println("加数：", *b)
	var result = add(*a, *b)
	fmt.Println("结果：", result)
}
func add(a int,b int) int {
	return a+b
}
