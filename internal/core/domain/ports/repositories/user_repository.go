// internal/core/domain/ports/repositories/user_repository.go

// คือการสร้าง interface สำหรับการจัดการข้อมูลผู้ใช้

package repositories

import (
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/entities"
)

// UserRepository กำหนดเมธอดที่จำเป็นสำหรับการจัดการข้อมูลผู้ใช้
type UserRepository interface {
	// สร้างผู้ใช้ใหม่
	Create(user *entities.User) error
	// แสดงรายละเอียดผู้ใช้ตามอีเมล
	GetByEmail(email string) (*entities.User, error)
	// แสดงรายละเอียดผู้ใช้ตาม ID
	GetByID(id uint) (*entities.User, error)
	// อัปเดตข้อมูลผู้ใช้
	Update(user *entities.User) error
	// ลบผู้ใช้ตาม ID
	Delete(id uint) error
	// ดึงรายชื่อผู้ใช้ทั้งหมด
	GetAll() ([]entities.User, error)
}
