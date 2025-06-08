package services

import (
	"errors"
	"testing"

	"github.com/naphat/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/naphat/fiber-ecommerce-api/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) GetByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepo) GetByID(id uint) (*entities.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepo) Update(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepo) GetAll() ([]entities.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.User), args.Error(1)
}

// --- TEST Register ---

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	authService := NewAuthService(mockRepo)

	req := entities.RegisterRequest{
		Email:     "newuser@example.com",
		Password:  "password123",
		FirstName: "New",
		LastName:  "User",
	}

	// ไม่มี user ซ้ำ
	mockRepo.On("GetByEmail", req.Email).Return(nil, errors.New("not found"))
	// สร้าง user สำเร็จ
	mockRepo.On("Create", mock.AnythingOfType("*entities.User")).Return(nil)

	user, err := authService.Register(req)

	assert.NoError(t, err)
	assert.Equal(t, req.Email, user.Email)
	assert.NotEmpty(t, user.Password) // ต้องเป็น hashed password
	mockRepo.AssertExpectations(t)
}

func TestRegister_UserExists(t *testing.T) {
	mockRepo := new(MockUserRepo)
	authService := NewAuthService(mockRepo)

	existingUser := &entities.User{Email: "exists@example.com"}

	mockRepo.On("GetByEmail", existingUser.Email).Return(existingUser, nil)

	req := entities.RegisterRequest{
		Email:    existingUser.Email,
		Password: "password123",
	}

	user, err := authService.Register(req)

	assert.Nil(t, user)
	assert.EqualError(t, err, "user already exists")
	mockRepo.AssertExpectations(t)
}

// --- TEST Login ---

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	authService := NewAuthService(mockRepo)

	hashedPass, _ := utils.HashPassword("password123")

	userFromDB := &entities.User{
		ID:       1,
		Email:    "loginuser@example.com",
		Password: hashedPass,
		IsActive: true,
		Role:     entities.RoleUser,
	}

	mockRepo.On("GetByEmail", userFromDB.Email).Return(userFromDB, nil)

	req := entities.LoginRequest{
		Email:    userFromDB.Email,
		Password: "password123",
	}

	resp, err := authService.Login(req)

	assert.NoError(t, err)
	assert.Equal(t, userFromDB.Email, resp.User.Email)
	assert.NotEmpty(t, resp.Token) // ต้องได้ token
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidPassword(t *testing.T) {
	mockRepo := new(MockUserRepo)
	authService := NewAuthService(mockRepo)

	hashedPass, _ := utils.HashPassword("correctpassword")

	userFromDB := &entities.User{
		Email:    "loginuser@example.com",
		Password: hashedPass,
		IsActive: true,
	}

	mockRepo.On("GetByEmail", userFromDB.Email).Return(userFromDB, nil)

	req := entities.LoginRequest{
		Email:    userFromDB.Email,
		Password: "wrongpassword",
	}

	resp, err := authService.Login(req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "invalid email or password")
	mockRepo.AssertExpectations(t)
}

func TestLogin_InactiveUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	authService := NewAuthService(mockRepo)

	userFromDB := &entities.User{
		Email:    "inactive@example.com",
		Password: "irrelevant",
		IsActive: false,
	}

	mockRepo.On("GetByEmail", userFromDB.Email).Return(userFromDB, nil)

	req := entities.LoginRequest{
		Email:    userFromDB.Email,
		Password: "any",
	}

	resp, err := authService.Login(req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "account is deactivated")
	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepo)
	authService := NewAuthService(mockRepo)

	mockRepo.On("GetByEmail", "notfound@example.com").Return(nil, errors.New("not found"))

	req := entities.LoginRequest{
		Email:    "notfound@example.com",
		Password: "any",
	}

	resp, err := authService.Login(req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "invalid email or password")
	mockRepo.AssertExpectations(t)
}
