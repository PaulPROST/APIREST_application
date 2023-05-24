package http

import (
	"net/http"
	"strconv"
	"time"

	"pricesAPI/domain"
	"pricesAPI/internal/usecase"

	"github.com/gin-gonic/gin"
)

// HTTP handler for the get price endpoint
// @Param date query string true "Date in the format YYYY-MM-DD-HH.MM.SS"
// @Param brandID query int true "Brand ID"
// @Param productID query int true "Product ID"
func GetPrice(cmd usecase.GetPriceCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		date := c.Query("applicationDate")
		brandID, err := strconv.ParseInt(c.Query("brandID"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brandID. Please provide a valid integer value."})
			return
		}
		productID, err := strconv.ParseInt(c.Query("productID"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid productID. Please provide a valid integer value."})
			return
		}

		// Parse the date parameter
		parsedDate, err := time.Parse("2006-01-02-15.04.00", date)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Please use YYYY-MM-DD-HH.MM.SS"})
			return
		}

		price, err := cmd(parsedDate, productID, brandID)
		if err != nil {
			switch err {
			case domain.ErrPriceNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			default:
				c.Status(http.StatusInternalServerError)
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"price": price,
		})
	}
}
