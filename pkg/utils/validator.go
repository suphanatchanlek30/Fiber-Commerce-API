// pkg/utils/validator.go

package utils

import (
	"github.com/go-playground/validator/v10"
)

// สร้างตัวแปร validate สำหรับการตรวจสอบความถูกต้องของข้อมูล เรียกใช้ validator.New()
var validate = validator.New()

// สร้างฟังก์ชัน Validate ที่ใช้สำหรับตรวจสอบความถูกต้องของข้อมูล
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
