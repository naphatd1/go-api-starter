package repositories

import (
	"github.com/naphat/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/naphat/fiber-ecommerce-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

// UserRepositoryImpl เป็นการใช้งาน GORM สำหรับจัดการผู้ใช้
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// CreateUser สร้างผู้ใช้ใหม่
func (r *UserRepositoryImpl) Create(user *entities.User) error {
	userModel := &models.User{}
	userModel.FromEntity(user)

	if err := r.db.Create(userModel).Error; err != nil {
		return err
	}

	*user = *userModel.ToEntity()
	return nil
}

// GetUserByEmail ค้นหาผู้ใช้ตามอีเมล
func (r *UserRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

// GetUserByID ค้นหาผู้ใช้ตาม ID
func (r *UserRepositoryImpl) GetByID(id uint) (*entities.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

// UpdateUser อัปเดตข้อมูลผู้ใช้
func (r *UserRepositoryImpl) Update(user *entities.User) error {
	userModel := &models.User{}
	userModel.FromEntity(user)
	return r.db.Save(userModel).Error
}

// Delete ลบผู้ใช้ตาม ID
func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// GetAllUsers ค้นหาผู้ใช้ทั้งหมด
func (r *UserRepositoryImpl) GetAll() ([]entities.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	var result []entities.User
	for _, user := range users {
		result = append(result, *user.ToEntity())
	}
	return result, nil
}