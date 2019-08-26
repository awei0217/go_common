package basic

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

// 邮件配置信息
const (
	ADDR     = ""
	USER     = "sunpengwei" //发送邮件的邮箱
	PASSWORD = ""           //发送邮件邮箱的密码
	FROM     = ""
)

/**
发送邮件实体， to  发给谁， subject 邮件主题 message 邮件内容
*/
type Email struct {
	To      string
	Subject string
	Message string
}

type LoginAuth struct {
	username, password string
}

/**
验证

*/
func NewLoginAuth(username, password string) smtp.Auth {

	return &LoginAuth{username, password}
}

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		fmt.Println(string(fromServer))
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

func send(addr, subject string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	host, _, _ := net.SplitHostPort(addr)
	if err != nil {
		log.Println("call dial")
		return err
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			fmt.Println("call start tls")
			return err
		}
	}

	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
				fmt.Println("check auth with err:", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}

	header := make(map[string]string)
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(msg)
	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

/**
发送邮件
*/
func SendMail(email *Email) {
	auth := NewLoginAuth(USER, PASSWORD)
	//auth := smtp.PlainAuth("", USER, PASSWORD, "mail.gomeplus.com")
	fmt.Println(auth)
	err := send(ADDR, email.Subject, auth, FROM, strings.Split(email.To, ","), []byte(email.Message))
	if err != nil {
		log.Println("发送邮件失败", err)
	} else {
		log.Println("发送邮件成功")
	}

}
