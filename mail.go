//1111
package main

import (
	"net/smtp"
	"strings"
)

const (
	//HOST 邮件服务器的地址
	HOST = "mail.bjsasc.com"
	//SERVERADDR 邮件服务器的地址和端口
	SERVERADDR = "mail.bjsasc.com:25"
	//USER 发送邮件的邮箱
	USER = "netdisk@bjsasc.com"
	//PASSWORD 发送邮件邮箱的密码
	PASSWORD = "Shaomina1984"
)

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

//Email userd to
type Email struct {
	//11111
	To      string
	Subject string
	Msg     string
}

/*NewEmail 用于初始化mail*/
func NewEmail(to, subject, msg string) *Email {
	//初始化邮件对象
	return &Email{To: to, Subject: subject, Msg: msg}
}

//SendEmail 发送邮件方法
func (email *Email) SendEmail() error {
	//auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	auth := unencryptedAuth{
		smtp.PlainAuth(
			"",
			USER, PASSWORD,
			HOST,
		),
	}
	sendTo := strings.Split(email.To, ";")
	done := make(chan error, 1024)

	go func() {
		defer close(done)
		for _, v := range sendTo {

			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.Subject+"~~", "~", "\r\n", -1) + email.Msg

			err := smtp.SendMail(
				SERVERADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			done <- err
		}
	}()

	for i := 0; i < len(sendTo); i++ {
		<-done
	}

	return nil
}

// SandTest sandtest
func SandTest() {
	//发送测试邮件
	mycontent := " 你好：" + "\r\n" +
		"      这是一封测试邮件"

	email := NewEmail("fengshaomin@bjsasc.com",
		"test golang email", mycontent)

	println(email.To)

	if err := email.SendEmail(); err != nil {

		println(err)
	}
}
