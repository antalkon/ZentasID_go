package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"-"`
	SSLMode  string `yaml:"ssl_mode"`
}

func GenerateDBConnectionString() (string, error) {
	dbConfig, err := DBConfig()
	if err != nil {
		return "", err
	}

	connectionString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.User, dbConfig.Password, dbConfig.SSLMode)
	fmt.Println(connectionString)
	return connectionString, nil
}

func DBConfig() (*DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ошибка загрузки файла .env: %s", err)
	}

	viper.SetConfigName("dataBase")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs") // предполагается, что файл конфигурации находится в том же каталоге

	// Чтение файла конфигурации
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ошибка чтения файла конфигурации: %s", err)
	}

	// Установка префикса переменной окружения для Viper
	viper.SetEnvPrefix("DB")

	// Автоматическое связывание переменных окружения
	viper.AutomaticEnv()

	// Чтение переменных окружения
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatalf("пароль не найден в переменных окружения")
	}

	var dbConfig DB
	if err := viper.Unmarshal(&dbConfig); err != nil {
		log.Fatalf("ошибка разбора конфигурации: %s", err)
	}

	dbConfig.Password = password
	fmt.Println(dbConfig.DbName)
	return &dbConfig, nil
}
