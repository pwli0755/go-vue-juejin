package service

import (
	"backend/model"
	"backend/util"
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"net/textproto"
	"os"
	"time"
)

const tpl = `
<div style="background-color:#ECECEC; padding: 35px;">
    <table cellpadding="0" align="center"
           style="width: 600px; margin: 0px auto; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:微软雅黑, 黑体; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;">
        <tbody>
        <tr>
            <th valign="middle"
                style="height: 25px; line-height: 25px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: #42a3d3; background-color: #49bcff; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;">
                <span font="微软雅黑" size="5" style="color: rgb(255, 255, 255); ">注册成功! （go-vue）</span>
            </th>
        </tr>
        <tr>
            <td>
                <div style="padding:25px 35px 40px; background-color:#fff;">
                    <h2 style="margin: 5px 0px; ">
                       <span color="#333333" style="line-height: 20px; ">
                           <span style="line-height: 22px; " size="4">
                               尊敬的{{.Name}}，您好,</span>
                       </span>
                    </h2>
                    <p>感谢您使用go-vue服务。<br/>
                    <p>请点击以下链接进行邮箱验证，以便开始使用您的go-vue服务账户：<br></p>
                    <a href="{{.ActiveUrl}}" target="_blank"
                       style="display: inline-block;color:#fff;line-height: 40px;background-color: #1989fa;border-radius: 5px;text-align: center;text-decoration: none;font-size: 14px;padding: 1px 30px;"
                       rel="noopener">
                        马上验证邮箱
                    </a>
                    </p>

                    <div style="width:700px;margin:0 auto;">
                        <div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
                            <p>如果您无法点击以上链接，请复制以下网址到浏览器里直接打开：
                            <p>{{.ActiveUrl}}</p>
                            <p>如果您并未申请go-vue服务账户，可能是其他用户误输入了您的邮箱地址。请忽略此邮件，或者<a
                                        href="mailto:go_vue@163.com?subject=go-vue意见反馈">联系我们</a>。</p>

                            <p>©go-vue  <span>&nbsp;&nbsp; {{.Date}}</span></p>
                        </div>
                    </div>
                </div>
            </td>
        </tr>
        </tbody>
    </table>
</div>
`

//var check = func(err error){
//	if err != nil {
//		util.Log.Fatal(err)
//	}
//}

func SendActiveEmail(Name, Token, To string) error {
	t, err := template.New("activeEmailTpl").Parse(tpl)
	if err != nil {
		util.Log.Error(err)
		return err
	}
	data := struct {
		Name      string
		ActiveUrl string
		Date      string
	}{
		Name:      Name,
		ActiveUrl: "http://127.0.0.1:8080/#/email/confirm?token=" + Token,
		Date:      time.Now().Format("2006-01-02"),
	}
	buffer := &bytes.Buffer{}
	err = t.Execute(buffer, data)
	if err != nil {
		util.Log.Error(err)
		return err
	}
	e := &email.Email{
		To:      []string{To},
		From:    "go-vue <go_vue@163.com>",
		Subject: "go-vue邮件中心",
		HTML:    buffer.Bytes(),
		Headers: textproto.MIMEHeader{},
	}
	err = e.Send("smtp.163.com:25", smtp.PlainAuth("", "go_vue@163.com", os.Getenv("MAIL_PW"), "smtp.163.com"))
	if err != nil {
		util.Log.Error(err)
		return err
	}
	return nil
}

func generateToken(user model.User) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":   user.ID,
		"name": user.UserName,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	fmt.Println(tokenString, err)
	return tokenString, err
}
