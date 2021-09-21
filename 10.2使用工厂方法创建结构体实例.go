package main

func main() {
	//Go语言不支持面向对象编程语言中的那样的构造子方法，但是可以很容易的在Go中实现“构造子工厂”方法。为了方便通常会为类型定义一个工厂，按照惯例，工厂的名字以new或New开头。假设定义了如下的File结构体类型：
	type File struct {
		fd int//文件描述符
		name string//文件名
	}
	//下面是这个结构体类型对应的工厂方法NewFile，它返回一个指向结构体实例的指针
	//...
	//然后这样调用它
	//f := NewFile(10, "./test.txt")
	//在Go语言中常常使用像上面这样在工厂方法里使用初始化来简便的实现构造函数。
	//
	//如何强制使用工厂方法
	//将结构体类型名命名为小写开头即可
	//
	//10.2.2 map和struct vs new()和make()
	//可以使用make()的三种类型：slices/maps/channels
	//
	//在映射上使用new和make的区别以及可能发生的错误
	//
	//OK
	y := new(File)
	(*y).fd = 2
	(*y).name = "ok"

	//NOT OK，Cannot make File
	//z := make(File)
	//(*z).fd = 2
	//(*z).name = "ok"

	//OK
	type Foo map[string]string
	x := make(Foo)
	x["x"] = "good"
	x["y"] = "ok"

	//NOT OK
	u := new(Foo)
	(*u)["x"] = "good"//运行时出错，panic: assignment to entry in nil map
	(*u)["y"] = "ok"

	//试图make()一个结构体变量，会引发一个编译错误，这还不是太糟糕的，但是new()一个映射并试图使用数据填充它，将会引发运行时错误！因为new(Foo)返回的是一个指向nil的指针，它尚未被分配内存。所以在使用map时要特别谨慎。

}

type File struct {
	fd int//文件描述符
	name string//文件名
}

func NewFile(fd int, filename string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd: fd, name: filename}
}