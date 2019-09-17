package services

import (
	"../config"
	"../models"
)

func DbPeriodoPorMPAdd(Estructura models.DbPeriodoPorMP) (models.DbPeriodoPorMP, error) {

	Estructura.IdPeriodo = 0

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Estructura)

	return Estructura, nil
}

func DbPeriodoPorMPUPdate(Estructura models.DbPeriodoPorMP) (models.DbPeriodoPorMP, error) {

	var updatedPeriodo models.DbPeriodoPorMP

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idPeriodo = ?", Estructura.IdPeriodo).Find(&updatedPeriodo)

	db.Model(&updatedPeriodo).Where("idPeriodo= ?", Estructura.IdPeriodo).Update(models.DbPeriodoPorMP{IdPeriodo: Estructura.IdPeriodo, IdMP: Estructura.IdMP, IdFrecuencia: Estructura.IdFrecuencia, IdMedicion: Estructura.IdMedicion})

	return updatedPeriodo, nil
}
