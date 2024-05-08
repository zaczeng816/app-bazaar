package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var(
	APP_INDEX  = "app" 
    USER_INDEX = "user"
    SERVER_PORT = ":8080"
    STRIPE_API_KEY  string
    STRIPE_CHECKOUT_SESSION_URL string
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
    SERVER_PORT = getEnv("SERVER_PORT", ":8080")
    STRIPE_API_KEY = getEnv("STRIPE_API_KEY", "sk_test_4eC39HqLyjWDarjtT1zdp7dc")
    STRIPE_CHECKOUT_SESSION_URL = getEnv("STRIPE_CHECKOUT_SESSION_URL", "http://localhost:4242")
    return nil
}