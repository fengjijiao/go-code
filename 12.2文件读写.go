package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//12.2.1读文件
//在Go语言中，文件使用指向os.File类型的指针来标识的，也叫文件句柄。标准输入os.Stdin和标准输出os.Stdout，它们的类型都是*os.File。
func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString ,readerError := inputReader.ReadString('\n')
		fmt.Printf("the input was %s\n", inputString)
		if readerError == io.EOF {
			return
		}
	}
	//inputFile是*os.File类型的。该类型是一个结构，表示一个打开文件的描述符（文件句柄）。然后，使用os包里的Open函数来打开一个文件。该函数的参数是文件名，类型为string。在上面的程序中，我们以只读模式打开input.txt文件。
	//如果文件不存在或者程序没有足够的权限打开这个文件，Open函数会返回一个错误。如果文件打开正常，我们就使用defer inputFile.Close()语句确保在程序退出前关闭该文件。然后，我们使用bufio.NewReader来获取一个读取器变量。
	//通过使用bufio包提供的读取器（写入器也类似），如上面程序所示，我们可以很方便的操作相对高层的string对象，而避免了去操作比较底层的字节。
	//接着，我们在一个无限循环中使用ReadString('\n')将或ReadBytes('\n')将文件的内容逐行（行结束符'\n'）读取出来。
	//注意：在使用ReadString和ReadBytes方法的时候，我们不需要关心操作系统的类型，直接使用\n就可以了。另外我们也可以使用ReadLine()方法来实现相同的功能。
	//一旦读取到文件末尾，变量readerError的值将变成非空(事实上，常量io.EOF的值为true)，我们就会执行return语句从而退出。
	//1) 将整个文件的内容读到一个字符串里：
	//可以使用io/ioutil包里的ioutil.ReadFile()方法，该方法第一个返回值的类型是[]byte，里面存放读取到的内容，第二个返回值是err。类似的，函数WriteFile()方法可以将[]byte的值写入文件。
	inputFileName := "p0.txt"
	outputFileName := "p0_out.txt"
	buf, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFileName, buf, 0644)
	if err != nil {
		panic(err)
	}
	//2)带缓冲的读取
	//在很多情况下，文件的内容是不按行划分的，或者干脆就是一个二进制文件。在这种情况下，ReadString()就无法使用了，我们可以使用bufio.Reader的Read()方法，它只接收一个参数：
	//buf = make([]byte, 1024)
	////...
	//n ,err := inputReader.Read(buf)
	//if n == 0 {
	//	break
	//}
	//n表示读取到的字节数
	//3)按列读取文件中的数据
	//如果数据是按列排列并用空格分割的，你可以使用fmt包提供的以FScan开头的一系列函数来读取它们。
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1,v2,v3 string
		_,err := fmt.Fscanln(file, &v1,&v2,&v3)
		//scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
	/**
	[ABC FUNC GO]
	[46 56 45]
	[150 280 356]
	 */
	//注意：path包里有一个子包filepath，这个子包提供了跨平台的函数，用于处理文件名和路径。例如Base()函数用于获取路径中的最后一个元素（不包含后面的分隔符）：
	//import "path/filepath"
	//filename := filepath.Base(path)



	//12.2.2 compress包：读取压缩文件。
	//compress包提供了读取压缩文件的功能，支持的压缩文件格式为：bzip2\flate\gzip\lzw\zlib。
	//下面的程序展示了如何读取一个gzip文件。
	fName := "myFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	}else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}

	//12.2.3写文件
	outputFile, outputError := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		panic(outputError)
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
	//除了文件句柄，我们还需要bufio的Writer。我们以只写模式打开文件output.txt,如果文件不存在则会自动创建。
	//os.O_RDONLY:只读
	//os.O_WRONLY:只写
	//os.O_CREATE:创建，如果指定文件不存在，就创建文件。
	//os.O_TRUNC:截断：如果指定文件存在，就将该文件的长度截为0
	//在读文件时，文件的权限是被忽略的，所以使用OpenFile时传入的第三个参数可以用0.而在写文件时，不管是Unix还是Windows都需要使用0666。
	//然后，我们创建一个写入器（缓冲区）对象：
	//outputWriter := bufio.NewWriter(outputFile)
	//接着，使用一个for循环，将字符串写入缓冲区，写10次。
	//写入文件Flush()
	//如果写入的东西简单，我们可以使用fmt.Fprintf(outputFile, "some test data\n")直接将内容写入文件。fmt包里的F开头的Print函数可直接写入任何io.Writer，包括文件。
	//os.Stdout.WriteString("hello world!")//输出到屏幕
	//outputFile, outputError := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	//outputFile.WriteString("hello world!")
	//不使用缓冲区直接将内容写入文件。
}

