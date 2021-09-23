package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('S')
	if err != nil {
		panic(err)
	}
	fmt.Printf("total length: %d\n", CalcTotalLen(input))
	fmt.Printf("word numbers: %d\n", CalcWordN(input))
	fmt.Printf("line numbers: %d\n", CalcLineN(input))
	/**
	w a a w c
	a d h t
	d r
	S
	total length: 20
	word numbers: 12
	line numbers: 4
	*/
}

func CalcTotalLen(str string) (res int) {
	for _, s := range str {
		if s != '\r' && s != '\n' {
			res++
		}
	}
	return
}

func CalcWordN(str string) (res int) {
	for _, s := range str {
		if s == ' ' {
			res++
		}
	}
	res+=CalcLineN(str)
	return
}

func CalcLineN(str string) (res int) {
	for _, s := range str {
		if s == '\n' {
			res++
		}
	}
	res++
	return
}
