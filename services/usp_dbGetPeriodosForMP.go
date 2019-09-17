package services

import (
	
	"fmt"
	"../config"
	"../models"
)

func GetFrecuencias() ([]models.Frecuencia, error) {

	 var result []models.Frecuencia

	 db := config.ConnectDB()
	 defer db.Close()
	 fmt.Println("Entra GetFrecuencias")
	 
	 db.Raw("EXEC usp_dbGetFrecuencias").Scan(&result)	 	

	return result, nil
	
}