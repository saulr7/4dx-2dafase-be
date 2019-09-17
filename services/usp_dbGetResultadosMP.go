package services

import (
	
	"fmt"
	"../config"
	"../models"
)

func GetResultados(idMP string) ([]models.Resultados, error) {

	 var result []models.Resultados

	 db := config.ConnectDB()
	 defer db.Close()
	 fmt.Println("Entra usp_dbGetResultadosMP")
	 
	 db.Raw("EXEC usp_dbGetResultadosMP ?",idMP).Scan(&result)	 	

	return result, nil
	
}