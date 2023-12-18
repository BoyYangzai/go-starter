package handler

import (
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
)

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth(password, user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func generateVerificationCode(length int) string {
	code := ""
	for i := 0; i < length; i++ {
		code += strconv.Itoa(rand.Intn(10)) // 生成随机数字（0-9）
	}
	return code
}
