package services

import (
	"../config"
	"../models"
)

func DbPeriodoPorMPUPdate(Estructura models.DbPeriodoPorMP) (models.DbPeriodoPorMP, error) {

	var updatedPeriodo models.DbPeriodoPorMP

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idPeriodo = ?", Estructura.IdPeriodo).Find(&updatedPeriodo)

	db.Model(&updatedPeriodo).Where("idPeriodo= ?", Estructura.IdPeriodo).Update(models.DbPeriodoPorMP{IdFrecuencia: Estructura.IdFrecuencia, IdMedicion: Estructura.IdMedicion})

	return updatedPeriodo, nil
}


func DbPeriodoPorMPAdd(Modelo models.DbPeriodoPorMP) (models.DbPeriodoPorMP, error) {

	Modelo.IdPeriodo = 0

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}