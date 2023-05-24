package http

import (
	"net/http"
	"time"

	"pricesAPI/domain"
	"pricesAPI/internal/usecase"

	"github.com/gin-gonic/gin"
)

type requestParameters struct {
	ApplicationDate time.Time `json:"applicationDate"`
	BrandID         int64     `json:"brandID"`
	ProductID       int64     `json:"productID"`
}

// HTTP handler for the get price endpoint
func GetPrice(cmd usecase.GetPriceCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestParameters
		if err := c.ShouldBind(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		_, err := cmd(c.Request.Context(), req.ApplicationDate, req.ProductID, req.BrandID)
		if err != nil {
			switch err {
			case domain.ErrPriceNotFound:
				c.Status(http.StatusNotFound)
				return
			default:
				c.Status(http.StatusInternalServerError)
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
