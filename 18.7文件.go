package main

import (
	"fmt"
	"os"
)

//1.如何打开一个文件并读取
f, err := os.Open("filename.txt")
if err != nil {
	panic(err)
}
defer f.Close()
iReader := bufio.NewReader(f)
for {
	str, err := iReader.ReadString('\n')
	if err != nil {//包含EOF
		panic(err)
	}
	fmt.Printf("the input was: %s", str)
}
//2.如何通过切片读写文件
func cat(f *file.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr <0:
			fmt.Fprintf(os.Stdout, "cat: error reading from %s: %s\n", f.String(), er.String())
			os.Exit(1)
		case nr == 0://EOF
			return
		case nr >0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stdout, "cat: error writing from %s: %s\n", f.String(), ew.String())
			}
		}
	}
}
