package main

import (
	"io"
	"log"
	"net/http"
)
const form = `
<html><body>
<form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="submit"/>
</form>
</body></html>`

func SimpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>hello, world!</h1>")
}

func FormServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, req.FormValue("in"))
	}
}

//当web服务器发生一个panic时，web服务器就会终止。这样非常糟糕：一个web服务必须是一个健壮的程序，能够处理可能会出现的问题。
//一个方法是可以在每一个处理函数(handler)中去使用defer/recover，但是这样会导致出现很多重复的代码，更加优雅的解决方法是使用13.5中的闭包方法处理错误。
//为了使代码更具可读性，我们为处理函数(HandleFunc)创建一个type
type HandleFnc func(w http.ResponseWriter, r *http.Request)
//模仿13.5中的errorHandler函数，创建一个logPanics函数
func logPanics(function HandleFnc) HandleFnc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		function(w, r)
	}
}
//然后就可以将处理函数作为回调包装进log.Panics:
//http.HandleFunc("/test1", logPanics(SimpleServer))
//http.HandleFunc("/test2", logPanics(FormServer))
//处理函数中应该包含一个panic调用，或者像1.3中的用check(error)函数。
func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}