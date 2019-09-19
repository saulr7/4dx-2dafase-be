package services

import("../models"
"../config"

"time")

func GetResultadosMCI(idMCI string) ([]models.ResultadosMCI, error) {

	var result []models.ResultadosMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetResultadosMCI ?", idMCI).Scan(&result)

	return result, nil
}


func ResultadosMCIUpdate(Modelo models.ResultadosMCI)(models.ResultadosMCI, error){

	var updatedResultado models.ResultadosMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idResultado = ?", Modelo.IdResultadoMCI).Find(&updatedResultado)

	db.Model(&updatedResultado).Where("idResultado= ?", Modelo.IdResultadoMCI).Update(models.ResultadosMCI{Valor: Modelo.Valor, FechaModificacion: time.Now()})

	return updatedResultado, nil

}