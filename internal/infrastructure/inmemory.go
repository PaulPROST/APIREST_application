package infrastructure

import (
	"fmt"
	"pricesAPI/domain"
	"time"
)

type InMemoryDatabase struct {
	prices []domain.Price
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}

// InitializeDatabase initializes the in-memory database with some data
func (db *InMemoryDatabase) InitializeDatabase() {
	db.prices = append(db.prices,
		domain.Price{BrandID: 1, StartDate: time.Date(2020, 6, 14, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC), PriceList: 1, ProductID: 35455, Priority: 0, Price: 35.50, CurrencyISO: "EUR"},
		domain.Price{BrandID: 1, StartDate: time.Date(2020, 6, 14, 15, 0, 0, 0, time.UTC), EndDate: time.Date(2020, 6, 14, 18, 30, 0, 0, time.UTC), PriceList: 2, ProductID: 35455, Priority: 1, Price: 25.45, CurrencyISO: "EUR"},
		domain.Price{BrandID: 1, StartDate: time.Date(2020, 6, 15, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2020, 6, 15, 11, 0, 0, 0, time.UTC), PriceList: 3, ProductID: 35455, Priority: 1, Price: 30.50, CurrencyISO: "EUR"},
		domain.Price{BrandID: 1, StartDate: time.Date(2020, 6, 15, 16, 0, 0, 0, time.UTC), EndDate: time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC), PriceList: 4, ProductID: 35455, Priority: 1, Price: 38.95, CurrencyISO: "EUR"},
	)
}

// GetPrice returns the price information for a given application date, product identifier, and string identifier.
func (db *InMemoryDatabase) GetPrice(applicationDay time.Time, productID int64, brandID int64) (*domain.Price, error) {
	var matchingPrices []domain.Price

	// Find the price with the highest priority in the database based on brand, date, and product
	for _, price := range db.prices {
		if applicationDay.After(price.StartDate) && applicationDay.Before(price.EndDate) && price.ProductID == productID && price.BrandID == brandID {
			matchingPrices = append(matchingPrices, price)
		}
	}

	// Return an error if no matching price is found
	if len(matchingPrices) == 0 {
		return nil, domain.ErrPriceNotFound
	}

	// Find the price with the highest priority
	var highestPriorityPrice = matchingPrices[0]
	for _, p := range matchingPrices {
		if highestPriorityPrice.Priority < p.Priority {
			highestPriorityPrice = p
		}
	}
	fmt.Println(highestPriorityPrice)

	return &highestPriorityPrice, nil
}

// GetPrice returns the price information for a given application date, product identifier, and string identifier.
func (db *InMemoryDatabase) GetBasePrice(productID int64, brandID int64) (*domain.Price, error) {
	for _, price := range db.prices {
		if price.ProductID == productID && price.BrandID == brandID && price.PriceList == 1 {
			return &price, nil
		}
	}
	return nil, fmt.Errorf("no price found for the given parameters")
}
