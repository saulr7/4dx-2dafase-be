package services

import (
	"../config"
	"../models"
)

func GetMediciones() ([]models.Mediciones, error) {

	var result []models.Mediciones

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetMedicion").Scan(&result)

	return result, nil

}

func TipoGraficos() ([]models.TipoGrafico, error) {

	var result []models.TipoGrafico

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetTipoGraficos").Scan(&result)

	return result, nil

}
