package services

import (
	"time"

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

	t := time.Now()
	Modelo.IdPeriodo = 0
	Modelo.Anio = t.Year()
	Modelo.FechaCreacion = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func PeriodoPorMCIAdd(Modelo models.PeriodoPorMCI) (models.PeriodoPorMCI, error) {

	Modelo.IdPeriodo = 0

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func GetPeriodicidadMCI() ([]models.PeriodicidadMCI, error) {

	var result []models.PeriodicidadMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetPeriodicidadMCI").Scan(&result)

	return result, nil
}
