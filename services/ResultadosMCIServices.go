package services

import (
	"fmt"

	"../config"
	"../models"

	"time"
)

func GetResultadosMCI(idMCI string) ([]models.ResultadosMCI, error) {

	var result []models.ResultadosMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbGetResultadosMCI ?", idMCI).Scan(&result)

	return result, nil
}

func ValorMCIAdd(Modelo models.ResultadosMCI) (models.ResultadosMCI, error) {

	var updatedResultado models.ResultadosMCI

	//  if Modelo.Autorizado == true{
	db := config.ConnectDB()
	defer db.Close()

	fmt.Println("Llego")

	db.Where("idResultadoMCI = ?", Modelo.IdResultadoMCI).Find(&updatedResultado)

	db.Model(&updatedResultado).Where("idResultadoMCI= ?", Modelo.IdResultadoMCI).Update(models.ResultadosMCI{Valor: Modelo.Valor, FechaModificacion: time.Now()})

	return updatedResultado, nil
}

func ResultadosMCICreate(idMCI string) ([]models.ResultadosMCI, error) {

	var result []models.ResultadosMCI

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbCreaResultadosMCI ?", idMCI).Scan(&result)

	return result, nil
}

func MetaMCIAdd(Modelo models.ResultadosMCI) (models.ResultadosMCI, error) {

	var updatedResultado models.ResultadosMCI

	//  if Modelo.Autorizado == true{
	db := config.ConnectDB()
	defer db.Close()

	db.Where("idResultadoMCI = ?", Modelo.IdResultadoMCI).Find(&updatedResultado)

	db.Model(&updatedResultado).Where("idResultadoMCI= ?", Modelo.IdResultadoMCI).Update(models.ResultadosMCI{Meta: Modelo.Meta, FechaModificacion: time.Now()})

	return updatedResultado, nil
}
