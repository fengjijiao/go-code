package main

import (
	"flag"
	"os"
	"runtime/pprof"
)

//13.10.1时间和内存消耗
//可以用这个便捷脚本xtime来测量
//#!/bin/sh
///usr/bin/time -f '%Uu %Ss %er %MkB %C' "$@"
//在Unix命令行中像这样使用xtime goprogexec，这里的progexec是一个go可执行程序，这句命令行输出类型： 56.63u 0.26s 56.92r 1642640kB progexec
//分别对应用户时间，系统时间，实际时间和最大内存占用。

//13.10.2用go test调试
//如果代码使用了Go中testing包的基准测试功能，我们可以用gotest标准的-cpuprofile和-memprofile标志向指定文件写入CPU或内存使用情况报告。
//使用方式：go test -x -v -cpuprofile=prof.out -file x_test.go
//编译执行x_test.go中的测试，并向prof.out文件中写入CPU性能分析信息。

//13.10.3用pprof调试
//你可以在单机程序progexec中引入runtime/pprof包；这个包以pprof可视化工具需要的格式写入运行时报告数据。对于CPU性能分析来说你需要添加一些代码：
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	//...
}
//代码定义了一个名为cpuprofile的flag，调用Go flag库来解析命令行flag,如果命令行设置了cpuprofile flag，则开始CPU性能分析并把结果重定向到那个文件。(os.Create 用拿到的名字创建了用来写入分析数据的文件)。这个分析程序最后需要在程序退出之前调用StopCPUProfile来刷新挂起的写操作到文件中；我们用defer来保证这一切会在main返回时触发。
//现在用这个flag运行程序：progexec -cpuprofile=progexec.prof
//然后可以像这样用gopprof工具：gopprof progexec progexec.prof
//gopprof程序是Google pprof C++分析器的一种轻微变种；更多信息：github.com/gperftools/gperftools
//如果开启了CPU性能分析，Go程序会以大约每秒100次的频率阻塞，并记录当前执行的goroutine栈上的程序计数器样本。
//此工具一些有趣的命令：
//1)topN
//用来展示分析结果中最开头的N份样本，例如top5它会展示在程序运行期间调用最频繁的5个函数，输出如下
//total: 3099 samples
//626 20.2% 20.2% 626 20.2% scanblock
//...
//第5列表示函数调用的频度。
//2) web或web函数名
//该命令生成一份SVG格式的分析数据图标，并在网络浏览器中打开它（还有一个gv命令可以生成PostScript格式的数据，并在GhostView中打开，这个命令需要安装graphviz），函数被表示成不同的矩形（被调用越多，矩形越大），箭头指示函数调用链。
//3) list 函数名 或 weblist 函数名
//展示对应函数名的代码行列表，第2行表示当前行执行消耗的时间，这样就很好地指出了运行过程中消耗最大的代码。
//如果发现函数runtime.mallocgc（分配内存并执行周期性的垃圾回收）调用频繁，那么是应该进行内存分析的时候了。找出垃圾回收频繁执行的原因，和内存大量分配的根源。
//为了做到这一点必须在合适的地方添加下面的代码
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
//...
func CallToFunctionWhichAllocatesLotsOfMemory() {
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			panic(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
//用-memprofile flag运行这个程序：progexec -memprofile=progexec.mprof
//然后你可以像这样再次使用gopprof工具：gopprof progexec progexec.mprof
//top5， list 函数名 等命令同样适用，只不过现在是以Mb为单位测量内存分配情况，这时top命令输出的例子：
//Total: 118.3MB
//66.1 55.8% 55.8% 103.7 87.7% main.FindLoops
//...
//从第1列可以看出，最上面的函数占用了最多的内存。
//同样有一个报告内存分配计数的有趣工具：
//gopprof --inuse_objects progexec progexec.mprof
//对于web应用来说，有标准的HTTP接口可以分析数据。
//在HTTP服务中添加
//import _ "http/pprof"
//会为/debug/pprof 下的一些URL安装处理器。然后你可以用一个唯一的参数 --- 你服务中的分析数据的URL来执行gopprof命令---它会下载并执行在线分析。
//gopprof http://127.0.0.1:6060/debug/pprof/profile
//#30 s CPU Profile
//gopprof http://127.0.0.1:6060/debug/pprof/heap
//heap profile