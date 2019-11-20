package services

import (
	"fmt"
	"time"

	"../config"
	"../models"
)

func BrujulaPorMPCreate(Modelo models.BrujulaLista) (models.BrujulaLista, error) {

	fmt.Println(Modelo.Actividades)

	var brujulaModel models.Brujula
	brujulaModel.IdBrujula = 0
	brujulaModel.FechaCreada = time.Now()
	brujulaModel.FechaModificada = time.Now()
	brujulaModel.Desde = time.Now()
	brujulaModel.Hasta = time.Now()
	brujulaModel.IdEstado = 1
	brujulaModel.IdColaborador = Modelo.IdColaborador
	brujulaModel.CreatedBy = Modelo.CreatedBy
	brujulaModel.ActividadComoLider = Modelo.ActividadComoLider

	db := config.ConnectDB()
	defer db.Close()

	for _, brujula := range Modelo.Actividades {
		brujulaModel.IdBrujula = 0
		brujulaModel.Actividad = brujula
		db.Create(&brujulaModel)
	}

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

func GetFBrujulaCantidad(idColaborador string) ([]models.BrujulaCantidad, error) {

	var result []models.BrujulaCantidad

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetActiviadesSiguienteBrujula ?", idColaborador).Scan(&result)

	return result, nil

}

func BrujulaActividadesPorMP(idMP string, mes string) ([]models.BrujulaConEstado, error) {

	var BrujulaActividades []models.BrujulaConEstado

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("SELECT * FROM dbBrujulaPorColaborador b INNER JOIN dbEstadosBrujula e on e.idEstado= b.IdEstado  ORDER BY FechaCreada DESC", idMP, mes).Scan(&BrujulaActividades)

	return BrujulaActividades, nil
}

func BrujulaActividadesPorColaborador(idEmpleado string, idEstado string, esLider string) ([]models.BrujulaConEstado, error) {

	var BrujulaActividades []models.BrujulaConEstado

	filtroPorEstado := ""

	if idEstado != "NO" {
		filtroPorEstado = "AND b.idEstado = " + idEstado
	}

	if esLider == "SI" {
		filtroPorEstado = filtroPorEstado + " AND (ActividadComoLider is null or ActividadComoLider <> 1) "
	}

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("SELECT * FROM dbBrujulaPorColaborador b inner join dbEstadosBrujula e on b.IdEstado = e.idEstado where IdColaborador = ? "+filtroPorEstado+" order by FechaCreada desc", idEmpleado).Scan(&BrujulaActividades)

	return BrujulaActividades, nil
}
