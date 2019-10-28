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

func GetColaboradoresSubArea(IdSubArea string) ([]models.Colaborador, error) {

	var result []models.Colaborador

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC dbo.usp_dbgetColaboradoresByEquipo ?", IdSubArea).Scan(&result)

	return result, nil
}

func ColaboradoresAdmins() ([]models.Colaborador, error) {

	var colaboradores []models.Colaborador

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("SELECT NombreColaborador, idColaborador, Correo FROM Colaboradores where idPerfil = ?", 2).Scan(&colaboradores)

	return colaboradores, nil
}
