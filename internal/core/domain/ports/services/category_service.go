// internal/core/domain/ports/services/category_service.go

package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/entities"
)

// CategoryService interface สำหรับการจัดการหมวดหมู่
type CategoryService interface {
	CreateCategory(ctx context.Context, req *entities.CreateCategoryRequest) (*entities.Category, error)
	GetCategories(ctx context.Context, page, limit int) ([]*entities.Category, *entities.PaginationResponse, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*entities.Category, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, req *entities.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}
