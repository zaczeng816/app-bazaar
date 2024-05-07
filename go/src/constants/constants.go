package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_INDEX       string 
    USER_INDEX      string
	SERVER_PORT     string
    ES_URL          string
    ES_USERNAME     string
    ES_PASSWORD     string
    STRIPE_API_KEY  string
    STRIPE_CHECKOUT_SESSION_API string
)

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func LoadEnv() error {
    err := godotenv.Load()
    if err != nil {
        return err
    }
    APP_INDEX   = getEnv("APP_INDEX", "app")
    USER_INDEX  = getEnv("USER_INDEX", "user")
    ES_URL      = getEnv("ES_URL", "http://localhost:9200")
    ES_USERNAME = getEnv("ES_USERNAME", "admin")
    ES_PASSWORD = getEnv("ES_PASSWORD", "")
    SERVER_PORT = getEnv("SERVER_PORT", ":8080")
    STRIPE_API_KEY = getEnv("STRIPE_API_KEY", "sk_test_4eC39HqLyjWDarjtT1zdp7dc")
    STRIPE_CHECKOUT_SESSION_API = getEnv("STRIPE_CHECKOUT_SESSION_API", "http://localhost:4242")
    return nil
}