package helper

import (
	"clou-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
)

func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "Get <wbr@bupt.cn>"
	e.To = []string{mail}
	e.Subject = "验证测试"
	e.HTML = []byte("您的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.exmail.qq.com:465", smtp.PlainAuth("", "wbr@bupt.cn", define.MailPassword, "smtp.exmail.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.exmail.qq.com"})
	if err != nil {
		return err
	}
	return nil

}

func RandCode() string {
	s := "1234567890"
	code := ""
	//rand.Seed()
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
