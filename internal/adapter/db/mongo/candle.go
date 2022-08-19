package mongo

import (
	"context"
	"fmt"
	candle2 "github.com/igilgyrg/crypto/internal/domain/candle"
	exception "github.com/igilgyrg/crypto/internal/error"
	"github.com/igilgyrg/crypto/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type repository struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewRepository(database *mongo.Database, collection string, logger *logging.Logger) candle2.Storage {
	return &repository{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (r repository) Store(ctx context.Context, candle *candle2.Candle) (string, error) {
	res, err := r.collection.InsertOne(ctx, candle)
	if err != nil {
		return "", fmt.Errorf("error of inserting candle %v", err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("error of convert insert id to primitive %v", err)
}

func (r repository) FindOneByDatetime(ctx context.Context, datetime time.Time) (*candle2.Candle, error) {
	filter := bson.M{"datetime": datetime}

	var result candle2.Candle
	res := r.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, exception.ErrNotFound(res.Err(), "candle not found")
		}
		return nil, fmt.Errorf("error of getting element by datetime")
	}

	err := res.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error of decode mongo document %v", err)
	}

	return &result, nil
}

func (r repository) FindAllBetweenDatetime(ctx context.Context, start time.Time, end time.Time) ([]candle2.Candle, error) {
	filter := bson.M{"datetime": bson.M{
		"$gte": start,
		"$lt":  end,
	}}

	var result []candle2.Candle
	find, err := r.collection.Find(ctx, filter)
	if find.Err() != nil || err != nil {
		if find.Err() == mongo.ErrNoDocuments {
			return nil, exception.ErrNotFound(find.Err(), "candle not found")
		}
		return nil, fmt.Errorf("error of getting element by datetime")
	}

	err = find.All(ctx, &result)
	if err != nil {
		return nil, fmt.Errorf("error of decode mongo documents %v", err)
	}

	return result, nil
}
