package domain

import "time"

type Price struct {
	BrandID   int       `json:"brandID"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	PriceList int       `json:"priceList"`
	ProductID int       `json:"productId"`
	Priority  int       `json:"priority"`
	Price     float64   `json:"price"`
	Curr      string    `json:"curr"`
}
