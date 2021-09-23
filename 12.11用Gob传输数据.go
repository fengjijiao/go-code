package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

//Gob是Go自己的以二进制形式序列化和反序列化程序数据的格式；可以在encoding包中找到。这种格式的数据简称为Gob(Go binary)。类似于Python的pickle和Java的Serialization.
//Gob通常用于远程方法调用（RPCs, 15.9 rpc包）参数和结果的传输，以及应用程序和机器之间的数据传输。
//它和JSON、XML有什么不同呢？
//Gob特定地用于纯Go的环境中，例如两个用Go写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。
//Gob不是可外部定义，语言无关的编码方式。因此它首选格式是二进制，而不是像JSON、XML那样的文本格式。Gob并不是一种不同于Go的语言，而是在编码和解码过程种用到了Go的反射。
//只有可导出的字段会被编码，零值会被忽略。在解码结构体的时候，只有同时匹配名称和可兼容类型的字段才会被解码。当源数据类型增加新字段后，Gob解码客户端仍然可以以这种方式正常工作：解码客户端会继续识别以前存在的字段。并且还提供了很大的灵活性，比如在发送者看来，整数被编码成没有固定长度的可变长度，而忽略具体的Go类型。
//假如在发送者这边有一个结构T
//type T struct {
//	X,Y,Z int
//}
//var t = T{7,0,8}
//而在接收者这边可以用一个结构体U类型的变量u来接收这个值
//type U struct {
//	X,Y *int8
//}
//var u U
//在接收者中，X的值是7，Y的值是0（Y的值并没有从t中传递过来，因为它是零值）
//和JSON的使用方式一样，Gob使用通用的io.Writer接口，通过NewEncoder()函数创建Encoder对象并调用Encode()；相反的过程使用io.Reader接口，通过NewDecoder()函数创建Decoder对象并调用Decode。

//你将会看到一个编解码，并且以字节缓冲模拟网络传输的简单例子。
type P struct {
	X, Y, Z int
	Name string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	err := enc.Encode(P{3,4,5,"ok"})
	if err != nil {
		panic(err)
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	/**
	"ok": {3, 4}
	*/
}

//编码到文件

type Address struct {
	Type string
	City string
	Country string
}

type VCard struct {
	FirstName string
	LastName string
	Addresses []*Address
	Remark string
}

var content string
func main() {
	pa := &Address{"private", "aart", "belgium"}
	wa := &Address{"work", "boom", "belgium"}
	vc := VCard{"jan", "kerss",[]*Address{pa, wa}, "none"}
	fmt.Printf("%v\n", vc)
	//encoder
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	//defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		panic(err)
	}
	file.Close()
	//decoder
	file, _ = os.OpenFile("vcard.gob", os.O_RDONLY, 0)
	dec := gob.NewDecoder(file)
	var vc2 VCard
	err = dec.Decode(&vc2)
	if err != nil {
		panic(err)
	}
	fmt.Println(vc2)
	file.Close()
	/**
	{jan kerss [0xc000100db0 0xc000100de0] none}
	{jan kerss [0xc000101560 0xc000101590] none}
	*/
}