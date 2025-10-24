// internal/config/config.go

package config

import (
	"os"

	"github.com/joho/godotenv"
)

// สร้างโครงสร้าง Config เพื่อเก็บค่าการตั้งค่าต่างๆ
type Config struct {
	AppEnv       string
	AppPort      string
	AppURL       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	JWTExpiresIn string
}

// LoadEnv โหลดตัวแปรสภาพแวดล้อมจากไฟล์ .env
func LoadEnv() (*Config, error) {
	return &Config{
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "8080"),
		AppURL:       getEnv("APP_URL", "http://localhost:8080"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASS", "your_password_here"),
		DBName:       getEnv("DB_NAME", "fiberecomapidb"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTSecret:    getEnv("JWT_SECRET", "fibernextcommerce_jwt_secret_key_2024"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24"),
	}, godotenv.Load()
}

// สร้างฟังก์ชันช่วยเหลือเพื่อดึงค่าจากตัวแปรสภาพแวดล้อมหรือใช้ค่าเริ่มต้น
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
