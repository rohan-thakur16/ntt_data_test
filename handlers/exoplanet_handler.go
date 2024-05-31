package handlers

///Users/rohanthakur/Documents/ntt_data_test/models
import (
	"net/http"
	"strconv"

	"ntt_data_test/models"
	"ntt_data_test/services"
	utils "ntt_data_test/utlis"
	"ntt_data_test/validation"

	"github.com/gin-gonic/gin"
)

// Create a new Exoplanet
func CreateExoplanet(c *gin.Context) {
	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validation.ValidateExoplanet(exoplanet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	createdExoplanet, err := services.CreateExoplanet(exoplanet)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, createdExoplanet)
}

// List all Exoplanets
func ListExoplanets(c *gin.Context) {
	exoplanets := services.ListExoplanets()
	utils.RespondWithJSON(c, http.StatusOK, exoplanets)
}

// Get a specific Exoplanet by ID
func GetExoplanet(c *gin.Context) {
	id := c.Param("id")

	exoplanet, err := services.GetExoplanet(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, exoplanet)
}

// Update an Exoplanet
func UpdateExoplanet(c *gin.Context) {
	id := c.Param("id")

	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validation.ValidateExoplanet(exoplanet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedExoplanet, err := services.UpdateExoplanet(id, exoplanet)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, updatedExoplanet)
}

// Delete an Exoplanet
func DeleteExoplanet(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteExoplanet(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// Estimate fuel cost
func EstimateFuel(c *gin.Context) {
	id := c.Param("id")

	crewCapacityStr := c.Query("crew")
	if crewCapacityStr == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Crew capacity is required")
		return
	}

	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid crew capacity")
		return
	}

	fuelCost, err := services.EstimateFuel(id, crewCapacity)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, map[string]float64{"fuelCost": fuelCost})
}
