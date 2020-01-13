package services

import (
	"../config"
	"../models"
)

func TableroColaborador(colaboradorId string, mesId string) ([]models.Tablero, error) {

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

		var resultadosMCI []models.ResultadoMCI

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

		db.Raw("SELECT Mes, Valor, Meta, Unidad FROM dbResultadosMCI where idMCI = ? AND Anio = YEAR(GETDATE())", dato.IdMCI).Scan(&resultadosMCI)

		for _, resultado := range resultadosMCI {
			tablero.ResultadosMCI = append(tablero.ResultadosMCI, resultado)
			tablero.Unidad = resultado.Unidad
		}

		var medidasPredictiva []models.MedidaPredictiva

		db.Raw("Select idMP,MedidaPredictiva,CriterioVerde,CriterioAmarillo,CriterioRojo,ROW_NUMBER()OVER(ORDER BY idMP)OrdenMP FROM LVMedidasPredictivas where idMCI = ?", dato.IdMCI).Scan(&medidasPredictiva)

		for _, medida := range medidasPredictiva {

			var resultasdoMP []models.ResultadoMP

			var meta models.MedidaPredictiva
			var metaMedicion models.MedidaPredictiva

			db.Raw("SELECT Mes, Semana ,FORMAT(ns.FechaInicioRendicion, 'dd/MM') + ' - ' +  FORMAT(ns.FechaFinRendicion, 'dd/MM') 'Periodo', d.Dia, Valor, Unidad, LlegoAMeta FROM dbResultados d inner join dbNoSemana ns on d.idNoSemana = ns.idNoSemana WHERE Mes = ? and idMP = ?", mesId, medida.IdMP).Scan(&resultasdoMP)

			db.Table("LVMPCampos").Select("ValorCampo").Where(" CampoMP = 'ValorY' AND  idMP = ?", medida.IdMP).Scan(&meta)

			db.Table("dbPeriodoPorMP").Select("idMedicion, idFrecuencia").Where(" idMP = ?", medida.IdMP).Scan(&metaMedicion)

			medida.MetaMP = meta.MetaMP
			medida.MedicionId = metaMedicion.MedicionId
			medida.FrecuenciaId = metaMedicion.FrecuenciaId

			for _, resulMP := range resultasdoMP {

				medida.ResultadosMP = append(medida.ResultadosMP, resulMP)
			}

			tablero.MedidaPredictiva = append(tablero.MedidaPredictiva, medida)

		}

		tableros = append(tableros, tablero)
	}

	return tableros, nil
}

func GetEstiloTablero(idEmpleado string) (string, error) {

	type resultEstilo struct {
		Estilo string `gorm:"column:Estilo"`
	}

	var result resultEstilo

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEstiloTablero ?", idEmpleado).Scan(&result)

	return result.Estilo, nil
}

func UpdateEstiloTablero(modelo models.ActualizarEstiloTablero) (string, error) {

	type resultEstilo struct {
		Estilo string `gorm:"column:Estilo"`
	}

	var result resultEstilo

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_ActualizarEstiloTablero ?, ?", modelo.Colaborador, modelo.Estilo).Scan(&result)

	return result.Estilo, nil
}
