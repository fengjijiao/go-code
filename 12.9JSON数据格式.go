package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//数据结构要在网络中传输或保存到文件，就必须对其编码和解码；目前存在很多编码格式：JSON、XML、Gob、Google缓冲协议等待。Go语言支持所有这些编码格式；在后面的章节，我们将讨论前3种格式。
//结构可能包含二进制数据，如果将其作为文本打印，那么可读性是很差的。另外结构内部可能包含匿名字段，而不清楚数据的用意。
//通过把数据转换成纯文本，使用命名的字段来标注，让其具有可读性。这样的数据格式可以通过网络传输，而且是与平台无关的，任何类型的应用都能够读取和输出，不与操作系统和编程语言的类型相关。
//下面是一些术语说明：
//1.数据结构 -> 指定格式 = 序列化 或 编码（传输之前）
//2.指定格式 -> 数据结构 = 反序列化 或 解码 （传输之后）
//JSON相比与XML格式非常简洁、轻量（占用更少的内存、磁盘、网络带宽）和更好的可读性。
//Go语言的json包可以使我们在程序中很方便的读取和写入json数据。

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

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"jan", "kersschot", []*Address{pa, wa}, "none"}
	fmt.Printf("%v\n", vc)
	js, _ := json.Marshal(vc)
	fmt.Printf("json format: %s\n", js)
	//using an encoder
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Println("error in encoding json")
	}
	//出于安全考虑，在web应用中最好使用json.MarshalforHTML()函数，其对数据执行HTML转码，所以文本可以被安全的嵌在HTML<script>标签中。
	//json.NewEncoder()的函数签名是func NewEncoder(w io.Writer) *Encoder，返回的Encoder类型的指针可调用方法Encode(v interface{})，将数据对象v的json编码写入io.Writer中。
	//JSON与Go类型对应如下：
	//bool 对应JSON的booleans
	//float64 对应JSON的numbers
	//string 对应JSON的strings
	//nil对应JSON的null
	//不是所有的数据都可以编码为json类型：只有验证通过的数据结构才能被编码：
	//1.JSON对象只支持字符串类型的KEY；要编码一个Go map类型，map必须是map[string]T(T是JSON包支持的任何类型)
	//2.Channel，复杂类型和函数类型不能被编码
	//不支持循环数据结构；它将引起序列化进入一个无限循环
	//指针可以被编码，实际上是对指针的值进行编码（或者指针是nil）
	//


	//反序列化
	//UnMarshal()的函数签名是func UnMarshal(data []byte, v interface{}) error，把JSON解码为数据结构。
	//对vc编码后的数据为js，对其解码时，我们首先创建结构VCard用来保存解码的数据：var v VCard并调用json.UnMarshal(js, &v)，解析[]byte中的JSON数据并将结果存入指针&v指向的值。
	//虽然反射能够让JSON字段去尝试匹配目标结构字段；但是只有真正匹配上的字段才会填充数据。字段没有匹配不会报错，而是直接忽略掉。



	//解码任意的数据：
	//json包使用map[string]interface{}和[]interface{}存储任意的JSON对象和数组；其可以被反序列化为任何的JSON blob存储到接口值中。
	//b := []byte(`{"name":"ok", "data":{"status":2}}`)
	//不用理解这个数据的结构，我们可以直接使用UnMarshal把这个数据编码并保存在接口值中：
	//var f interface{}
	//err := json.UnMarshal(b, &f)
	//f指向的值是一个map，key是一个字符串，value是本身存储作为空接口类型的值。
	//map[string]interface{} {
	//	"name":"ok",
	//	"data": {
	//		"status":2
	//	}
	//}
	//要访问这个数据，我们可以使用类型断言
	//m := f.(map[string]interface{})
	//我们可以通过for-range语法和type switch来访问其实际类型
	//for k,v := range m {
	//	switch vv := v.(type) {
	//	case string:
	//		fmt.Println(k, "is string", vv)
	//	case int:
	//		fmt.Println(k, "is int", vv)
	//	case []interface{}:
	//		fmt.Println(k, "is an array")
	//		for i, u := range vv {
	//			fmt.Println(i, u)
	//		}
	//	default:
	//		fmt.Println(k, "is of a type I dont know how to handle")
	//}
	//通过这种方式，你可以处理未知的JSON数据，同时可以确保类型安全。
	//



	//解码数据到结构
	//如果我们事先直到JSON数据，我们可以定义一个适当的结构并对JSON数据反序列化。
	type FamilyMember struct {
		Name string
		Age int
		Parents []string
	}
	//并对其反序列化
	//var m FamilyMember
	//err := json.Unmarshal(b, &m)
	//程序实际上是分配了一个新的切片。这是一个典型的反序列化引用类型（指针、切片和map）的例子。




	//编码和解码流
	//json包提供了Decoder和Encoder类型来支持常用JSON数据流的读写。NewDecoder和NewEncoder函数分别封装了io.Reader和io.Writer接口。
	//func NewDecoder(r io.Reader) *Decoder
	//func NewEncoder(r io.Writer) *Encoder
	//要想把JSON直接写入文件，可以使用json.NewEncoder初始化文件（或任何实现了io.Writer的类型），并调用Encode()；反过来与其对应的是使用json.Decoder和Decode()函数。
	//func NewDecoder(r io.Reader) *Decoder
	//func (dec *Decoder) Decode(v interface{}) error
	//来看下接口是如何对实现进行抽象的：数据结构可以是任何类型，只要其实现了某种接口，目标或源数据要能够被编码就必须实现io.Writer或io.Reader接口。由于Go语言中到处都实现了Reader和Writer，因此Encoder和Decoder可被应用的场景非常广泛，例如读取或写入HTTP连接、websocket或文件。
}
