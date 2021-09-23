package main

import (
	"bufio"
	"fmt"
	"os"
)

//下面的程序很好的阐述了io包中的接口概念
func main() {
	//unbuffered
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	//buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	//and now so does buf
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
	//下面是fmt.Fprintf()函数的实际签名
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	//其不是写入一个文件，而是写入一个io.Writer接口类型的变量，下面是Writer接口在io包中的定义：
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
	//fmt.Fprintf()依据指定的格式向第一个参数内写入字符串，第一个参数必须实现了io.Writer接口。Fprintf()能够写入任何类型，只要其实现了Write方法，包括os.Stdout，文件（例如os.File），管道，网络连接，通道等等，通用的也可以使用bufio包中缓冲写入。bufio包中定义了 type Writer struct{}
	//bufio.Writer实现了Write方法：
	//func (b *Writer) Write(p []byte) (nn int, err error)
	//它还有一个工厂函数：传给它一个io.Writer类型的参数，它会返回一个缓冲的bufio.Writer类型的io.Writer:
	//func NewWriter(wr io.Writer) (b *Writer)
	//其适合任何形式的缓冲写入。
	//在缓冲写入的最后千万不要忘了使用Flush()，否则最后的输出不会被写入。
	//在15.2-15.8，我们将使用fmt.Fprintf函数向http.ResponseWriter写入，其同样实现了io.Writer接口。
}
