package main

import (
	"fmt"
	"os"
	"text/template"
)

//在上一节中，我们使用了模板去合并结构体与html模板中的数据。这对构建web应用程序确实非常有用，但是模板技术比这更通用：数据驱动模板可以用于生成文本输出，HTML仅仅是其中一个特例。

//将结构传递给template.Execute()可以替换生成新的内容。

//1.字段替代{{.FieldName}}
//使用template.New创建一个新的模板，需要一个string参数作为模板的名称
//Parse方法通过解析一些模板定义的字符串来生成一个template作为内部表示。
//当参数是一个定义好的模板路径时，使用ParseFile。

type Person struct {
	Name string
	nonExportedAgeField string//不可导出，没有添加这个，代码会报错
}

func main2() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!")//使用{{.}}可以将两个字段都导出
	p := Person{
		Name: "Mary",
		nonExportedAgeField: "28",
	}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("there was an error: ", err.Error())
	}
}
//当模板应用于浏览器时，要先用html过滤器去过滤输出的内容：{{html .}}，或者使用一个FieldName: {{.FieldName |html}}。
//|html告诉template在输出FieldName的值之前要通过html格式化它。它会转义特殊的html字符。防止用户数据破坏html表单。
//|html用起来和Linux中的|(管道)类似，将前面命令的输出作为|后面命令的输入。


//15.7.2模板验证
//检查模板的语法是否正确，对Parse的结果执行Must函数。在下面的示例中tOK是正确的，tErr的验证会出错并会导致一个panic。
func main3() {
	tOK := template.New("ok")
	//一个有效的模板，所以Must的时候不会出现panic
	template.Must(tOK.Parse("/*这是一个注释*/some static text: {{ .Name }}"))
	fmt.Println("the first one parsed OK.")
	fmt.Println("the next one ought to fail.")
	tErr := template.New("err_template")
	template.Must(tErr.Parse("some static text {{ .Name }"))//缺少一个}
}


//下面3个基本函数在代码中经常被链接使用
//var strTempl = template.Must(template.New("TName").Parse(strTemplateHTML))




//15.7.3 If-else
func main4() {
	t := template.New("test")
	//t = template.Must(t.Parse("will be {{`ok`}}"))
	//使用if-else-end来调整管道数据的输出
	//t = template.Must(t.Parse("will be {{if ``}} failed {{end}}"))//failed
	t = template.Must(t.Parse("will be {{if `an`}} ok {{else}}failed {{end}}"))//ok
	t.Execute(os.Stdout, nil)
}


//15.7.4 点于with-end
//在Go模板中使用(.)：它的值{{.}}被设置为当前管道的值。
//with语句将点的值设置为管道的值。如果管道是空的，就会跳过with到end之前的内容；当嵌套使用时，点会从最近的范围取值。
func main5() {
	t := template.New("test")
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")//hello!
	t.Execute(os.Stdout, nil)
	t, _ = t.Parse("{{with `hello`}}{{.}}{{with `Mary`}}{{.}}{{end}}{{end}}!\n")//hello Mary!
	t.Execute(os.Stdout, nil)
}


//15.7.5 模板变量$
//你可以在变量名前加一个$符号来为模板中的管道创建一个局部变量。变量名称只能由字母、数字、下划线组成。在下面的示例中，我使用了几种可以使用的变量名称。
func main6() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))//hello!
	t.Execute(os.Stdout, nil)
	t =template.Must(t.Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))//hola!
	t.Execute(os.Stdout, nil)
	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}}{{.}}{{$x_1}}{{end}}!"))//hey hey hey!
}



//15.7.6 range-end
//这个构造的格式
//{{range pipeline}} T1 {{else}} T0 {{end}}
//range 在循环的集合中使用：管道的值必须是一个数字、切片或者map。如果管道的值的长度为0,点不会被影响并且T0将会被执行；否则将点设置为拥有连续元素的数组、切片或者map，T1就会被执行。
//如果它是模板：{{range .}}
//				{{.}}
//				{{end}}
//然后是这个代码：s := []int{1,2,3,4}
//t.Execute(os.Stdout, s)
//将会输出:
//			1
//			2
//			3
//			4
//20.7中，
/*
{{range .}}
	{{with .Author}}
		<p><b>{{html .}}</b> wrote:</p>
	{{else}}
		<p>An anonymous person wrote:</p>
	{{end}}
	<pre>{{html .Content}}</pre>
	<pre>{{html .Date}}</pre>
{{end}}
 */
//range . 这里循环了一个结构体切片，每个结构体都包含了一个Author、Content和Date字段。





//15.7.7 预定义模板函数
//还可以在代码中使用一些模板函数，例如：和fmt.Printf函数类似的printf函数：
func main7() {
	t :=template.New("test")
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}


//在15.6中也这样使用过
//{{printf "%s" .Body|html}}
//否则Body的字节会被当作数字显示（字节默认都是int8类型的数字）