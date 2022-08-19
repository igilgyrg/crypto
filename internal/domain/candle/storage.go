package candle

import (
	"context"
	"time"
)

type Storage interface {
	Store(ctx context.Context, candle *Candle) (string, error)
	FindOneByDatetime(ctx context.Context, datetime time.Time) (*Candle, error)
	FindAllBetweenDatetime(ctx context.Context, start time.Time, end time.Time) ([]Candle, error)
}
