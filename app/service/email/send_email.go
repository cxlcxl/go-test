package email

import (
	"fmt"
	JordanEmail "github.com/jordan-wright/email"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"net/smtp"
	"strings"
)

type Email struct {
	Host     string
	Port     string
	Subject  string
	From     string
	Username string
	Password string
}

// Factory 创建邮件
func Factory() *Email {
	return &Email{
		Host:     variable.ConfigYml.GetString("Email.Host"),
		Port:     variable.ConfigYml.GetString("Email.Port"),
		From:     variable.ConfigYml.GetString("Email.From"),
		Subject:  consts.EmailGlobalSubject,
		Username: variable.ConfigYml.GetString("Email.Username"),
		Password: variable.ConfigYml.GetString("Email.Password"),
	}
}

// Send 发送邮件
// to 邮件接受者，多个接受者使用“,”分割
func (e *Email) Send(to, content string, arguments ...interface{}) error {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)
	je := JordanEmail.NewEmail()
	je.Text = e.combinationContent(content, arguments...)
	je.Subject = e.Subject
	je.To = strings.Split(to, ",")
	je.From = e.From

	return je.Send(e.Host+":"+e.Port, auth)
}

// combinationContent 组合内容
func (e *Email) combinationContent(content string, arguments ...interface{}) []byte {
	return []byte(fmt.Sprintf(content, arguments...))
}
