package services

import (
	"../config"
	"../models"

	"time"
	"errors"
)

func GetResultadosMCI(idMCI string) ([]models.ResultadosMCI, error) {

	var result []models.ResultadosMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetResultadosMCI ?", idMCI).Scan(&result)

	return result, nil
}

func ResultadosMCIUpdate(Modelo models.ResultadosMCI) (models.ResultadosMCI, error) {

	var updatedResultado models.ResultadosMCI
	
	
	 if Modelo.Autorizado == true{
	 	db := config.ConnectDB()
	 	defer db.Close()
		 
		 db.Where("idResultadoMCI = ?", Modelo.IdResultadoMCI).Find(&updatedResultado)

	 	db.Model(&updatedResultado).Where("idResultadoMCI= ?", Modelo.IdResultadoMCI).Update(models.ResultadosMCI{Valor: Modelo.Valor, FechaModificacion: time.Now(), Autorizado : Modelo.Autorizado})
	
		return updatedResultado, nil	
	 }else{
		 return updatedResultado, errors.New("no autorizado para almacenar")
		}

	
	 //return updatedResultado, nil			
}
