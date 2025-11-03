// internal/core/domain/ports/services/payment_service.go

package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/core/domain/entities"
)

// PaymentService interface สำหรับการจัดการการชำระเงิน
type PaymentService interface {
	CreatePayment(ctx context.Context, req *entities.CreatePaymentRequest) (*entities.Transaction, error)
	GetPaymentByID(ctx context.Context, id uuid.UUID) (*entities.Transaction, error)
	VerifyPayment(ctx context.Context, id uuid.UUID, req *entities.VerifyPaymentRequest) error
	CancelPayment(ctx context.Context, id uuid.UUID) error
}
