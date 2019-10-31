package services

import (
	"fmt"
	"time"

	"../config"
	"../models"
)

func ReunionMCIreate(Modelo models.ReunionesMCI) (models.ReunionesMCI, error) {

	Modelo.IdReunion = 0
	Modelo.FechaInicio = time.Now()
	Modelo.UltimaActualizacion = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func SetDetalleReunionMCI(Modelo models.DetalleReunionesMCI) (models.DetalleReunionesMCI, error) {

	Modelo.IdDetalleReunion = 0
	Modelo.Fecha = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func GetReunionDelDia(empleadoId string) (models.ReunionesMCI, error) {

	var result models.ReunionesMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("SELECT top 1 * FROM dbReunionesMCI where idLider = ? and cast(FechaInicio as date) = cast(GETDATE() as date) order by 1 desc ", empleadoId).Scan(&result)

	fmt.Println(result)

	return result, nil
}

func UpdateTiempoReunion(Model models.ReunionesMCI) (models.ReunionesMCI, error) {

	var updateReunion models.ReunionesMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idReunion = ?", Model.IdReunion).Find(&updateReunion)

	db.Model(&updateReunion).Where("idReunion= ?", Model.IdReunion).Update(map[string]interface{}{"UltimaActualizacion": time.Now(), "tiempoSegundos": Model.TiempoSegundos})

	return updateReunion, nil

}
