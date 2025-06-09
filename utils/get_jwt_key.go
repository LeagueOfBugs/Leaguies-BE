package utils

import (
    "log"
    "os"
)

func GetJWTSecret() []byte {
    key := os.Getenv("JWT_SECRET")
    if key == "" {
        log.Fatal("JWT_SECRET environment variable not set")
    }
    return []byte(key)
}
