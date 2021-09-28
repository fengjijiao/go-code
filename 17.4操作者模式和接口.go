package main

//运算符是个一元或者二元函数，它返回一个新的对象并且不能修改它的参数，比如+和*。在C++中，可以重载中缀运算符（+,-,*等）用来实现数学类语法，但是除了一些特殊情况，Go不支持运算符重载：为了克服这个限制，运算符必须用函数进行模拟。由于Go支持程序以及一个面向对象的范例，因此有两种方案：

//17.4.1用函数实现运算符
//运算符被实现为了一个包级别的函数，在专用于它们所在的对象的包中，它去操作一个或两个参数，并返回一个新的对象。
//例如：我们如果想在一个matrix包中实现矩阵操作，在matrix的结果中，他要包含添加矩阵的Add()和相乘的Mult()。这些将用包本身的名称去调用，所以我们可以这样使用：m := matrix.Add(m1, matrix.Mult(m2, m3))
//如果我们想在这个操作中区分不同的矩阵（sparse、dense），因为它不能函数重载，我们要给它们不同的名字，例如：
func addSpareToDense(a *sparseMatrix, b *denseMatrix) *denseMatrix
func addDenseToDense(a *denseMatrix, b *denseMatrix) *denseMatrix
func addSpareToSpare(a *sparseMatrix, b *sparseMatrix) *sparseMatrix
//这个非常不优雅，我们最好能作为一个私有的函数隐藏这些，并且通过一个单独的公共函数Add()去暴露它们。可以通过嵌套switch type对它们进行类型检测，来操作任意组合的被支持的参数:
func Add(a Matrix, b Matrix) Matrix {
	switch a.(type) {
	case sparseMatrix:
		switch b.(type) {
		case sparseMatrix:
			return addSpareToSpare(a.(sparseMatrix), b.(sparseMatrix))
		case denseMatrix:
			return addSpareToDense(a.(sparseMatrix), b.(denseMatrix))
		}
	default:
		//不支持的参数
	}
}
//但是更加优雅和首选的做法是将运算符作为一个方法去实现，因为它在标准库中任何地方都可以完成。code.google.com/p/gomatrix/



//17.4.2用方法实现运算符
//方法可以根据它们的接收器类型进行区分，所以不必使用不同名称的函数（上一小节的方法），我们可以简单的为每种类型定义一个Add方法
func (a *sparseMatrix) Add(b Matrix) Matrix
func (a *denseMatrix) Add(b Matrix) Matrix
//每个方法都会返回一个新对象，该对象将成为下一个方法调用的接收者，因此我们可以创建链式表达式：m1.Mult(m2).Add(m3)
//这种方法比17.4.1的程序更简洁、清晰。
//基于type-switch，正确的实现可以在运行时再次被选择
func (a *sparseMatrix) Add(b Matrix) Matrix {
	switch b.(type) {
	case sparseMatrix:
		return addSpareToSpare(a.(sparseMatrix), b.(sparseMatrix))
	case denseMatrix:
		return addSpareToDense(a.(sparseMatrix), b.(denseMatrix))
	default:
		//不支持的参数
	}
}
//比17.4.1章节中的type switch更简洁




//17.4.3使用接口
//当在不同类型中使用相同的方法进行操作时，应该想到创建一个泛化接口去实现这种多态性：
//例如，我们可以定义接口Algebraic
type Algebraic interface {
	Add(b Algebraic) Algebraic
	Min(b Algebraic) Algebraic
	Mult(b Algebraic) Algebraic
	//...
	Elements()
}
//并为我们的matrix类型定义方法Add()、Min()、Mult()......实现上述Algebraic接口的每种类型都将允许方法链接。每个方法的实现都应该根据参数类型使用type-switch来提供优化的实现。此外，应指定一个仅依赖于接口中方法的默认情况：
func (a *denseMatrix) Add(b Algebraic) Algebraic {
	switch b.(type) {
	case sparseMatrix:
		return addDenseToSparse(a,b.(sparseMatrix))
	default:
		for x in range b.Elements()
		//...
	}
}
//如果通用实现不能仅使用接口中的方法实现，你可能正在处理那些不够相似的类，这种操作模式应该被抛弃。例如：
//如果a是一个set、b是一个matrix，写一个a.Add(b)就不太合理了；因此，在一个set和matrix操作条件中，实现一个通用的a.Add(b)非常困难。在这种情况下，将你的包分为两个部分，并定义单独的AlgebraicSet和AlgebraicMatrix接口。