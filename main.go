package main

import (
	"log"
	"pricesAPI/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

func initRoute(router *gin.Engine, inMemorydatabase *infrastructure.InMemoryDatabase) {
	// router.GET("/price", )
}

func main() {
	// Create an instance of the in-memory database
	inMemorydatabase := infrastructure.NewInMemoryDatabase()

	router := gin.Default()

	initRoute(router, inMemorydatabase)

	err := router.Run()
	if err != nil {
		log.Fatal("error while running the router")
	}
}
