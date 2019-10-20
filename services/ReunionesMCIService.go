package services

import (
	"time"

	"../config"
	"../models"
)

func ReunionMCIreate(Modelo models.ReunionesMCI) (models.ReunionesMCI, error) {

	Modelo.IdReunion = 0
	Modelo.FechaInicio = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func SetDetalleReunionMCI(IdReunion string, IdColaborador string, HoraInicio string, horaFin string) ([]models.DetalleReunionesMCI, error) {

	var result []models.DetalleReunionesMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC dbo.usp_dbSetDetalleReuniones ?, ?, ?, ?", IdReunion, IdColaborador, HoraInicio, horaFin).Scan(&result)

	return result, nil
}
