package models

import (
	"time"

	"github.com/naphat/fiber-ecommerce-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

// สร้างโมเดล User สำหรับ GORM
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	FirstName string         `gorm:"not null" json:"first_name"`
	LastName  string         `gorm:"not null" json:"last_name"`
	Role      entities.Role  `gorm:"type:varchar(20);default:'user'" json:"role"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// ToEntity แปลงโมเดล User เป็น Entity
func (u *User) ToEntity() *entities.User {
	return &entities.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// FromEntity แปลง Entity เป็นโมเดล User
func (u *User) FromEntity(entity *entities.User) {
	u.ID = entity.ID
	u.Email = entity.Email
	u.Password = entity.Password
	u.FirstName = entity.FirstName
	u.LastName = entity.LastName
	u.Role = entity.Role
	u.IsActive = entity.IsActive
	u.CreatedAt = entity.CreatedAt
	u.UpdatedAt = entity.UpdatedAt
}