package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

/**
<Person>
	<FirstName>laura</FirstName>
	<LastName>lynn</LastName>
</Person>
 */
//如同json包一样，也有Marshal()和UnMarshal()从XML中编码和解码数据；但这个更通用，可以从文件中读取和写入（或者任何实现了io.Reader和io.Writer接口的类型）
//encoding/xml包实现了一个简单的xml解析器（SAX）。用来解析XML数据内容。

var t, token xml.Token
var err error

func main() {
	input := `<Person>
	<FirstName>laura</FirstName>
	<LastName>lynn</LastName>
</Person>`
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("an attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("end of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("this is the content: %v\n", content)
		default:
			//...
		}
	}
	/**
	token name: Person
	this is the content:

	token name: FirstName
	this is the content: laura
	end of token
	this is the content:

	token name: LastName
	this is the content: lynn
	end of token
	this is the content:

	end of token
	*/
	//包中定义了若干XML标签类型：StartElement,Chardata（这是从开始标签到结束标签之间的实际文本），EndElement,Comment,Directive或ProcIns。
	//包中同样定义了一个结构解析器：NewParser方法持有一个io.Reader（这里具体类型是strings.NewReader）并生成一个解析器类型的对象。还有一个Token()方法返回输入流里的下一个XML token。在输入流的结尾处，会返回(nil, io.EOF)
	//XML文本被循环处理直到Token()返回一个错误，因为已经达到文件尾部，再没有内容可供处理了。通过一个type-switch可以根据一些XML标签进一步处理。Chardata中的内容只是一个[]byte,通过字符串转换让其变得可读性强一些。
}
