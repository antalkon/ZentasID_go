package z_mail

import (
	"fmt"
	// Замените на путь к вашему пакету

	mail_config "github.com/antalkon/ZentasID_go/pkg/mail"
	"gopkg.in/gomail.v2"
)

func LoginCodeMail(email string, code int) error {
	// Загрузите конфигурацию
	mailConfig, err := mail_config.ReadConfigMAIL()
	if err != nil {
		return fmt.Errorf("failed to read mail config: %s", err)
	}

	smtpServer := mailConfig.Host
	smtpPort := mailConfig.Port
	smtpUser := "id@zentas.ru"             // Замените на ваш email
	smtpPassword := "94M77miuriiyDPfuaE7a" // Замените на ваш пароль

	// Создайте новое письмо
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(smtpUser, "Зентас ID"))
	m.SetHeader("To", email) // Замените на email получателя
	m.SetHeader("Subject", "Код для входа в аккаунт")
	messageBody := fmt.Sprintf(`
        <html>
        <body>
            <h1>Здраствуйте!</p>
            <p>Ваш код для входа в аккаунт:</p>
            <a href="#">%d</a>
            <h2>Если это не вы - перейдите по ссылке чтобы сбросить все сессии...</p>
        </body>
        </html>
    `, code)
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
