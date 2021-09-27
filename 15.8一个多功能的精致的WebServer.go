package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

//hello world计数器
var helloRequests = expvar.NewInt("hello-requests")

//flags:
var webroot = flag.String("root", "/home/user", "web root directory")
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

//简单的计数器
type Counter struct {
	n int
}
type Chan chan int

func main() {
	flag.Parse()
	http.Handle("/", http.HandleFunc(Logger))
	http.Handle("/go/hello", http.HandleFunc(HelloServer))
	//计数器直接作为一个变量被发布
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))
	http.Handle("/flags", http.HandleFunc(FlagServer))
	http.Handle("/args", http.HandleFunc(ArgServer))
	http.Handle("/chan", ChanCreate())
	http.Handle("/date", http.HandleFunc(DateServer))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func Logger(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello,world!\n")
}

//通过这个方法满足expvar.Var接口，所以就可以直接导出它。
func (ctr *Counter) String() string {
	return fmt.Sprintf("%d", ctr.n)
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": //给n加1
		ctr.n++
	case "POST": //设置n值去发布
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST %v\n body: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprintf(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func FlagServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	fmt.Fprintf(w, "Flags:\n")
	flag.VisitAll(func(f3 *flag.Flag) {
		if f3.Value.String() != f3.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f3.Name, f3.Value.String(), f3.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f3.Name, f3.Value)
		}
	})
}

func ArgServer(w http.ResponseWriter, r *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func() {
		for x := 0;;x++ {
			c <- x
		}
	}()
	return c
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("channel sned #%d\n", <-ch))
}

//执行一个程序输出重定向
func DateServer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r,w,err := os.Pipe()
	if err != nil {
		fmt.Fprintf(w, "pipe: %s\n", err)
		return
	}
	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil,w,w}})
	defer r.Close()
	w.Close()
	if err != nil {
		fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	io.Copy(rw, r)
	wait,err := p.Wait()
	if err != nil {
		fmt.Fprintf(rw,"wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		fmt.Fprintf(rw, "date: %v\n", wait)
	}
}
