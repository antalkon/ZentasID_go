package z_mail

import (
	"fmt"
	// Замените на путь к вашему пакету

	mail_config "github.com/antalkon/ZentasID_go/pkg/mail"
	"gopkg.in/gomail.v2"
)

func VerifyMail(token, email, name string) error {
	// Загрузите конфигурацию
	mailConfig, err := mail_config.ReadConfigMAIL()
	if err != nil {
		return fmt.Errorf("failed to read mail config: %s", err)
	}

	smtpServer := mailConfig.Host
	smtpPort := mailConfig.Port
	smtpUser := "verify@zentas.ru"         // Замените на ваш email
	smtpPassword := "Yp60Y2sgtRg6twJisyay" // Замените на ваш пароль

	// Создайте новое письмо
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(smtpUser, "Зентас ID"))
	m.SetHeader("To", email) // Замените на email получателя
	m.SetHeader("Subject", "Подтверждение электронной почты")
	verificationURL := fmt.Sprintf("https://localhost:8080/auth/api/v1/verify/%s", token)
	messageBody := fmt.Sprintf(`
        <html>
        <body>
            <h1>Здраствуйте, %s!</p>
            <p>Пожалуйста подтвердите свой электронный адрес для активации аккаунта по ссылки:</p>
            <a href="%s">%s</a>
            <h2>Спасибо!</p>
        </body>
        </html>
    `, name, verificationURL, verificationURL)
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
