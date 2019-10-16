package services

import (
	"../config"
	"../models"
)

func GetFrecuencias() ([]models.Frecuencia, error) {

	var result []models.Frecuencia

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetFrecuencias").Scan(&result)

	return result, nil

}
