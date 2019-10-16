package services

import (
	"../config"
	"../models"
)

func BrujulaPorMPCreate(Modelo models.Brujula) (models.Brujula, error) {

	Modelo.IdBrujula = 0

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}
