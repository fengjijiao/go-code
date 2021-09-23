package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	Arr []byte
	Top int
}

func (s *Stack) Push(b byte) (err error) {
	if s.Top == cap(s.Arr) {
		err = errors.New("the stack is full")
	} else {
		err = nil
		s.Arr[s.Top] = b
		s.Top++
	}
	return
}

func (s *Stack) Pop() (res byte, err error) {
	if s.Top == 0 {
		err = errors.New("the stack is empty")
	} else {
		err = nil
		res = s.Arr[s.Top]
		s.Top--
	}
	return
}

func (s *Stack) Clear() (err error) {
	if s.Top == 0 {
		err = errors.New("the stack is empty")
	} else {
		err = nil
		s.Top = 0
	}
	return
}

func (s *Stack) Get() (str string, err error) {
	if s.Top == 0 {
		str = ""
		err = errors.New("the stack is empty")
	} else {
		for i := 0; i < s.Top; i++ {
			str += string(s.Arr[i])
		}
		err = nil
	}
	return
}

func main() {
	stack := &Stack{Arr: make([]byte, 6), Top: 0}
	inputReader := bufio.NewReader(os.Stdin)
	buffer := make([]string, 3)
	bufferN := 0
	for {
		bi, err := inputReader.ReadByte()
		if err != nil {
			fmt.Println("input error, " + err.Error())
			continue
		}
		if bi == 'q' {
			fmt.Println("logout")
			return
		}else if bi == '\r' {
		}else if bi != '\n' {
			err = stack.Push(bi)
			if err != nil {
				fmt.Println("stack.Push error, " + err.Error())
				continue
			}
		} else {
			buffer[bufferN], err = stack.Get()
			if err != nil {
				panic(err)
			}
			bufferN++
			err = stack.Clear()
			if err != nil {
				panic(err)
			}
			if bufferN == 3 {
				n1, err := strconv.Atoi(buffer[0])
				if err != nil {
					panic(err)
				}
				n2, err := strconv.Atoi(buffer[1])
				if err != nil {
					panic(err)
				}
				op := buffer[2][0]
				switch op {
				case '+':
					fmt.Printf("result: %d\n", n1+n2)
				case '-':
					fmt.Printf("result: %d\n", n1-n2)
				case '*':
					fmt.Printf("result: %d\n", n1*n2)
				case '/':
					fmt.Printf("result: %d\n", n1/n2)
				default:
					fmt.Println("unknown operator.")
				}
				bufferN = 0
			}
		}
	}
	//for {
	//	number1, err := inputReader.ReadString('\n')
	//	if err != nil {
	//		panic(err)
	//	}
	//	number2, err := inputReader.ReadString('\n')
	//	if err != nil {
	//		panic(err)
	//	}
	//	operator, err := inputReader.ReadString('\n')
	//	if err != nil {
	//		panic(err)
	//	}
	//	n1, err := strconv.Atoi(number1)
	//	if err != nil {
	//		panic(err)
	//	}
	//	n2, err := strconv.Atoi(number2)
	//	if err != nil {
	//		panic(err)
	//	}
	//	op := operator[0]
	//	switch op {
	//	case '+':
	//		fmt.Printf("result: %d\n", n1 + n2)
	//	case '-':
	//		fmt.Printf("result: %d\n", n1 - n2)
	//	case '*':
	//		fmt.Printf("result: %d\n", n1 * n2)
	//	case '/':
	//		fmt.Printf("result: %d\n", n1 / n2)
	//	default:
	//		fmt.Println("unknown operator.")
	//	}
	//}
}
