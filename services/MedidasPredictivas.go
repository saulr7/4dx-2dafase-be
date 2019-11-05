package services

import (
	"fmt"

	"../config"
	"../models"
)

func GetMedidasPredictivas(idEmpleado string) ([]models.GetMPbyMCI, error) {

	var result []models.GetMPbyMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbgetMPbyMCI ?", idEmpleado).Scan(&result)

	return result, nil
}

func GetMCIColaborador(idEmpleado string) ([]models.GetMPbyMCI, error) {

	var result []models.GetMPbyMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbgMCIbyColaborador ?", idEmpleado).Scan(&result)

	return result, nil
}

func GetMetasColaborador(colaboradorId string) ([]models.Tablero, error) {

	var tableros []models.Tablero

	type Result struct {
		IdMCI            int    `gorm:"column:idMCI"`
		MCI              string `gorm:"column:MCI"`
		IdMP             int    `gorm:"column:idMP"`
		MedidaPredictiva string `gorm:"column:MedidaPredictiva "`
		Orden            int    `gorm:"column:orden"`
	}

	var result []Result

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetMCIbyColaborador ?", colaboradorId).Scan(&result)

	for _, dato := range result {

		var tablero models.Tablero

		tablero.IdMCI = dato.IdMCI
		tablero.MCI = dato.MCI
		tablero.Orden = dato.Orden

		type Frecuencia struct {
			FrecuenciaId int `gorm:"column:idFrecuencia"`
		}

		var frecuencia Frecuencia

		db.Table("dbPeriodoPorMCI").Select("idFrecuencia").Where(" idMCI = ?", dato.IdMCI).Scan(&frecuencia)

		tablero.Periodicidad = frecuencia.FrecuenciaId

		var medidasPredictiva []models.MedidaPredictiva

		type Autorizado struct {
			Autorizado bool `gorm:"column:Autorizado"`
		}

		var estaAutorizado Autorizado

		db.Raw("SELECT COUNT(1) as Autorizado from  dbPeriodoPorMCI WHERE  idMCI  = ?", dato.IdMCI).Scan(&estaAutorizado)

		tablero.Autorizado = estaAutorizado.Autorizado
		fmt.Println("aUROR", estaAutorizado.Autorizado)

		db.Raw("Select idMP,MedidaPredictiva,ROW_NUMBER()OVER(ORDER BY idMP)OrdenMP FROM LVMedidasPredictivas where idMCI = ?", dato.IdMCI).Scan(&medidasPredictiva)

		for _, medida := range medidasPredictiva {

			// var resultasdoMP []models.ResultadoMP

			var meta models.MedidaPredictiva
			var metaMedicion models.MedidaPredictiva

			db.Table("dbPeriodoPorMP").Select("idMedicion, idFrecuencia").Where(" idMP = ?", medida.IdMP).Scan(&metaMedicion)

			medida.MetaMP = meta.MetaMP
			medida.MedicionId = metaMedicion.MedicionId
			medida.FrecuenciaId = metaMedicion.FrecuenciaId

			// for _, resulMP := range resultasdoMP {

			// 	medida.ResultadosMP = append(medida.ResultadosMP, resulMP)
			// 	tablero.Unidad = resulMP.Unidad
			// }

			tablero.MedidaPredictiva = append(tablero.MedidaPredictiva, medida)

		}

		tableros = append(tableros, tablero)
	}

	return tableros, nil
}
