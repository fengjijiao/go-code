package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//下面，数值中的url都将被访问：或发送一个简单的http.Head()请求查看返回值；它的声明如下：func Head(url string)(r *Response, err error)。
//返回状态码会被打印出来。

var urls = []string {
	"http://www.bing.com/",
	"http://www.golang.org/",
	"http://www.google.com.hk/",
}

func main2() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			panic(err)
		}
		fmt.Println(url, ":", resp.Status)
	}
}

//下面我们使用http.Get()获取网页内容；Get的返回值res中的Body属性包含了网页内容，然后我们用ioutil.ReadAll把它读出来：
func main() {
	res, err := http.Get("http://www.google.com.hk/")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got: %q\n", string(data))
	//if res.Body is json/xml format
	//xml.Unmarshal(res.Body, &user)
	//json.Unmarshal(res.Body, &user)
}

//http.Redirect让浏览器重定向到url
//http.NotFound返回网页没有找到404
//http.Error返回特定错误信息和http代码


//可以使用w.header().Set("Accept", "zh-cn“)设置头部
