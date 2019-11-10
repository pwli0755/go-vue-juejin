package service

import (
	"bytes"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"net/textproto"
	"testing"
)

func TestSendEmail(t *testing.T) {
	tp, err := template.New("activeEmailTpl").Parse(tpl)
	if err != nil {
		t.Fatal(err)
	}
	data := struct {
		Name      string
		ActiveUrl string
		Date      string
	}{
		Name:      "pwli",
		ActiveUrl: "http://localhost:8080/#/email/confirm?token=",
		Date:      "2019-11-11",
	}
	buffer := &bytes.Buffer{}
	err = tp.Execute(buffer, data)
	if err != nil {
		t.Fatal(err)
	}
	e := &email.Email{
		To:      []string{"763045276@qq.com"},
		From:    "go-vue <go_vue@163.com>",
		Subject: "go-vue邮件中心",
		//Text:    []byte("测试text"),
		HTML:    buffer.Bytes(),
		Headers: textproto.MIMEHeader{},
	}
	err = e.Send("smtp.163.com:25", smtp.PlainAuth("", "go_vue@163.com", "fakepasswd", "smtp.163.com"))
	if err != nil {
		t.Fatal(err)
	}
}
