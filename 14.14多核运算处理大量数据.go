package main

//假设我们必须处理大量的彼此独立的数据项，通过一个输入通道进入，并且全部处理完成后放到一个输出通道，就像一个工厂的管道。每个数据项的处理也许会涉及多个步骤：
//预处理/步骤A/步骤B/.../后期处理
//一个典型的顺序流水线算法可以用来解决这个问题，下面的示例展示了每一步执行的顺序：
//func SerialProcessData(in <- chan *Data, out <- chan *Data) {
//	for data := range in {
//		tmpA := PreprocessData(data)
//		tmpB := ProcessStepA(tmpA)
//		tmpC := ProcessStepB(tmpB)
//		out <- PostProcessData(tmpC)
//	}
//}
//一次只执行一步，并且每个项目按顺序处理：在第一个项目被预处理完并将结果放到输出通道之前第二个项目不会开始。
//如果仔细想想，你很快就会意识到这样会非常的浪费时间。
//一个更有效的计算是让每一步鄹都作为一个协程独立工作。每个步骤都从上一步的输出通道获取输入数据。这样可以尽可能的避免时间浪费，并且大部分时间所有的步骤都会繁忙的执行：
//func ParallelProcessData(in <- chan *Data, out <- chan *Data) {
//	//make channels
//	preOut := make(chan *Data, 100)
//	stepAOut := make(chan *Data, 100)
//	stepBOut := make(chan *Data, 100)
//	stepCOut := make(chan *Data, 100)
//	//start parallel computations
//	go PreprocessData(in, preOut)
//	go ProcessStepA(preOut, stepAOut)
//	go ProcessStepB(stepAOut, stepBOut)
//	go ProcessStepC(stepBOut, stepCOut)
//	go PostProcessData(stepCout, out)
//}

//通道缓冲区可以用于进一步优化整个过程。