package infrastructure

import (
	"pricesAPI/domain"
)

type InMemoryDatabase struct {
	prices []domain.Price
}

// NewCache ...
func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}

func (db *InMemoryDatabase) AddPrice(price []domain.Price) {
	db.prices = append(db.prices, price...)
}
