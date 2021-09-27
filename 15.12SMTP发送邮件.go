package main

import (
	"bytes"
	"net/smtp"
)

//smtp包实现了一个简单的邮件传输协议来发送邮件。
//它包含了一个代表客户端连接到SMTP服务器的Client类型：
//1.Dial返回一个用于连接到SMTP服务器的客户端
//2.设置Mail(=寄件人)和Rcpt(=收件人)
//3.Data返回一个可以写入数据的writer,这里用buf,WriteTo(wc)写入

func sendSimple() {
	//连接
	client, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		panic(err)
	}
	//设置寄件人和收件人
	client.Mail("sender@example.org")
	client.Rcpt("recipient@example.net")
	//发送邮件主体
	wc, err := client.Data()
	if err != nil {
		panic(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("this is the mail body.")
	if _, err = buf.WriteTo(wc); err != nil {
		panic(err)
	}
}


//如果需要权限认证并且有多个收件人，可以使用SendMail函数。
func sendAuth() {
	//设置认证信息
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
		)
	//连接到服务器，认证，设置发件人，收件人，发送的内容
	//然后发送邮件
	err := smtp.SendMail(
		"mail.example.com:25",
		auth,
		"sender@example.org",
		[]string{"recipient@example.net"},
		[]byte("this is the email body."),
		)
	if err != nil {
		panic(err)
	}
}
//SendMail的源码要求msg参数需要符合RFC822电子邮件的标准格式。
//"To: recipient@example.net\r\nFrom: sender@example.org\r\nSubject: 邮件主题\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\nHello World"