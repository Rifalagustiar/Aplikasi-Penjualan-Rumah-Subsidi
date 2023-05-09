package controllers

import (
	"crud/config"
	"crud/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreatePropertyEndpoint is a controller function to create a new property
func CreatePropertyEndpoint(c echo.Context) error {
	property := new(models.Properties)
	if err := c.Bind(property); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&property).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    property,
	})
}

// GetPropertyEndpoint is a controller function to get a single property by ID
func GetPropertyEndpoint(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	property := new(models.Properties)
	if err := config.DB.Preload("Agent").Preload("Transactions").First(&property, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"user":    property,
	})

}

// UpdatePropertyEndpoint is a controller function to update an existing property by ID
func UpdatePropertyEndpoint(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	property := new(models.Properties)
	if err := config.DB.First(&property, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(property); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&property).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Agent Properties successfully",
		"user":    property,
	})
}

// DeletePropertyEndpoint is a controller function to delete a property by ID
func DeletePropertyEndpoint(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	property := new(models.Properties)
	if err := config.DB.First(&property, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&property).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Properties by id",
	})
}

// ListPropertiesEndpoint is a controller function to list all properties
func ListPropertiesEndpoint(c echo.Context) error {
	var properties []models.Properties
	if err := config.DB.Preload("Agent").Preload("Transactions").Find(&properties).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, properties)
}
