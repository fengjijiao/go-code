package main

import (
	"bufio"
	"fmt"
	"os"
)

//我们如何读取用户的键盘（控制台）输入呢？从键盘和标准输入os.Stdin读取输入，最简单的办法是使用fmt包提供的Scan和Sscan开头的函数。

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.2 / 5212 / Go"
	format = "%f / %d / %s"
)

var (
	inputReader *bufio.Reader
	input string
	err error
)

func main() {
	fmt.Println("please enter your full name: ")
	_, err := fmt.Scanln(&firstName, &lastName)
	if err != nil {
		panic(err)
		return
	}
	//fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	_, err = fmt.Sscanf(input, format, &f, &i, &s)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("from the string we read: ", f, i, s)
	/**
	please enter your full name:
	a a
	Hi a a!
	from the string we read:  56.2 5212 Go
	*/
	//Scanln扫描来自标准输入的文本，将空格分隔的值依此存放到后续的参数内，直到碰到换行。
	//Scanf与其类似，除了Scanf的第一个参数用作格式字符串，用来决定如何读取。
	//Sscan和以Sscan开头的函数则是从字符串读取，除此之外，与Scanf相同。如果这些函数读取到的结果与预想的不同，你可以检查成功读入数据的个数和返回的错误。
	//也可以使用bufio包提供的缓冲读取（buffered reader）来读取数据，例如：
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("the input was: %s\n", input)
	}
	//inputReader是一个指向bufio.Reader的指针。
	//inputReader := bufio.NewReader(os.Stdin)这行代码，将会创建一个读取器，并将其与标准输入绑定。
	//bufio.NewReader()构造函数的签名为：func NewReader(rd io.Reader) *Reader
	//该函数的实参可以是满足io.Reader接口的任意对象，函数返回一个带缓冲的io.Reader对象，它将从指定读取器（例如：os.Stdin）读取内容。
	//返回的读取器对象提供一个方法ReadString(delim byte)，该方法从输入中读取内容，直到碰到delim指定的字符，然后将读取到的内容连同delim指定的字符一起放到缓冲区。
	//ReadString返回读取到的字符串，如果碰到错误则返回nil。如果它一直读到文件结束，则返回读取到的字符串和io.EOF。如果读取过程中没有碰到delim字符，将返回错误err != nil。
	//在上面的例子中，我们会读取键盘输入，直到回车键(\n)被摁下。
	//屏幕是标准输出os.Stdout；os.Stderr用于显示错误信息，大多数情况下等同于os.Stdout。
	//一般情况下，我们会省略变量声明，而使用:=，例如：
	//inputReader := bufio.NewReader(os.Stdin)
	//input, err := inputReader.ReadString('\n')
	//第二个例子，从键盘读取输入，使用了switch语句：
	//
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your name: ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Printf("your name is %s\n", input)
	//Unix delimiter \n, Windows: \r\n
	switch input {
	case "Philip\r\n":
		fmt.Println("welcome philip!")
	case "chris\r\n":
		fmt.Println("welcome chris!")
	case "ivo\r\n":
		fmt.Println("welcome ivo!")
	default:
		fmt.Println("you are not welcome here! goodbye!")
	}
	//Version 2
	switch input {
	case "Philip\r\n": fallthrough
	case "chris\r\n": fallthrough
	case "ivo\r\n":
		fmt.Printf("welcome %s!", input)
	default:
		fmt.Println("you are not welcome here! goodbye!")
	}
	//Version 3
	switch input {
	case "Philip\r\n", "chris\r\n", "ivo\r\n":
		fmt.Printf("welcome %s!", input)
	default:
		fmt.Println("you are not welcome here! goodbye!")
	}
	//注意：Unix和Windows的行结束符是不同的!

}
