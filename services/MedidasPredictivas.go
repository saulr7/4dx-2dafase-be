package services

import (
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
