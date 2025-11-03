// internal/core/domain/ports/services/stats_service.go

package services

import (
	"context"

	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/entities"
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/ports/repositories"
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/ports/services"
)

type statsService struct {
	statsRepo repositories.StatsRepository
}

func NewStatsService(statsRepo repositories.StatsRepository) services.StatsService {
	return &statsService{
		statsRepo: statsRepo,
	}
}

func (s *statsService) GetSalesStats(ctx context.Context) (*entities.SalesStats, error) {
	return s.statsRepo.GetSalesStats(ctx)
}

func (s *statsService) GetProductStats(ctx context.Context) (*entities.ProductStats, error) {
	return s.statsRepo.GetProductStats(ctx)
}

func (s *statsService) GetUserStats(ctx context.Context) (*entities.UserStats, error) {
	return s.statsRepo.GetUserStats(ctx)
}
