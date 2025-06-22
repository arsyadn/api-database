package config

import (
	"os"
	"time"
)

// mengambil JWT secret key dari environment variable
func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

// mengambil durasi kedaluwarsa JWT dari environment variable
func GetJWTExpirationDuration() time.Duration {
	duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))

	// jika terjadi kesalahan dalam parsing, kembalikan nilai default 24 jam
	// ini bisa diubah sesuai kebutuhan
	if err != nil {
		return time.Hour * 24
	}

	return duration
}