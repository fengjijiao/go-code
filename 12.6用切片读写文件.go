package main

import (
	"fmt"
	"os"
)

//切片提供了Go中处理IO缓冲的标准方式，下面cat函数的第二版中，在一个切片缓冲内使用无限for循环(直到文件尾部EOF)读取文件，并写入标准输出（os.Stdout）。
func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			panic(err)
		case nr == 0://EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

