package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
	"os"
	"testing"
)

func TestSendActiveEmail(t *testing.T) {
	e := &email.Email{
		To:      []string{"763045276@qq.com"},
		From:    "go-vue <go_vue@163.com>",
		Subject: "go-vue邮件中心",
		Text:    []byte("测试text"),
		HTML:    []byte("<h1>测试HTML!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	//err:=e.SendWithTLS("smtp.163.com:465", smtp.CRAMMD5Auth("go_vue@163.com", os.Getenv("MAIL_PW")), &tls.Config{InsecureSkipVerify:true})
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "go_vue@163.com", os.Getenv("MAIL_PW"), "smtp.163.com"))
	if err != nil {
		t.Fatal(err)
	}
}
