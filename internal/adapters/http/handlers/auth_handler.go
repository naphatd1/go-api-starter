package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/naphat/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/naphat/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/naphat/fiber-ecommerce-api/pkg/utils"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// ฟังก์ชันสำหรับลงทะเบียนผู้ใช้ใหม่
// Register godoc
// @Summary Register a new user
// @Description Register a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.RegisterRequest true "Registration data"
// @Success 201 {object} entities.User
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {

	var req entities.RegisterRequest

	// ตรวจสอบว่า request body ถูกต้องหรือไม่
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// ตรวจสอบความถูกต้องของข้อมูลที่ส่งมา
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// เรียกใช้ service เพื่อทำการลงทะเบียนผู้ใช้
	user, err := h.authService.Register(req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)

}

// ฟังก์ชันสำหรับเข้าสู่ระบบ
// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.LoginRequest true "Login credentials"
// @Success 200 {object} entities.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req entities.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.authService.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}

// ฟังก์ชัน getProfile สำหรับดึงข้อมูลโปรไฟล์ของผู้ใช้
// GetProfile godoc
// @Summary Get user profile
// @Description Get current user's profile information
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} entities.User
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/user/profile [get]
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	id, _ := strconv.ParseUint(userID, 10, 32)

	user, err := h.authService.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}