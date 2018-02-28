package main

import (
    "net/smtp"
    "strings"
)

const (
    HOST        = "mail.bjsasc.com"
    SERVER_ADDR = "mail.bjsasc.com:25"
    USER        = "netdisk@bjsasc.com" //发送邮件的邮箱
    PASSWORD    = "Shaomina1984"         //发送邮件邮箱的密码
)

type unencryptedAuth struct {
    smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    s := *server
    s.TLS = true
    return a.Auth.Start(&s)
}



type Email struct {
    to      string "to"
    subject string "subject"
    msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
    return &Email{to: to, subject: subject, msg: msg}
}

func (email *Email)SendEmail() error {
    //auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
    auth := unencryptedAuth {
    smtp.PlainAuth(
        "",
        USER,PASSWORD,
        HOST,
    ),
}
    sendTo := strings.Split(email.to, ";")
    done := make(chan error, 1024)

    go func() {
        defer close(done)
        for _, v := range sendTo {

            str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg

            err := smtp.SendMail(
                SERVER_ADDR,
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


//测试方法，可放倒main中


func SendTest() {
	mycontent := " 你好："+"\r\n"+
	"      这是一封测试邮件"

	email := NewEmail("fengshaomin@bjsasc.com",
		"test golang email", mycontent)

	println(email.to)

	if err := email.SendEmail(); err != nil {

		println(err)
	}
}
