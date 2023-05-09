package controllers

import (
	"crud/config"
	"crud/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAgentsController mengembalikan daftar agen dari database
func GetAgentsController(c echo.Context) error {
	var agents []models.Agents
	if err := config.DB.Find(&agents).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve agents")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": agents,
	})

}

// CreateAgentController menambahkan agen baru ke database
func CreateAgentController(c echo.Context) error {
	// Menerima informasi agen dari request body
	agent := models.Agents{}
	c.Bind(&agent)
	// Menambahkan agen baru ke database
	if err := config.DB.Create(&agent).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create agent")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Agents created successfully",
		"data":    agent,
	})
}

// UpdateAgentController mengubah informasi agen di database
func UpdateAgentController(c echo.Context) error {
	// Menerima informasi agen dari request body
	agent := models.Agents{}
	c.Bind(&agent) // Mencari agen yang akan diubah di database
	var dbAgent models.Agents
	if err := config.DB.First(&dbAgent, agent.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Agents not found")
	}

	// Mengubah informasi agen di database
	if err := config.DB.Model(&dbAgent).Updates(&agent).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update agent")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Agents updated successfully",
		"data":    dbAgent,
	})
}

// DeleteAgentController menghapus agen dari database
func DeleteAgentController(c echo.Context) error {
	// Menerima ID agen dari parameter URL
	id := c.Param("id") // Menghapus agen dari database
	if err := config.DB.Where("id = ?", id).Delete(&models.Agents{}).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete agent")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Agents deleted successfully",
	})
}
