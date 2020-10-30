package mailSender

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/smtp"
)

func genUUID() string {
	return uuid.New().String()
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendConfirmationEmail(email string) string {
	ui := genUUID()
	from := "morlandemailconfirmation@gmail.com"
	password := "Dirtysexy80085~"
	to := []string{email}
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	message := []byte("Thank you for registration on CrackTheBet. \nPlease verify your email: http://127.0.0.1:5555/confirm-email/?token=" + ui)
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		log.Println(err)
		return ""
	} else {
		return ui
	}
}

func SendPasswordRecoveryEmail(email string) string {
	ui := genUUID()
	from := "morlandemailconfirmation@gmail.com"
	password := "Dirtysexy80085~"
	to := []string{email}
	fmt.Println(to)
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	message := []byte("You password recovery link is down below. Thank you for staying with us. Link: http://127.0.0.1:5555/recovery?token=" + ui)
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		log.Println(err)
		return ""
	} else {
		return ui
	}
}
