package noticer

import (
	"gopkg.in/gomail.v2"
)

type EmailPackage struct {
	Destination []string  `json:"destination"`
	Subject     string    `json:"subject"`
	Body        EmailBody `json:"body"`
}

type EmailBody struct {
	Sender   string `json:"sender"`
	SendTime string `json:"send_time"`
	Message  string `json:"message"`
}

func SendEmail(email *EmailPackage) error {
	// 设置邮箱主体
	mailConn := map[string]string{
		"user": "develop_pqbs@163.com",
		"pass": "HYZXNSQZRZTQMCUP",
		"host": "smtp.163.com",
	}

	m := gomail.NewMessage(
		//就选择utf-8格式保存
		gomail.SetEncoding(gomail.Base64),
	)
	m.SetHeader("From", m.FormatAddress(mailConn["user"], email.Body.Sender))
	m.SetHeader("To", email.Destination...)
	m.SetHeader("Subject", email.Subject)
	message := email.Body.Message + "\ntime: " + email.Body.SendTime
	m.SetBody("text/html", message)

	// 163 port: 25
	d := gomail.NewDialer(mailConn["host"], 25, mailConn["user"], mailConn["pass"]) // 设置邮件正文
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}
