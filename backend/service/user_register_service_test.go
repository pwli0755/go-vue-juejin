package service

import (
	"bytes"
	"crypto/tls"
	"github.com/pwli-star/email"
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
		ActiveUrl: "http://47.105.53.57/#/email/confirm?token=",
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
	//err = e.Send("smtp.163.com:25", smtp.PlainAuth("", "go_vue@163.com", "fakepasswd", "smtp.163.com"))
	auth := smtp.PlainAuth("", "go_vue@163.com", "", "smtp.163.com")
	err = e.SendWithTLS("smtp.163.com:465", auth, &tls.Config{ServerName: "smtp.163.com",InsecureSkipVerify:true})
	if err != nil {
		t.Fatal(err)
	}
}

