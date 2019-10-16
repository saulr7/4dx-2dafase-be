package services

import (
	"../config"
	"../models"
)

func ColaboradoresPorArea(AreaId string) ([]models.Colaborador, error) {

	var colaborador []models.Colaborador

	db := config.ConnectDB()
	defer db.Close()

	var codigos []int

	db.Raw("SELECT idColaborador FROM EquiposSubAreas WHERE idSubArea = ?", AreaId).Pluck("idColaborador", &codigos)

	db.Where("idColaborador IN (?)", codigos).Find(&colaborador)

	return colaborador, nil
}
