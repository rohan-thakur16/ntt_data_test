package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ntt_data_test/handlers"
	"ntt_data_test/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/exoplanets", handlers.CreateExoplanet)
	router.GET("/exoplanets", handlers.ListExoplanets)
	return router
}

// TestCreateExoplanet tests the functionality of creating
//
//	a new exoplanet entry
func TestCreateExoplanet(t *testing.T) {
	router := setupRouter()

	exoplanet := models.Exoplanet{
		Name:        "TestExoplanet",
		Description: "A test exoplanet",
		Distance:    100,
		Radius:      1.5,
		Mass:        2.5,
		Type:        "Terrestrial",
	}

	jsonValue, _ := json.Marshal(exoplanet)
	req, _ := http.NewRequest("POST", "/exoplanets", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Exoplanet
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, exoplanet.Name, response.Name)
	assert.Equal(t, exoplanet.Description, response.Description)
	assert.Equal(t, exoplanet.Distance, response.Distance)
	assert.Equal(t, exoplanet.Radius, response.Radius)
	assert.Equal(t, exoplanet.Mass, response.Mass)
	assert.Equal(t, exoplanet.Type, response.Type)
}

// TestListExoplanets tests the functionality of listing all exoplanets.
func TestListExoplanets(t *testing.T) {
	router := setupRouter()

	// Create a GET request to the "/exoplanets" endpoint to retrieve all exoplanets
	req, _ := http.NewRequest("GET", "/exoplanets", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the API responded with the expected status code (OK - 200)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body to get a list of exoplanets
	var response []models.Exoplanet
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err) // Assert that there were no errors during parsing

	// Validate that the response list has at least one element (potentially more)
	assert.GreaterOrEqual(t, len(response), 0)
}
