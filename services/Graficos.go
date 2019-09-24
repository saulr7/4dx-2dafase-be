package services

import (
	"../config"
	"../models"
)

func GraficoPorMCICreate(Modelo models.GraficoPorMCI) (models.GraficoPorMCI, error) {

	Modelo.IdRelacion = 0

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}
