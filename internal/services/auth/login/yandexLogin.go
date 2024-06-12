package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/gin-gonic/gin"
)

func LoginYandex(c *gin.Context) {
	// Получаем токен из хэша URL
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not provided"})
		return
	}

	userInfo, err := getYandexUserInfo(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func getYandexUserInfo(token string) (*models.YandexUserInfo, error) {
	// URL API для получения информации о пользователе
	apiURL := "https://login.yandex.ru/info?format=json"

	// Создаем запрос
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Добавляем заголовок авторизации с токеном
	req.Header.Add("Authorization", "OAuth "+token)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Проверяем статус код ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Декодируем ответ
	var userInfo models.YandexUserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
