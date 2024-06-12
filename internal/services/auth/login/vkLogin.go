package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/gin-gonic/gin"
)

func LoginVK(c *gin.Context) {
	payloadStr := c.Query("payload")

	var payload models.PayloadVK
	err := json.Unmarshal([]byte(payloadStr), &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// Осуществляем обмен silent token на access token
	accessTokenResponse, err := exchangeSilentAuthToken(payload.Token, payload.UUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange silent token for access token", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accessTokenResponse)
}

func exchangeSilentAuthToken(silentToken, uuid string) (*models.AccessTokenResponseVK, error) {
	// Параметры запроса
	data := url.Values{}
	data.Set("v", "5.131")
	data.Set("token", silentToken)
	data.Set("access_token", "0400befc0400befc0400befcd7071823e1004000400befc62693e9041c06bf3ccb513a5")
	data.Set("uuid", uuid)

	// Выполняем POST-запрос к API ВКонтакте
	resp, err := http.PostForm("https://api.vk.com/method/auth.exchangeSilentAuthToken", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Проверяем код статуса ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Декодируем ответ
	var accessTokenResponse models.AccessTokenResponseVK
	err = json.NewDecoder(resp.Body).Decode(&accessTokenResponse)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponse, nil
}
