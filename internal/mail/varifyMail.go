package z_mail

import (
	"fmt"
	// Замените на путь к вашему пакету

	mail_config "github.com/antalkon/ZentasID_go/pkg/mail"
	"gopkg.in/gomail.v2"
)

func VerifyMail(token, email string) error {
	// Загрузите конфигурацию
	mailConfig, err := mail_config.ReadConfigMAIL()
	if err != nil {
		return fmt.Errorf("failed to read mail config: %s", err)
	}

	smtpServer := mailConfig.Host
	smtpPort := mailConfig.Port
	smtpUser := "verify@zentas.ru" // Замените на ваш email
	smtpPassword := ""             // Замените на ваш пароль

	// Создайте новое письмо
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", email) // Замените на email получателя
	m.SetHeader("Subject", "Greetings!")
	verificationURL := fmt.Sprintf("https://localhost:8080/auth/api/verify/%s", token)
	messageBody := fmt.Sprintf(`
        <html>
        <body>
            <p>Hello,</p>
            <p>Please verify your email by clicking the following link:</p>
            <a href="%s">%s</a>
            <p>Thank you!</p>
        </body>
        </html>
    `, verificationURL, verificationURL)
	m.SetBody("text/html", messageBody)
	// Создайте новый SMTP диалект
	d := gomail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)

	// Отправьте письмо
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("Error sending mail: %s", err)
	} else {
		return nil
	}
}
