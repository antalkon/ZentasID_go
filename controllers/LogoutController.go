package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// LogoutHandler обрабатывает выход пользователя
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Удаляем cookie файлы у пользователя
	http.SetCookie(w, &http.Cookie{
		Name:   "refresh_token",
		MaxAge: -1,
		Path:   "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		MaxAge: -1,
		Path:   "/",
	})

	// Получаем refresh токен из cookie файла
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		fmt.Println("Ошибка при получении refresh токена из cookie файла:", err)
		return
	}

	// Удаляем строку в таблице refresh_tokens
	err = deleteRefreshToken(refreshToken.Value)
	if err != nil {
		fmt.Println("Ошибка при удалении refresh токена из базы данных:", err)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Выход выполнен успешно"))
}

// deleteRefreshToken удаляет строку с указанным refresh токеном из таблицы refresh_tokens
func deleteRefreshToken(refreshToken string) error {
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM refresh_tokens WHERE refresh_token = $1", refreshToken)
	if err != nil {
		return err
	}

	return nil
}
