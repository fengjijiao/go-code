package main

//名为testing的包被专门用来进行自动化测试，日志和错误报告。并且还包含一些基准测试函数的功能。
//对一个包做（单元）测试，需要写一些可以频繁（每次更新后）执行的小块测试单元来检查代码的正确性。于是我们必须写一些Go源文件来测试代码。测试程序必须属于被测试的包，并且文件名满足这种形式*_test.go，所以测试代码和包中的业务代码分开的。
//_test程序不会被普通的Go编译器编译，所以当放应该部署到生产环境时它们不会被部署；只有gotest会编译所有的程序：普通程序和测试程序。
//测试文件中必须导入“testing”包，并写一些名字以TestZzz打头的全局函数，这里的Zzz是被测试函数的字母描述，如TestPay等.
//测试函数必须有这种形式的头部：
//func TestAbcde(t *testing.T)
//T是传给测试函数的结构类型，用来管理测试状态，支持格式化测试日志，如t.Log，t.Error,t.Errorf等。在函数的结尾把输出跟想要的结果对比，如果不等就打印一个错误。成功的测试则直接返回。
//用下面这些函数来通知测试失败：
//1) func (t *T) Fail()
//标记测试函数为失败，然后继续执行（剩下的测试）。
//2) func (t *T) FailNow()
//标记测试函数为失败并中止执行；文件中别的测试也被忽略，继续执行下一个文件。
//3) func (t *T) Log(args ...interface{})
//args被用默认的格式格式化并打印到错误日志中
//4) func (t *T) Fatal(args ...interface{})
//结合 先执行3）后执行2)的效果。
//运行go test来编译测试程序，并执行程序中所有的TestZZZ函数。如果所有的测试都通过会打印出PASS。
//go test可以接收一个或多个函数程序作为参数，并指定一些选项。
//结合--chatty或-v选择，每个执行的测试函数以及测试状态会被打印。
//例如:
//go test fmt_test.go --chatty
//===RUN fmt.TestFlagParser
//---PASS: fmt.TestFlagParser
//...
//...
//testing包中有一些类型和函数可以用来做简单的基准测试；测试代码中必须包含以BenchmarkZzz打头的函数并接收一个*testing.B类型的参数，比如：
//func BenchmarkReverse(b *testing.B) {
//...
//}
//命令go test -test.bench=.* 会运行所有的基准测试函数；代码中的函数会被调用N次（N是非常大的数，如N=1000000），并展示N的值和函数执行的平均时间，单位为ns。如果是用testing.Benchmark调用这些函数，直接运行程序即可。
//具体例子见14.16中goroutines运行基准测试。