package main

import (
	"fmt"
)

func main() {
	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v\n", c1)
	var re,im float32 = 3,12
	var c2 = complex(re,im)
	fmt.Printf("float32 -> complex32:%v\n",c2)
	fmt.Printf("real part: %f & image part: %f\n",real(c2),imag(c2))
}
