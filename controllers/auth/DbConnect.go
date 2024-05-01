package auth

import (
	"fmt"
	"github.com/spf13/viper"
)

func getDBConnectionString() string {
	// Получаем данные для подключения к базе данных из YAML-файла
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")

	// Формируем строку подключения к базе данных
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	return connStr
}
