package usecase

import (
	"pricesAPI/internal/infrastructure"
	"time"
)

type GetPriceCmd func(applicationDate time.Time, ProductID int64, StringID int64) (RequestResponse, error)

type RequestResponse struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	BrandID   int64     `json:"stringID"`
	ProductID int64     `json:"productId"`
	Price     float64   `json:"price"`
	Rate      float64   `json:"rate"`
}

// GetPrice returns a function that implements the business logic for the get price endpoint
func GetPrice(inMemorydatabase *infrastructure.InMemoryDatabase) GetPriceCmd {
	return func(applicationDate time.Time, ProductID int64, BrandID int64) (RequestResponse, error) {
		// Get the price from the database
		price, err := inMemorydatabase.GetPrice(applicationDate, ProductID, BrandID)
		if err != nil {
			return RequestResponse{}, err
		}

		// Get the base price from the database to calculate the rate
		basePrice, err := inMemorydatabase.GetBasePrice(ProductID, BrandID)
		rate := 0.0
		if err == nil {
			rate = (price.Price - basePrice.Price) / basePrice.Price * 100
		}

		return RequestResponse{
			StartDate: price.StartDate,
			EndDate:   price.EndDate,
			BrandID:   BrandID,
			ProductID: ProductID,
			Price:     price.Price,
			Rate:      rate,
		}, nil
	}
}
