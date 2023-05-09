package controllers

import (
	"crud/config"
	"crud/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateTransaction(c echo.Context) error {
	transaction := new(models.Transactions)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if err := config.DB.Create(&transaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, transaction)
}

func GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := new(models.Transactions)
	if err := config.DB.First(&transaction, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "transaction not found"})
	}
	return c.JSON(http.StatusOK, transaction)
}

func UpdateTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := new(models.Transactions)
	if err := config.DB.First(&transaction, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "transaction not found"})
	}
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if err := config.DB.Save(&transaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := new(models.Transactions)
	if err := config.DB.First(&transaction, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "transaction not found"})
	}
	if err := config.DB.Delete(&transaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "transaction deleted"})
}

func GetTransactions(c echo.Context) error {
	transactions := []models.Transactions{}
	if err := config.DB.Find(&transactions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, transactions)
}
