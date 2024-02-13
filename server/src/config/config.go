package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetDbName() string {
	return get("DB_NAME")
}

func GetSecterKey() []byte {
	return []byte(get("SECRET_KEY"))
}

// Config func to get env value from key ---
func get(key string) string {
	// load .env file
	err := godotenv.Load("letsgo-blog.env")
	if err != nil {
		log.Error("Error loading .env file", err)
	}
	return os.Getenv(key)

}
