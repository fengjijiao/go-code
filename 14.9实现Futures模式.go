package main

//所谓Futures指：有时候你使用某一值之前需要先对其进行计算。这种情况下，你就可以在另一个处理器上进行该值的计算，到使用时，该值就已经计算完毕了。
//Futures模式通过闭包和通道可以很容易实现，类似于生成器，不同地方在于Futures需要返回一个值。
//例子：假设我们有一个矩阵类型，我们需要计算两个矩阵A和B乘积的逆，首先我们通过函数Inverse(M)分别对其进行求逆运算，在将结果相乘。如下函数InverseProduct()实现了如上过程：
//func InverseProduct(a Matrix, b Matrix) {
//	a_inv := Inverse(a)
//	b_inv := Inverse(b)
//	return Product(a_inv, b_inv)
//}

//在这个例子中，a和b的求逆需要先被计算。那么为什么在计算b的逆矩阵时，需要等待a的逆计算完成呢？显然不必要，这两个求逆运算其实可以并行执行的。如下代码实现了并行计算方式：
//func InverseProduct(a Matrix, b Matrix) {
//	a_inv_future := InverseFutrue(a)//start as a goroutine
//	b_inv_future := InverseFutrue(b)//start as b goroutine
//	a_inv := <-a_inv_future
//	b_inv := <-b_inv_future
//	return Product(a_inv, b_inv)
//}

//InverseFuture函数起了一个goroutine协程，在其执行闭包运算，该闭包会将矩阵求逆结果放入到future通道中
//func InverseFuture(a Matrix) chan Matrix {
//	future := make(chan Matrix)
//	go func() {
//		future <- Inverse(a)
//	}()
//	return future
//}

//当开发一个计算密集型库时，使用Futures模式设计API接口是很有意义的。