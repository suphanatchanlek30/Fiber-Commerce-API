// internal/core/domain/ports/services/auth_service.go

package services

import (
	"errors"

	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/entities"
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/ports/repositories"
	"github.com/suphanatchanlek30/fiber-commerce-api/pkg/utils"
)

// AuthServiceImpl โครงสร้างที่ใช้ในการดำเนินการเกี่ยวกับการพิสูจน์ตัวตน
type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

// สร้างฟังก์ชัน NewAuthService สำหรับสร้าง NewAuthService ใหม่
func NewAuthService(userRepo repositories.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

// Register ฟังก์ชันสำหรับลงทะเบียนผู้ใช้ใหม่
func (s *AuthServiceImpl) Register(req entities.RegisterRequest) (*entities.User, error) {
	// เช็คถ้ามี user ที่มีอีเมลเดียวกันอยู่แล้ว
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash รหัสผ่าน
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// สร้าง user ใหม่
	user := &entities.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      entities.RoleUser, // กำหนดบทบาทเป็น User
		IsActive:  true,              // กำหนดสถานะเป็น Active
	}

	// เรียกใช้ userRepo เพื่อบันทึก user ใหม่
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// คืนค่าผู้ใช้ที่ถูกสร้างขึ้น
	return user, nil
}

// AdminRegister ฟังก์ชันสำหรับลงทะเบียนผู้ใช้ใหม่เป็น admin
func (s *AuthServiceImpl) AdminRegister(req entities.AdminRegisterRequest) (*entities.User, error) {
	// เช็คถ้ามี user ที่มีอีเมลเดียวกันอยู่แล้ว
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash รหัสผ่าน
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// สร้าง user ใหม่ด้วย role ที่กำหนด
	user := &entities.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role, // ใช้ role ที่ส่งมา
		IsActive:  true,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login ฟังก์ชันสำหรับเข้าสู่ระบบ
func (s *AuthServiceImpl) Login(req entities.LoginRequest) (*entities.LoginResponse, error) {

	user, err := s.userRepo.GetByEmail(req.Email)

	// Check if user exists
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// return login response
	return &entities.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthServiceImpl) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *AuthServiceImpl) UpdateUser(user *entities.User) error {
	return s.userRepo.Update(user)
}
