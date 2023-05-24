package usecase

import (
	"context"
	"pricesAPI/internal/infrastructure"
	"time"
)

type GetPriceCmd func(ctx context.Context, applicationDate time.Time, ProductID int64, StringID int64) (requestResponse, error)

type requestResponse struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	StringID  int64     `json:"stringID"`
	ProductID int64     `json:"productId"`
	Price     float64   `json:"price"`
	Rate      float64   `json:"rate"`
}

// GetPrice returns a function that implements the business logic for the get price endpoint
func GetPrice(inMemorydatabase *infrastructure.InMemoryDatabase) GetPriceCmd {
	return func(ctx context.Context, applicationDate time.Time, ProductID int64, BrandID int64) (requestResponse, error) {
		return requestResponse{}, nil
	}
}
