package repositories

import (
	"EchoAPI/core/models"
	"context"
)

type FilterProps struct {
	FromAge *int
	ToAge   *int

	FromDate *int64
	ToDate   *int64
}

type UsersI interface {
	Create(ctx context.Context, user models.User) error

	GetList(ctx context.Context, props FilterProps) ([]models.User, int64, error)
}
