package main

import (
	"fmt"
	"net/http"
	"strings"
)

type OBJ struct {}

func (_ *OBJ) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	HelloServer(w, req)
}

// simple http server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprint(w, "Hello, "+req.URL.Path[1:])// get /world, return "Hello, world"
	//fmt.Fprintf(w, "hello, %s", req.URL.Path[1:])
}

func main2() {
	var obj *OBJ = new(OBJ)
	http.HandleFunc("/", HelloServer)
	http.Handle("/obj", obj)//obj实现了http的Handler接口ServeHTTP
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

//如果req是一个POST请求类型的form，可以通过以下方式获取表单内的值
//1.req.FormValue("var1")
//2.req.ParseForm();var1, found := req.Form["var1"]
//表单属性实际上是一个map[string][]string类型。


//练习15.2
func HelloServerV2(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "hello %s", req.URL.Path[1+5+1:])
	fmt.Fprintf(w, "hello %s", strings.ToUpper(req.URL.Path[1+5+1:]))
}

func main() {
	http.HandleFunc("/hello/", HelloServerV2)
	http.ListenAndServe(":8000", nil)
}