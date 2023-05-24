package http_test

import (
	"pricesAPI/internal/infrastructure"
	"pricesAPI/internal/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var wantPrices = []usecase.RequestResponse{
	{
		StartDate: time.Date(2020, 6, 14, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC),
		BrandID:   1,
		ProductID: 35455,
		Price:     35.50,
		Rate:      0.00,
	}, {
		StartDate: time.Date(2020, 6, 14, 15, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2020, 6, 14, 18, 30, 0, 0, time.UTC),
		BrandID:   1,
		ProductID: 35455,
		Price:     25.45,
		Rate:      -28.30985915492958,
	},
	{
		StartDate: time.Date(2020, 6, 14, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC),
		BrandID:   1,
		ProductID: 35455,
		Price:     35.5,
		Rate:      0,
	},
	{
		StartDate: time.Date(2020, 6, 15, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2020, 6, 15, 11, 0, 0, 0, time.UTC),
		BrandID:   1,
		ProductID: 35455,
		Price:     30.5,
		Rate:      -14.084507042253522,
	},
	{
		StartDate: time.Date(2020, 6, 15, 16, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC),
		BrandID:   1,
		ProductID: 35455,
		Price:     38.95,
		Rate:      9.718309859154937,
	},
}

func TestGetPrice(t *testing.T) {
	inMemorydatabase := infrastructure.NewInMemoryDatabase()

	// Initialize the database with the test data
	inMemorydatabase.InitializeDatabase()

	for _, tt := range []struct {
		name string

		wantResponse    usecase.RequestResponse
		applicationDate time.Time
		brandID         int64
		productID       int64
	}{
		{
			name:            "Test 1",
			applicationDate: time.Date(2020, 6, 14, 10, 0, 0, 0, time.UTC),
			productID:       35455,
			brandID:         1,
			wantResponse:    wantPrices[0],
		},
		{
			name:            "Test 2",
			applicationDate: time.Date(2020, 6, 14, 16, 0, 0, 0, time.UTC),
			productID:       35455,
			brandID:         1,
			wantResponse:    wantPrices[1],
		},
		{
			name:            "Test 3",
			applicationDate: time.Date(2020, 6, 14, 21, 0, 0, 0, time.UTC),
			productID:       35455,
			brandID:         1,
			wantResponse:    wantPrices[2],
		},
		{
			name:            "Test 4",
			applicationDate: time.Date(2020, 6, 15, 10, 0, 0, 0, time.UTC),
			productID:       35455,
			brandID:         1,
			wantResponse:    wantPrices[3],
		},
		{
			name:            "Test 5",
			applicationDate: time.Date(2020, 6, 16, 21, 0, 0, 0, time.UTC),
			productID:       35455,
			brandID:         1,
			wantResponse:    wantPrices[4],
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			price, err := usecase.GetPrice(inMemorydatabase)(tt.applicationDate, tt.productID, tt.brandID)
			require.NoError(t, err)

			assert.Equal(t, tt.wantResponse, price)
		})
	}
}
