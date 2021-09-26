package main

import (
	"fmt"
	"strconv"
)

//为了保护一个对象的并发修改，我们可以使用一个后台的协程来顺序处理执行一个匿名函数，而不是通过同步互斥锁(Mutex)进行锁定。
//在下面的程序中，我们有一个Person类型，它包含了一个匿名函数类型的通道字段chF。它在构造器方法NewPerson中初始化，用一个协程启动backend()方法。这个方法在一个无限for循环中执行所有被放到chF上的函数，有效的序列化它们，从而提供安全的并发访问。改变和获取salary可以通过一个放在chF上的匿名函数来实现，backend()会顺序执行它们。注意如何在Salary方法中的闭合（匿名）函数中去包含fChan通道。

//这是一个简化的例子，并且他不应该在这种情况下应用
type Person struct {
	Name string
	salary float64
	chF chan func()
}

func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

//设置salary
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() {
		p.salary = sal
	}
}

//取回salary
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() {
		fChan <- p.salary
	}
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: "+p.Name+"-salary is: "+ strconv.FormatFloat(p.Salary(), 'f', 2,64)
}

func main() {
	bs := NewPerson("smit", 5550.22)
	fmt.Println(bs)
	bs.SetSalary(9000.0)
	fmt.Println("salary changed!")
	fmt.Println(bs)
	/**
	Person - name is: smit-salary is: 5550.22
	salary changed!
	Person - name is: smit-salary is: 9000.00
	*/
}
