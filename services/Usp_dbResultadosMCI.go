package services

import (
	
	"../config"
	"../models"
)


func GetResultadosMCI(idMCI string) ([]models.ResultadosMCI, error) {

	var result []models.ResultadosMCI

	db := config.ConnectDB()
	defer db.Close()
   
	db.Raw("EXEC usp_dbGetResultadosMCI ?", idMCI).Scan(&result)	

   return result, nil	
}