package services

import (
	"fmt"

	"../config"
	"../models"
)

func TableroColaborador(colaboradorId string) ([]models.Tablero, error) {

	var tableros []models.Tablero

	type Result struct {
		IdMCI            int    `gorm:"column:idMCI"`
		MCI              string `gorm:"column:MCI"`
		IdMP             int    `gorm:"column:idMP"`
		MedidaPredictiva string `gorm:"column:MedidaPredictiva "`
	}

	var result []Result

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetMedidasPorColaborador ?", colaboradorId).Scan(&result)

	for _, dato := range result {

		var resultadosMCI []models.ResultadoMCI

		var tablero models.Tablero

		tablero.IdMCI = dato.IdMCI
		tablero.MCI = dato.MCI

		db.Raw("SELECT Mes, Valor, Meta FROM dbResultadosMCI where idMCI = ?", dato.IdMCI).Scan(&resultadosMCI)

		for _, resultado := range resultadosMCI {

			tablero.ResultadosMCI = append(tablero.ResultadosMCI, resultado)
		}

		var medidasPredictiva []models.MedidaPredictiva

		db.Raw("SElect idMP, MedidaPredictiva FROM LVMedidasPredictivas where idMCI = ?", dato.IdMCI).Scan(&medidasPredictiva)

		for _, medida := range medidasPredictiva {

			var resultasdoMP []models.ResultadoMP

			db.Raw("SELECT Valor, Unidad FROM dbResultados WHERE idMP = ?", medida.IdMP).Scan(&resultasdoMP)

			fmt.Println(resultasdoMP)

			for _, resulMP := range resultasdoMP {

				medida.ResultadosMP = append(medida.ResultadosMP, resulMP)
			}

			tablero.MedidaPredictiva = append(tablero.MedidaPredictiva, medida)

		}

		tableros = append(tableros, tablero)
	}

	return tableros, nil
}
