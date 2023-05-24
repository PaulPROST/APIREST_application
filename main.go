package main

import (
	"log"
	"pricesAPI/internal/infrastructure"
	"pricesAPI/internal/transport/http"
	"pricesAPI/internal/usecase"

	"github.com/gin-gonic/gin"
)

func initRoute(router *gin.Engine, inMemorydatabase *infrastructure.InMemoryDatabase) {
	router.GET("/price", http.GetPrice(usecase.GetPrice(inMemorydatabase)))
}

func main() {
	// Create an instance of the in-memory database
	inMemorydatabase := infrastructure.NewInMemoryDatabase()
	inMemorydatabase.InitializeDatabase()

	router := gin.Default()

	initRoute(router, inMemorydatabase)

	err := router.Run()
	if err != nil {
		log.Fatal("error while running the router")
	}
}
