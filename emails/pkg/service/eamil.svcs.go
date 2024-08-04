package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"common/grpcs/email/pb"

	"gopkg.in/gomail.v2"
)

type EmailSvcs struct {
	pb.UnimplementedEmailServiceServer
}

func NewEmailSvcs() *EmailSvcs {
	return &EmailSvcs{}
}

func (e *EmailSvcs) SendEmail(ctx context.Context, req *pb.SendEmailReq) (*pb.SendEmailRes, error) {

	fmt.Println("receive email form other svcs")

	email := req.GetEmail()

	host := os.Getenv("SMTPHOST")
	port, err := strconv.Atoi(os.Getenv("SMTPPORT"))
	if err != nil {
		return nil, err
	}
	user := os.Getenv("SMTPUSER")
	password := os.Getenv("SMTPPASS")

	fmt.Println(host, user, password)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s<%s>", os.Getenv("SMTPFROMNAME"), os.Getenv("SMTPFROMEMAIL")))
	m.SetHeader("To", email.To)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Message)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(host, port, user, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pb.SendEmailRes{
		Success: true,
	}, nil
}
