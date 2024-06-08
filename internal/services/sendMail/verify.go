package sendmail

import (
	"fmt"

	mail_config "github.com/antalkon/ZentasID_go/pkg/mail"
	"gopkg.in/gomail.v2"
)

func SendVerify(email, token string) error {
	cfg, err := mail_config.ReadConfigMAIL()
	if err != nil {
		return fmt.Errorf("failed to read mail config: %v", err)
	}

	smtpHost := cfg.Host
	smtpPort := cfg.Port
	smtpUsername := ""
	smtpPassword := ""

	message := gomail.NewMessage()
	message.SetHeader("From", "your_email@example.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Your Access Token")
	message.SetBody("text/html", fmt.Sprintf("To complete your registration, click the following link: <a href=\"http://example.com/verify?token=%s\">Link</a>", token))

	// Настройка SMTP клиента
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// Отправка письма
	if err := d.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
