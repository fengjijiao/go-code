package main

import (
	"fmt"
)


func main() {
	a :=1.000
	b :=999
	c :=3.87777
	d :=99349
	f :=0.000000000022
	fmt.Printf("%d\n",a)
	fmt.Printf("%d\n",b)
	fmt.Printf("%x or %X\n",d,d)
	fmt.Printf("%e\n",f)
	fmt.Printf("%2d\n",b)//2位
	fmt.Printf("%5.2g\n",c)
}
