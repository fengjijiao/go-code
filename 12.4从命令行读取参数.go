package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//12.4.1 os包
//os包中有一个string类型的切片变量os.Args，用来处理一些基本的命令行参数，它在程序启动后读取命令行输入的参数。
func main() {
	who := "alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("good morning, %s\n", who)
	/**
	go run 12.4从命令行读取参数.go ok
	good morning, alice ok
	*/
}

//12.4.2 flag包
//flag包有一个扩展功能用来解析这些命令行选项。但是通常被用来替换基本常量。例如，在某些情况下我们希望在命令行给常量一些不一样的值（参见19）。
//在flag包中有一个Flag被定义成一个含有如下字段的结构体：
//type Flag struct {
//	Name string//name as it appears on command line
//	Usage string//help message
//	Value Value//value as set
//	DefValue string//default value (as text);for usage message
//}
//模拟Unix的echo功能
var NewLine = flag.Bool("n", false, "print newline")

const (
	Spance = ""
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()//命令行提示信息
	flag.Parse()//scan the arg list and sets up flags
	var s string = ""
	for i:=0;i<flag.NArg();i++ {
		if i>0 {
			s += " "
			if *NewLine {
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
//flag.Parse()扫描参数列表（或常量列表）并设置flag，flag.Arg(i)表示第i个参数。Parse()之后flag.Arg(i)全部可用，flag.Arg(0)就是第一个真实的flag，而不像os.Args(0)放置程序的名字。
//flag.Narg()返回参数的数量。解析后flag或常量就可以用了。
//flag.Bool()定义了一个默认值是false的flag：当在命令行出现的第一个参数（这里是n），flag被设置成true(NewLine是*bool类型)。flag被解引用到*NewLine，所以当值是true时将添加一个newline("\n")。
//flag.PrintDefaults()打印flag的使用帮助信息，本例中打印的是
//-n=false; print newline
//flag.VisitAll(fn func(*Flag))是另一个有用的功能：按照字典顺序遍历flag，并且对每个标签调用fn(参考15.8)。
