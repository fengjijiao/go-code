package main

import "log"

//当资源不再被使用时，使用defer延迟调用其后的代码，确保资源能够被关闭或返回给连接池。其次最重要的是从panic中恢复程序运行。
//1.关闭文件流（12.7）
//open a file f
defer f.Close()
//2.解锁一个已加锁资源（a mutex）（9.3）
mu.Lock()
defer mu.UnLock()
//3.关闭channel（如果必要的话）
ch := make(chan float64)
defer close(ch)
//4.从panic中恢复(13.3)
defer func() {
	if err := recover(); err != nil {
		log.Printf("run time panic: %v\n", err)
	}
}
//5.停止一个ticker(14.5)
tick1 := time.NewTicker(updateInterval)
defer tick1.Stop()
//6.释放一个进程p(13.6)
p, err := os.StartProcess(...,...,...)
defer p.Release()
//7.停止CPU分析并刷新信息(13.10)
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
//它也能用于不要忘了一个报告中打印页脚