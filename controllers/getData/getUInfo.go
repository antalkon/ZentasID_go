package getData

import (
	"TestWebServer/controllers/getData/db"
	"TestWebServer/controllers/getData/functions"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserinfoHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := functions.GetAccessToken(w, r)
	if accessToken != "" {
		data := reqDbDat(accessToken)
		if data != nil {
			// Если данные найдены, отправляем их клиенту
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data) // Отправка данных в формате JSON клиенту
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Пользователь не найден"))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка при получении доступа"})

	}
}

func reqDbDat(id string) []functions.Users {
	db, err := sql.Open("postgres", db.GetDBConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id, email, phone_number, name, surname, birth_date, verify, two_fa FROM users WHERE id = '%s'", id)

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []functions.Users
	fmt.Println(rows)
	// Перебор результатов запроса
	for rows.Next() {
		var user functions.Users
		err := rows.Scan(&user.Id, &user.Email, &user.Phone_number, &user.Name, &user.Surname, &user.Birth_date, &user.Verify, &user.Two_fa) // Исправлено на все поля структуры Users
		if err != nil {
			panic(err)
		}
		// Добавление пользователя в срез
		users = append(users, user)
	}
	// Проверка на наличие ошибок после перебора результатов запроса
	if err := rows.Err(); err != nil {
		panic(err)
	}

	if len(users) == 0 {
		return nil // Если пользователь не найден, возвращаем nil
	}

	return users
}
