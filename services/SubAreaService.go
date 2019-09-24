package services

import (
	"../config"
	"../models"
)

func SubAreasAll() ([]models.SubArea, error) {

	var SubArea []models.SubArea

	db := config.ConnectDB()
	defer db.Close()

	db.Order("SubArea").Find(&SubArea)

	return SubArea, nil
}
