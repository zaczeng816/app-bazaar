package constants

import "os"

var (
	APP_INDEX   =  "app" 
    USER_INDEX  =  "user"
	SERVER_PORT = GetEnv("SERVER_PORT", ":8080")
)

func GetEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}