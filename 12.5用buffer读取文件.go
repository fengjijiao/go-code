package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader, showNN bool) {
	var i int = 0
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		i++
		fmt.Fprintf(os.Stdout, "%d   %s", i, buf)
	}
	return
}

func main() {
	var showNN *bool = flag.Bool("n", false, "show line number")
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin), *showNN)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f), *showNN)
		f.Close()
	}
	/**
	 go run 12.5用buffer读取文件.go -n products.txt
	1   "The ABC of Go";25.5;1500
	2   "Functional Programming with Go";56;280
	3   "Go for It";45.9;356
	*/
}