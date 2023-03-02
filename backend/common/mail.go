package common

import (
	"atm/global"
	"atm/model"
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

func SendVerifyCode(email, content string) error {
	smtp := global.Conf.MailConf.SmtpServer
	secret := global.Conf.MailConf.AuthCode
	sender := global.Conf.MailConf.Sender

	fmt.Println("mail--", secret, sender, content, email)
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", email)
	m.SetHeader("Cc", email)
	m.SetHeader("Bcc", email)
	m.SetHeader("Subject", "AcademicTeamManager VerifyCode")
	m.SetBody("text/html", content)
	d := gomail.NewDialer(smtp, 465, sender, secret)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("qq mail send error : %s", err)
		return err
	}
	return nil
}

func SendMail(mp model.MailParam) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mp.Sender)
	m.SetHeader("To", mp.Receiver)
	m.SetHeader("Subject", mp.Subject)
	m.SetBody("text/html", mp.Content)
	if mp.Attachment != "" {
		m.Attach(mp.Attachment)
	}
	d := gomail.NewDialer(mp.Smtp, mp.Port, mp.Sender, mp.AuthCode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Printf("send mail to student error : %s", err)
		return err
	}
	return nil
}

func DialMail(smtp string, port int, sender, authCode string) error {
	d := gomail.NewDialer(smtp, port, sender, authCode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	_, err := d.Dial()
	return err
}
