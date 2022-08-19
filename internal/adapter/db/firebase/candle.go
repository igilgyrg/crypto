package firebase

import (
	"context"
	candle2 "github.com/igilgyrg/crypto/internal/domain/candle"
	"github.com/igilgyrg/crypto/pkg/logging"
	"time"
)

type repository struct {
	logger *logging.Logger
}

func NewRepository(logger *logging.Logger) candle2.Storage {
	return &repository{
		logger: logger,
	}
}

func (r repository) Store(ctx context.Context, candle *candle2.Candle) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindOneByDatetime(ctx context.Context, datetime time.Time) (*candle2.Candle, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindAllBetweenDatetime(ctx context.Context, start time.Time, end time.Time) ([]candle2.Candle, error) {
	//TODO implement me
	panic("implement me")
}
