package services

import (
	"fmt"

	"../config"
	"../models"
)

func ColaboradoresPorArea(AreaId string) ([]models.Colaborador, error) {

	var colaborador []models.Colaborador

	fmt.Println(AreaId)

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
	fmt.Println("Entra Colaboradores por sub área Servicio")
	
	db.Raw("EXEC dbo.usp_dbgetColaboradoresByEquipo ?", IdSubArea).Scan(&result)

   return result, nil
}