package main

import (
	"io"
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

func main2() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}


//练习15.6编写一个网页程序，可以让用户输入一串的数字，然后将它们打印出来，计算出这些数字的均值和中值。
//...