package infrastructure

import (
	"pricesAPI/domain"
)

type InMemoryDatabase struct {
	prices []domain.Price
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}

// AddPrice adds a price to the in-memory database
func (db *InMemoryDatabase) AddPrice(price []domain.Price) {
	db.prices = append(db.prices, price...)
}
