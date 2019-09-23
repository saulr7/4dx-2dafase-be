package services

import (
	"../config"
	"../models"
)

func TipoGraficos() ([]models.TipoGrafico, error) {

	var graficos []models.TipoGrafico

	db := config.ConnectDB()
	defer db.Close()

	db.Find(&graficos)

	return graficos, nil
}
