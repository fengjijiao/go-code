package main


//如果你在一个for循环内部处理一系列文件，你需要使用defer确保文件在处理完毕后被关闭。
//for _, file := range files {
//	if f, err = os.Open(file); err != nil {
//		return
//	}
//	//这是错误的方式，当循环结束时文件没有关闭
//	defer f.Close()
//	//对文件进行操作
//	f.Process(data)
//}
//但在循环结尾处的defer没有执行，所以文件一直没有关闭！垃圾回收机制可能会自动关闭文件，但是这会产生一个错误，更好的做法是：
//for _, file := range files {
//	if f, err = os.Open(file); err != nil {
//		return
//	}
//	//对文件进行操作
//	f.Process(data)
//  //关闭文件
//  f.Close()
//}


//defer仅在函数返回时才会执行，在循环的结尾或其他一些有限范围的代码内不会执行。