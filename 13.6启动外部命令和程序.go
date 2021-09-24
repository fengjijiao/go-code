package main

import (
	"fmt"
	"os"
	"os/exec"
)

//os包有一个StartProcess函数可以调用或启动外部系统命令和二进制可执行文件；它的第一个参数是要运行的进程，第二个参数是用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。
//这个函数返回被启动进程的id(pid)，或者启动失败返回错误。
//exec包中也有同样功能的更简单的结构体和函数；主要是exec.Command(name string, arg ...string)和Run().首先需要用系统命令或可执行文件的名字创建一个Command对象，然后用这个对象作为接收者调用Run()。下面的程序（因为是执行Linux命令，只能在Linux下运行）演示了它们的使用：
func main() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	//1st example: list files
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("the process id is %v", pid)
	//2st example: show all processes
	pid, err = os.StartProcess("/bin/ps", []string{"-e", "-opid,ppid,comm"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("the process id is %v", pid)
	//exec.Run
	cmd := exec.Command("gedit")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("the process id is %v", pid)
}

