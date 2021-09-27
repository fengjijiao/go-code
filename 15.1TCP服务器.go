package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"syscall"
)

//将使用TCP协议和在14中的协程范式编写一个简单的客户端-服务器程序，一个web服务器应用需要响应众多客户端的并发请求：go会为每一个客户端产生一个协程用来处理请求。我们需要使用net包中的网络通信功能。它包含了用于TCP/IP协议、域名解析等
////server.go
func doServerStuff(conn net.Conn) {
	for{
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("received data: %v\n", string(buf[:len]))
	}
}

func runServer() {
	fmt.Println("starting the server...")
	//创建listener
	listener, err := net.Listen("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	//监听并接受来自客户端的请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

////client.go
func runClient() {
	//打开连接
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")//windows, linux: \n
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient +" says: "+ trimmedInput))
	}
}





//tcp, udp, ipv6 socket
func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("err %v connecting", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v\n", conn)
}

func main2() {
	conn, err := net.Dial("tcp", "192.168.100.1:80")//tcp ipv4
    checkConnection(conn, err)
	conn, err = net.Dial("udp", "192.168.100.1:53")//udp ipv4
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "[240c::6666]:53")
}





//下面是一个使用net包从socket中打开，写入，读取数据的例子：
func main3() {
	var (
		host = "www.fengjijiao.com"
		port = "80"
		remote = host +":"+port
		msg string = "GET / \n"
		data = make([]uint8, 4096)
		read = true
		count = 0
	)
	conn, err := net.Dial("tcp", remote)
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, msg)
	for read {
		count, err = conn.Read(data)
		read = err == nil
		fmt.Printf("%v", string(data[0:count]))
	}
	conn.Close()
	/**
	HTTP/1.1 400 Bad Request
	Content-Type: text/plain; charset=utf-8
	Connection: close

	400 Bad Request
	*/
}




//练习15.1
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Client struct {
	Name string
	Online bool
}

type Clients []*Client

func (c Clients) IndexOf(name string) int {
	for i2, c2 := range c {
		if c2.Name == name {
			return i2
		}
	}
	return -1
}

//func (c *Clients) Unregister() chan *Client {
//	var a chan *Client = make(chan *Client)
//	go func() {
//		for v := range a {
//			i := c.IndexOf(v.Name)
//			if i >= 0 {
//				*c = append((*c)[:i], (*c)[i+1:]...)
//			}
//		}
//	}()
//	return a
//}
func (c *Clients) Unregister() chan *Client {
	var a chan *Client = make(chan *Client)
	go func() {
		for v := range a {
			i := c.IndexOf(v.Name)
			if i >= 0 {
				(*c)[i].Online = false
			}
		}
	}()
	return a
}

func (c *Clients) Register() chan *Client {
	var a chan *Client = make(chan *Client)
	go func() {
		for v := range a {
			(*c)[len(*c)] = v
		}
	}()
	return a
}

//var runFlag = true

func runServerV2() {
	var clients Clients = make(Clients, 0)
	ch1 := clients.Register()
	ch2 := clients.Unregister()
	listener, err := net.Listen("tcp", ":8000")
	checkError(err)
	//for runFlag {
	for {
		conn, err := listener.Accept()
		checkError(err)
		go doServerStuffV2(conn, ch1, ch2, clients)
	}
}

func doServerStuffV2(conn net.Conn, ch1 chan *Client, ch2 chan *Client, clients Clients) {
	var client *Client
	defer func() {
		ch2 <-client
	}()
	var buf = make([]byte, 512)
	len, err := conn.Read(buf)
	checkError(err)
	client = &Client{
		Name: string(buf[:len]),
		Online: true,
	}
	ch1 <-client
	for {
		len, err = conn.Read(buf)
		checkError(err)
		trimmedData := strings.Trim(string(buf[:len]), "\r\n")
		if trimmedData == "SH" {
			//runFlag = false
			//break
			os.Exit(1)
		} else if trimmedData == "WHO" {
			var data string = "thid is the client list:  1:active, 0:inactive\r\n"
			for _, c := range clients {
				data += fmt.Sprintf("User %s is %t\r\n", c.Name, c.Online)
			}
			len, err = conn.Write([]byte(data))
			checkError(err)
		} else {
			fmt.Println("data: ", string(buf[:len]))
			len, err = conn.Write(buf[:len])
			checkError(err)
		}
	}
}

//client similar





//优化后的tcp服务端
const maxRead = 25

func initServer(hostAndPort string) net.Listener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err)
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err)
	fmt.Println("listening to :", listener.Addr().String())
	return listener
}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', ',', 'g', 'o', '!'}
	wrote, err := to.Write(obuf)
	checkError(err)
	fmt.Println("wrote ", string(wrote), "bytes.")
}

func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	fmt.Println("connection from :", connFrom)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN://try again
			continue
		default:
			goto DISCONNECT
		}
	}
	DISCONNECT:
		err := conn.Close()
		println("close connection: ", connFrom)
		checkError(err)
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		fmt.Print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		fmt.Print(">")
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err)
		go connectionHandler(conn)
	}
}