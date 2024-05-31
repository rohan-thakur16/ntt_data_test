package services

import (
	"errors"
	"fmt"
	"math"
	"ntt_data_test/models"
	"sync"
)

// In-memory data storage
var (
	exoplanets = make(map[string]models.Exoplanet)
	mutex      = &sync.Mutex{}
)

// CreateExoplanet creates a new exoplanet and stores it in memory
func CreateExoplanet(exoplanet models.Exoplanet) (models.Exoplanet, error) {
	exoplanet.ID = generateID()

	mutex.Lock()
	exoplanets[exoplanet.ID] = exoplanet
	mutex.Unlock()

	return exoplanet, nil
}

// ListExoplanets lists all stored exoplanets
func ListExoplanets() []models.Exoplanet {
	mutex.Lock()
	defer mutex.Unlock()

	var planetList []models.Exoplanet
	for _, planet := range exoplanets {
		planetList = append(planetList, planet)
	}

	return planetList
}

// GetExoplanet retrieves an exoplanet by ID
func GetExoplanet(id string) (models.Exoplanet, error) {
	mutex.Lock()
	defer mutex.Unlock()

	exoplanet, exists := exoplanets[id]
	if !exists {
		return exoplanet, errors.New("exoplanet not found")
	}

	return exoplanet, nil
}

// UpdateExoplanet updates an existing exoplanet
func UpdateExoplanet(id string, updatedExoplanet models.Exoplanet) (models.Exoplanet, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := exoplanets[id]; !exists {
		return updatedExoplanet, errors.New("exoplanet not found")
	}

	updatedExoplanet.ID = id
	exoplanets[id] = updatedExoplanet
	return updatedExoplanet, nil
}

// DeleteExoplanet removes an exoplanet from memory
func DeleteExoplanet(id string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := exoplanets[id]; !exists {
		return errors.New("exoplanet not found")
	}
	delete(exoplanets, id)
	return nil
}

// EstimateFuel estimates the fuel cost for a trip to an exoplanet
func EstimateFuel(id string, crewCapacity int) (float64, error) {
	mutex.Lock()
	exoplanet, exists := exoplanets[id]
	mutex.Unlock()

	if !exists {
		return 0, errors.New("exoplanet not found")
	}

	var gravity float64
	switch exoplanet.Type {
	case models.GasGiant:
		gravity = 0.5 / math.Pow(exoplanet.Radius, 2)
	case models.Terrestrial:
		gravity = exoplanet.Mass / math.Pow(exoplanet.Radius, 2)
	default:
		return 0, errors.New("invalid exoplanet type")
	}

	fuelCost := float64(exoplanet.Distance) / math.Pow(gravity, 2) * float64(crewCapacity)
	return fuelCost, nil
}

// Helper function to generate unique IDs
func generateID() string {
	return fmt.Sprintf("%d", len(exoplanets)+1)
}
