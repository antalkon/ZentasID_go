package reg

import (
	"log"
	"net/http"
	"time"

	regapi_pg "github.com/antalkon/ZentasID_go/internal/database/postgres/regApi_pg"
	z_mail "github.com/antalkon/ZentasID_go/internal/mail"
	"github.com/antalkon/ZentasID_go/internal/models"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/antalkon/ZentasID_go/pkg/UUID"
	randnum "github.com/antalkon/ZentasID_go/pkg/randNum"
	"github.com/gin-gonic/gin"
)

func RegistrationApi(c *gin.Context) {
	// init variables
	var regUser models.RegUser
	var (
		displayID     = randnum.GenerateRandomNumber()
		currentTime   = time.Now()
		formattedTime = currentTime.Format("02.01.2006")
	)

	// Generation UUID
	userID, err := UUID.GenerateUserID() // ID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userNick, err := UUID.GenerateUserNick() // NickName
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Разбите JSON по структуре
	if err := c.ShouldBindJSON(&regUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Присываивание доп.Значений к струтуре
	regUser.UserID = userID
	regUser.DisplayID = displayID
	regUser.RegDate = formattedTime
	regUser.NickName = userNick

	// Валидация структуры
	if err := models.ValidateUser(regUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validate error: ": err.Error(), "error": "Ошибка валидации: Не все поля заполнены или содержат неверные значения."})
		return
	}

	// Сохранения пользователя в базу данных
	joinDB, err := regapi_pg.DbRegistr(regUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "warn": "Такой пользователь уже существует!"})
		return
	}
	_ = joinDB

	// Генерация токена для подтверждения Email адреса
	verifyToken, err := jwt.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Запись токена в базу данных
	err = regapi_pg.SaveVerifyToken(userID, verifyToken, formattedTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Горутина для отправки письма подтверждения регистрации
	go func() {
		err := z_mail.VerifyMail(verifyToken, regUser.UserEmail, regUser.UserName)
		if err != nil {
			log.Println("Error verifying email:", err)
		}
	}()

	// Если успешно - возвращаем.
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
