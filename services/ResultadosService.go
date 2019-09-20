package services

import (
	"../config"
	"../models"

	"time"
	"errors"
)

func GetResultados(idMP string) ([]models.Resultados, error) {

	var result []models.Resultados

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetResultadosMP ?", idMP).Scan(&result)

	return result, nil
}


func ResultadosUpdate(Modelo models.Resultados)(models.Resultados, error){

	var updatedResultado models.Resultados

	if Modelo.Autorizado == true{

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idResultado = ?", Modelo.IdResultado).Find(&updatedResultado)

	db.Model(&updatedResultado).Where("idResultado= ?", Modelo.IdResultado).Update(models.Resultados{Valor: Modelo.Valor, FechaModificacion: time.Now(), LlegoAMeta: Modelo.LlegoAMeta})
	return updatedResultado, nil	
}else{
	return updatedResultado, errors.New("No autorizado para almacenar")	
}
	
	
}