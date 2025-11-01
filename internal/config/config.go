// internal/config/config.go

package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// สร้างโครงสร้าง Config เพื่อเก็บค่าการตั้งค่าต่างๆ
type Config struct {
	AppEnv         string
	AppPort        string
	AppURL         string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	JWTSecret      string
	JWTExpiresIn   string
	AdminEmail     string
	AdminPassword  string
	AdminFirstName string
	AdminLastName  string
}

// LoadEnv โหลดตัวแปรสภาพแวดล้อมจากไฟล์ .env
func LoadConfig() (*Config, error) {

	// โหลดไฟล์ .env ถ้ามี
	err := godotenv.Load()

	// ตรวจสอบข้อผิดพลาดในการโหลดไฟล์ .env
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	config := &Config{
		// ค่าที่ปลอดภัยสำหรับ default สามารถเปลี่ยนแปลงได้ตามความต้องการ
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "3000"),
		AppURL:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24"),

		// ค่าที่ไม่ปลอดถัย default ต้องกำหนดใน .env
		DBPassword:     getEnv("DB_PASS", ""),
		DBName:         getEnv("DB_NAME", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		AdminEmail:     getEnv("ADMIN_EMAIL", ""),
		AdminPassword:  getEnv("ADMIN_PASSWORD", ""),
		AdminFirstName: getEnv("ADMIN_FIRST_NAME", ""),
		AdminLastName:  getEnv("ADMIN_LAST_NAME", ""),
	}

	// ตรวจสอบค่าที่จำเป็นต้องมี
	if err := validateConfig(config); err != nil {
		log.Printf("Configuration validation error: %v\n", err)
		return nil, err
	}

	// เมื่อไม่มีข้อผิดพลาด คืนค่า config
	return config, nil
}

// ฟังก์ชั่นตรวจสอบค่าจำเป็นต้องมี สำหรับ production
func validateConfig(config *Config) error {
	// ตรวจสอบว่าค่าที่จำเป็นถูกตั้งค่าในสภาพแวดล้อม production
	if config.AppEnv == "production" {
		if config.DBPassword == "" {
			return fmt.Errorf("DB_PASS is required in production")
		}
		if config.DBName == "" {
			return fmt.Errorf("DB_NAME is required in production")
		}
		if len(config.JWTSecret) < 32 {
			return fmt.Errorf("JWT_SECRET must be at least 32 characters long in production")
		}
		if config.DBSSLMode == "disable" {
			log.Println("Warning: DB_SSL is set to disable in production")
		}
		if config.AdminEmail == "" {
			return fmt.Errorf("ADMIN_EMAIL is required in production")
		}
		if config.AdminPassword == "" {
			return fmt.Errorf("ADMIN_PASSWORD is required in production")
		}
		if config.AdminFirstName == "" {
			return fmt.Errorf("ADMIN_FIRST_NAME is required in production")
		}
		if config.AdminLastName == "" {
			return fmt.Errorf("ADMIN_LAST_NAME is required in production")
		}
	}

	// ตรวจสอบรูปแบบ email (เฉพาะเมื่อมีค่า)
	if config.AdminEmail != "" && !isValidEmail(config.AdminEmail) {
		return errors.New("ADMIN_EMAIL format is invalid")
	}

	// ตรวจสอบค่าพื้นฐานที่ควรมีเสมอ
	if config.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	// ถ้าผ่านการตรวจสอบทั้งหมด
	return nil
}

// สร้างฟังก์ชันช่วยเหลือเพื่อดึงค่าจากตัวแปรสภาพแวดล้อมหรือใช้ค่าเริ่มต้น
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ฟังก์ชั่น check email ว่ามีรูปแบบถูกต้องหรือไม่
func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	// ตรวจสอบพื้นฐาน - ต้องมี @ และ . และไม่เริ่มหรือจบด้วย @
	return len(email) > 0 && len(email) <= 254 &&
		strings.Contains(email, "@") &&
		strings.Contains(email, ".") &&
		!strings.HasPrefix(email, "@") &&
		!strings.HasSuffix(email, "@")
}
