// internal/adapters/http/handlers/admin_handle.go

package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

// GetAdminDashboard godoc
// @Summary Get admin dashboard
// @Description Get admin dashboard information (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/admin/dashboard [get]
func (h *AdminHandler) GetDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the admin dashboard",
		"user_id": c.Locals("user_id"),
		"role":    c.Locals("role"),
	})
}
