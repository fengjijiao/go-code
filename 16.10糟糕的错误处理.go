package main

import (
	"errors"
	"io"
	"net/http"
)

//13,17.1,17.2.4总结

//16.10.1 不要使用布尔值
//像下面的代码，创建一个布尔型变量用于测试错误条件是多余的。
//var good bool//测试一个错误
//if !good {
//	return errors.New("things are not good")
//}
//立即检查一个错误
//..., err1 := api.Func1()
//if err1 != nil { ... }



//16.10.2避免错误检测使代码变得混乱
//避免写出这样的代码
..., err1 := api.Func1()
if err1 != nil {
	fmt.Println("err: "+err1.Error())
	return
}
err2 := api.Func2()
if err2 != nil {
fmt.Println("err: "+err2.Error())
return
}
//首先，包括在一个初始化的if语句中对函数的调用。但即使代码中到处都是以if语句的形式通知错误（通过打印错误信息）。通过这种方式，很难分辨什么是正常的程序逻辑，什么是错误检测或错误通知。还需要注意的是，大部分代码都是致力于错误的检测。通常解决此问题的好办法是尽可能以闭包的形式封装你的错误检测，例如：
func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	err := func() error {
		if r.Method != "GET" {
			return errors.New("expected GET")
		}
		if input := parseInput(r); input != "command" {
			return errors.New("malformed command")
		}
		//可在这里进行其他错误检测
		return nil
	}()
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, err)
		return
	}
	doSomething()
	//...
}
//这种方法可以很容易分辨出错误检测、错误通知和正常的程序逻辑。