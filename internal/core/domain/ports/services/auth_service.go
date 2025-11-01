// internal/core/domain/ports/services/auth_service.go

package services

import (
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/entities"
)

type AuthService interface {
	// สร้างผู้ใช้ใหม่
	Register(req entities.RegisterRequest) (*entities.User, error)
	// สร้างผู้ดูแลระบบใหม่
	AdminRegister(req entities.AdminRegisterRequest) (*entities.User, error)
	// ผู้ใช้เข้าสู่ระบบ
	Login(req entities.LoginRequest) (*entities.LoginResponse, error)
	// ดึงข้อมูลผู้ใช้ตาม ID
	GetUserByID(id uint) (*entities.User, error)
	// อัปเดตข้อมูลผู้ใช้
	UpdateUser(user *entities.User) error
}
