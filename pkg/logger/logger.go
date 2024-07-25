package logger

import (
	"log/slog" // Импортируем пакет slog для логирования
	"os"       // Импортируем пакет os для работы с файловой системой
)

// Определяем константы для различных окружений
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// SetupLogger инициализирует и возвращает новый экземпляр логгера в зависимости от окружения.
// env - параметр, указывающий текущее окружение (local, dev, prod).
func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger // Объявляем переменную для логгера
	var logFile *os.File // Объявляем переменную для файла логов
	var err error        // Объявляем переменную для хранения ошибок

	// Открываем файл log.json в режиме добавления или создаем его, если он не существует.
	// Файл будет доступен для записи.
	logFile, err = os.OpenFile("cmd/log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // В случае ошибки открытия файла логов, завершаем выполнение программы с паникой.
	}

	// Выбираем обработчик логов и уровень логирования в зависимости от окружения.
	switch env {
	case envLocal, envDev, envProd:
		log = slog.New(
			// Используем JSON-обработчик для записи логов в формате JSON.
			// Файл logFile будет использоваться для записи логов.
			// Уровень логирования определяется функцией getLogLevel.
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: getLogLevel(env)}),
		)
	default:
		log = slog.New(
			// Если окружение не распознано, используем JSON-обработчик с уровнем логирования INFO.
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log // Возвращаем настроенный логгер.
}

// getLogLevel возвращает уровень логирования в зависимости от окружения.
// env - параметр, указывающий текущее окружение (local, dev, prod).
func getLogLevel(env string) slog.Level {
	switch env {
	case envLocal, envDev:
		return slog.LevelDebug // Для локального и девелоперского окружений используем уровень DEBUG.
	case envProd:
		return slog.LevelInfo // Для продакшен окружения используем уровень INFO.
	default:
		return slog.LevelInfo // Для неизвестных окружений используем уровень INFO.
	}
}
