package services

import (
	"time"

	"../config"
	"../models"
)

func BrujulaPorMPCreate(Modelo models.Brujula) (models.Brujula, error) {

	Modelo.IdBrujula = 0
	Modelo.FechaCreada = time.Now()
	Modelo.FechaModificada = time.Now()
	Modelo.IdEstado = 1

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func BrujulasPorMPGet(codigoEmpleado string) ([]models.BrujulaConEstado, error) {

	var brujulas []models.BrujulaConEstado

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("SELECT * FROM dbBrujulaPorColaborador b INNER JOIN dbEstadosBrujula e ON e.idEstado = b.idEstado where b.idColaborador = ? order by b.fechaCreada desc", codigoEmpleado).Scan(&brujulas)

	return brujulas, nil
}

func BrujulaEstadoUpdate(Modelo models.Brujula) (models.Brujula, error) {

	var updatedBrujula models.Brujula

	//  if Modelo.Autorizado == true{
	db := config.ConnectDB()
	defer db.Close()

	db.Where("IdBrujula = ?", Modelo.IdBrujula).Find(&updatedBrujula)

	db.Model(&updatedBrujula).Where("IdBrujula= ?", Modelo.IdBrujula).Update(models.Brujula{IdEstado: Modelo.IdEstado, FechaModificada: time.Now()})

	return updatedBrujula, nil
}

func BrujulaEstadoGet() ([]models.BrujulaEstado, error) {

	var BrujulaEstados []models.BrujulaEstado

	db := config.ConnectDB()
	defer db.Close()

	db.Find(&BrujulaEstados)

	return BrujulaEstados, nil
}

func BrujulaActividadesPorMP(idMP string, mes string) ([]models.BrujulaConEstado, error) {

	var BrujulaActividades []models.BrujulaConEstado

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("SELECT * FROM dbBrujulaPorColaborador b INNER JOIN dbEstadosBrujula e on e.idEstado= b.IdEstado  ORDER BY FechaCreada DESC", idMP, mes).Scan(&BrujulaActividades)

	return BrujulaActividades, nil
}
