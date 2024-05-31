package main

import (
	"ntt_data_test/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.POST("/exoplanets", handlers.CreateExoplanet)
	router.GET("/exoplanets", handlers.ListExoplanets)
	router.GET("/exoplanets/:id", handlers.GetExoplanet)
	router.PUT("/exoplanets/:id", handlers.UpdateExoplanet)
	router.DELETE("/exoplanets/:id", handlers.DeleteExoplanet)
	router.GET("/exoplanets/:id/fuel", handlers.EstimateFuel)

	router.Run(":8000")
}
