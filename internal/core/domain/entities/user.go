// core/domain/entities/user.go

package entities

import (
	"time"
)

type Role string

// กำหนดค่าคงที่สำหรับบทบาทผู้ใช้
const (
	RoleAdmin     Role = "admin"
	RoleUser      Role = "user"
	RoleModerator Role = "moderator"
)

// สร้างโครงสร้าง User เพื่อเก็บข้อมูลผู้ใช้
type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      Role      `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// สร้างโครงสร้างสำหรับคำขอและการตอบกลับที่เกี่ยวข้องกับผู้ใช้
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterRequest โครงสร้างสำหรับคำขอลงทะเบียนผู้ใช้ใหม่
type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

// AdminRegisterRequest represents admin registration request payload
type AdminRegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      Role   `json:"role" validate:"required,oneof=admin user moderator"`
}

// LoginResponse โครงสร้างสำหรับการตอบกลับการเข้าสู่ระบบ
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
