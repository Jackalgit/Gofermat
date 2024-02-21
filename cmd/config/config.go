package config

import (
	"flag"
	"github.com/Jackalgit/Gofermat/internal/models"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

var Config struct {
	ServerPort    string
	LogLevel      string
	DatabaseDSN   string
	AccrualSystem string
	SecretKey     string
}

func ConfigServerPort() {
	flag.StringVar(&Config.ServerPort, "a", "localhost:8080", "Addres local port")

	if envRunServerAddr := os.Getenv("RUN_ADDRESS"); envRunServerAddr != "" {
		Config.ServerPort = envRunServerAddr

	}

}

func ConfigLogger() {
	flag.StringVar(&Config.LogLevel, "l", "info", "log level")

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		Config.LogLevel = envLogLevel
	}

}

func ConfigDatabaseDSN() {
	//ps := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
	//	`localhost`, `5432`, `gofermat`)

	flag.StringVar(&Config.DatabaseDSN, "d", "", "Database source name")

	if envDatabaseDSN := os.Getenv("DATABASE_URI"); envDatabaseDSN != "" {
		Config.DatabaseDSN = envDatabaseDSN
	}

}

func ConfigAccrualSystem() {

	flag.StringVar(&Config.AccrualSystem, "r", "http://localhost:8081", "System Accrual")

	if envDatabaseDSN := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envDatabaseDSN != "" {
		Config.DatabaseDSN = envDatabaseDSN
	}

}

func ConfigSecretKey() {

	secret := models.Secret{}

	for _, fileName := range []string{".env"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Println("[SECRET_KEY]: ", err)
		}
	}

	if err := envconfig.Process("", &secret); err != nil {
		log.Println("[SECRET_KEY]: ", err)
	}
	Config.SecretKey = secret.SecretKey

}
