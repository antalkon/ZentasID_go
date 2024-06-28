package userdata

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/dataApi_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func UserEditAvatar(c *gin.Context) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}
	decoder, err := jwt.DecodeAccessToken(cookie)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file", "details": err.Error()})
		return
	}
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String()[:12] + ext
	filePath := filepath.Join("storage", "users", "avatar", newFileName)

	// Сохраняем файл
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
		return
	}

	// Открываем сохраненный файл
	imgIn, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file", "details": err.Error()})
		return
	}
	defer imgIn.Close()

	var img image.Image
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(imgIn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JPEG image", "details": err.Error()})
			return
		}
	case ".png":
		img, err = png.Decode(imgIn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode PNG image", "details": err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file format"})
		return
	}

	// Изменяем размер и обрезаем изображение
	img = resize.Resize(600, 0, img, resize.Bicubic)
	croppedImg := cropImage(img, 500, 500)

	// Сохраняем обрезанное изображение
	imgOut, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file", "details": err.Error()})
		return
	}
	defer imgOut.Close()

	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(imgOut, croppedImg, nil)
	case ".png":
		err = png.Encode(imgOut, croppedImg)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image", "details": err.Error()})
		return
	}

	// Сохранение информации об аватаре в базе данных
	err = dataApi_pg.SaveDbAvatar(decoder.UserID, newFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar to database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Avatar has been set"})
}

// Функция для обрезки изображения
func cropImage(img image.Image, width, height int) image.Image {
	originalBounds := img.Bounds()
	originalWidth := originalBounds.Dx()
	originalHeight := originalBounds.Dy()
	startX := (originalWidth - width) / 2
	startY := (originalHeight - height) / 2
	croppedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			croppedImg.Set(x, y, img.At(startX+x, startY+y))
		}
	}
	return croppedImg
}
