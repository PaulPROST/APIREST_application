package domain

import (
	"errors"
	"time"
)

var ErrPriceNotFound = errors.New("price has not been found")

type Price struct {
	BrandID   int64     `json:"brandID"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	PriceList int64     `json:"priceList"`
	ProductID int64     `json:"productId"`
	Priority  int64     `json:"priority"`
	Price     float64   `json:"price"`
	Curr      string    `json:"curr"`
}
