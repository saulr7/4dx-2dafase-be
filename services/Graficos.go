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

func GetGraficoColaborador(idEmpleado string) ([]models.GraficoGlobal, error) {

	var result []models.GraficoGlobal

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetGraficosAnualesColaborador ?", idEmpleado).Scan(&result)

	return result, nil
}
