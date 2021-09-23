package main

import (
	"io/ioutil"
	"os"
)

//defer关键字对于函数在结束时关闭打开的文件非常有用。
func data(name string) string {
	f, _ := os.OpenFile(name, os.O_RDONLY, 0)
	defer f.Close()
	contents, _ := ioutil.ReadAll(f)
	return string(contents)
}
//在函数return后执行了f.Close()
