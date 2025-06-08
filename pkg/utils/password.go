package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// สร้างฟังก์ชันสำหรับ เข้ารหัสรหัสผ่าน
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// สร้างฟังก์ชันสำหรับ ตรวจสอบรหัสผ่าน
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}