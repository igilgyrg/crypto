package candle

import (
	"context"
	"github.com/igilgyrg/crypto/pkg/logging"
	"time"
)

type Service interface {
	Store(ctx context.Context, candle *CreateCandle) (string, error)
	GetByDatetime(ctx context.Context, datetime time.Time) (*Candle, error)
	GetBetweenDatetime(ctx context.Context, start time.Time, end time.Time) ([]Candle, error)
}

type service struct {
	storage Storage
	logger  *logging.Logger
}

func NewService(storage Storage, logger *logging.Logger) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

func (s service) Store(ctx context.Context, dto *CreateCandle) (string, error) {
	var candle *Candle

	return s.storage.Store(ctx, candle)
}

func (s service) GetByDatetime(ctx context.Context, datetime time.Time) (*Candle, error) {
	return s.storage.FindOneByDatetime(ctx, datetime)
}

func (s service) GetBetweenDatetime(ctx context.Context, start time.Time, end time.Time) ([]Candle, error) {
	return s.storage.FindAllBetweenDatetime(ctx, start, end)
}
