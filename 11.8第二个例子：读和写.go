package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

//只要类型实现了读写接口，提供Read()和Write()方法，就可以从它读取数据，或向它写入数据。一个对象要是可读的，他必须实现io.Reader接口，这个接口只有一个前名是Read(p []byte) (n int, err error)的方法，它从调用它的对象上读取数据，并把读取到的数据放入参数中的字节切片中，然后返回读取的字节数和一个error对象，如果没有错误发生返回nil，如果已经达到输入的尾端，会返回io.EOF("EOF")，如果读取的过程中发生了错误，就会返回具体的错误信息。
//类似的，一个对象要是可写的，它必须实现io.Writer接口，这个接口也只有一个签名是Write(p []byte) (n int, err error)的方法，它将指定字节切片中的数据写入调用它的对象里，然后返回实际写入的字节数和一个error对象（如果没有错误发生就是nil）。
//io包里的Readers和Writers都是不带缓冲区的，bufio包里提供了对应的带缓冲的操作，在读写UTF-8编码的文本文件时它们尤其有用。详见12
//在实际的编程中尽可能的使用这些接口，会使程序变得更通用，可以在任何实现了这些接口的类型上使用读写方法。
//例如：一个JPEG图形解码器，通过一个Reader参数，它可以解码来自磁盘、网络连接或以gzip压缩的HTTP流中的JPEG图像数据，或其他任何实现了Reader接口的对象。

type TestO struct {
	data []byte
	io.Reader
	io.Writer
}

func (t *TestO) Read(p []byte) (n int, err error) {
	if len(t.data) >= 3 {
		p[0] = t.data[0]
		p[1] = t.data[1]
		p[2] = t.data[2]
		return 3, nil
	}
	return 0, errors.New("data length < 3")
}

func (t *TestO) Write(p []byte) (n int, err error) {
	if len(p) >= 3 {
		t.data[0] = p[0]
		t.data[1] = p[1]
		t.data[2] = p[2]
		return 3, nil
	}
	return 0, errors.New("data length < 3")
}

func main() {
	var t TestO
	t.data = make([]byte, 3)
	data := []byte {0x01, 0x55, 0xff}
	fmt.Println("data: ", data)
	if _, err := t.Write(data); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	res := make([]byte, 3)
	if _, err := t.Read(res); err == nil {
		fmt.Println("res: ", res)
	}
	/**
	data:  [1 85 255]
	res:  [1 85 255]
	*/
}
