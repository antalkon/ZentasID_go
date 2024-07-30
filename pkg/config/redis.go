package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type REDIS struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string
}

func GenerateRedisConnectionString() (string, error) {
	redisConfig, err := RedisConfig()
	if err != nil {
		return "", err
	}

	connectionString := fmt.Sprintf("redis://:%s@%s:%s",
		redisConfig.Password, redisConfig.Host, redisConfig.Port)
	fmt.Println(connectionString)
	return connectionString, nil
}

func RedisConfig() (*REDIS, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ошибка загрузки файла .env: %s", err)
	}

	viper.SetConfigName("redis")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs") // предполагается, что файл конфигурации находится в папке configs

	// Чтение файла конфигурации
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ошибка чтения файла конфигурации: %s", err)
	}

	// Установка префикса переменной окружения для Viper
	viper.SetEnvPrefix("REDIS")

	// Автоматическое связывание переменных окружения
	viper.AutomaticEnv()

	// Чтение переменных окружения
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		log.Fatalf("пароль не найден в переменных окружения")
	}

	var redisConfig REDIS
	if err := viper.Unmarshal(&redisConfig); err != nil {
		log.Fatalf("ошибка разбора конфигурации: %s", err)
	}

	redisConfig.Password = password
	fmt.Println(password)
	return &redisConfig, nil
}
