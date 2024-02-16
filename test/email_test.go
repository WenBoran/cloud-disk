package test

import (
	"clou-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <wbr@bupt.cn>"
	e.To = []string{"254750850@qq.com"}
	e.Subject = "验证测试"
	e.HTML = []byte("您的验证码为：<h1>Fancy HTML is supported, too!</h1>")
	err := e.SendWithTLS("smtp.exmail.qq.com:465", smtp.PlainAuth("", "wbr@bupt.cn", define.MailPassword, "smtp.exmail.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.exmail.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
