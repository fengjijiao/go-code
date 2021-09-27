package main

//nexter是一个接口类型，并且定义了一个next()方法读取下一字节。函数nextFew1将nexter接口作为参数并读取接下来n个字节，并返回一个切片；这是正确的做法。但是nextFew2是使用一个指向nexter接口类型的指针作为参数传递给函数：当使用next()函数时，系统会给出一个编译错误：n.next() undefined(type *nexter has no field or method next)

type nexter interface {
	next() byte
}

func nextFew1(n nexter, num int) []byte {
	var b []byte
	for i := 0; i< num;i++ {
		b[i] = n.next()
	}
	return b
}

func nextFew2(n *nexter, num int) []byte {
	var b []byte
	for i := 0; i< num;i++ {
		b[i] = n.next()//编译错误：n.next未定义(*nexter类型没有next方法或next成员)
	}
	return b
}


//永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
