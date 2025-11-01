// internal/config/database.go

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/suphanatchanlek30/fiber-commerce-api/internal/adapters/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabase ตั้งค่าการเชื่อมต่อฐานข้อมูล
func SetupDatabase(config *Config) *gorm.DB {

	// สร้าง Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// ตรวจสอบว่าควรทำการ migrate หรือไม่
	if shoulRundMigratuion() {

		// runMigration จะทำการ migrate ถ้าเงื่อนไขเป็นจริง
		runMigrations(db)
	} else {
		// แสดง message ที่ชัดเจนขึ้นตามสาเหตุ
		autoMigrate := os.Getenv("AUTO_MIGRATE")
		appEnv := os.Getenv("APP_ENV")

		if autoMigrate == "false" {
			log.Printf("Skipping database migration (AUTO_MIGRATE=false)")
		} else if appEnv == "production" && autoMigrate != "true" {
			log.Printf("Skipping database migration (production environment, set AUTO_MIGRATE=true to enable)")
		} else {
			log.Printf("Skipping database migration (set AUTO_MIGRATE=true to enable)")
		}
	}

	return db

}

// สร้างฟังก์ชั่นตรวจสอบว่าควร migrate หรือไม่
func shoulRundMigratuion() bool {
	// ถ้ากำหนด AUTO_MIGRATE เป็น false ให้ไม่ migration เลย (ทุก environment)
	if os.Getenv("AUTO_MIGRATE") == "false" {
		return false
	}

	// ถ้ากำหนด AUTO_MIGRATE เป็น true ให้ migration เลย (ทุก environment)
	if os.Getenv("AUTO_MIGRATE") == "true" {
		return true
	}

	// ถ้าไม่ได้กำหนด AUTO_MIGRATE ให้ใช้ default ตาม environment
	// Development - migrate อัตโนมัติเสมอ
	// Production - ไม่ต้อง migrate อัตโนมัติ
	if os.Getenv("APP_ENV") == "development" {
		return true
	}

	// Production or อื่นๆ - ไม่ต้อง migrate อัตโนมัติ
	return false
}

// ฟังก์ชั่นสำหรับ migration
func runMigrations(db *gorm.DB) {
	// เริ่มต้น migration
	log.Println("Starting database migration...")

	// ตรวจสอบว่าต้องการ migrate หรือไม่
	err := db.AutoMigrate(&models.User{})

	// ตรวจสอบข้อผิดพลาด
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// สำเร็จ
	log.Println("Database migration completed successfully")
}

// ฟังก์ชันสำหรับ migrate แบบ manual (สำหรับ CLI)
// คือ ไม่ตรวจสอบเงื่อนไข AUTO_MIGRATE หรือ APP_ENV ใดๆ แล้ว migrate ทันที
func RunMigrationManual(config *Config) error {
	// ตั้งค่าการเชื่อมต่อฐานข้อมูล
	db := SetupDatabase(config)

	// เริ่มต้น migration
	log.Println("Running manual migration...")

	// ทำการ migrate
	err := db.AutoMigrate(&models.User{})

	// ตรวจสอบข้อผิดพลาด
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	// สำเร็จ
	log.Println("Manual migration completed successfully")

	// คืนค่า nil เมื่อสำเร็จ
	return nil
}
