package domain

import (
	"errors"
	"time"
)

var ErrPriceNotFound = errors.New("no price has been found")

type Price struct {
	BrandID     int64     `json:"brand_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	PriceList   int       `json:"price_list"`
	ProductID   int64     `json:"product_id"`
	Priority    int64     `json:"priority"`
	Price       float64   `json:"price"`
	CurrencyISO string    `json:"currency_iso"`
}
