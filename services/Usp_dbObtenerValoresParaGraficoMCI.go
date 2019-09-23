package services

import (
	
	"fmt"
	"../config"
	"../models"
)

func GetGraficaMCI(TipoGrafico string, IdMCI string, Anio string) ([]models.GraficaMCI, error) {

	var result []models.GraficaMCI

	db := config.ConnectDB()
	defer db.Close()
	fmt.Println("Entra GraficaMCI Servicio")
	
	db.Raw("EXEC dbo.usp_dbObtenerValoresParaGraficoMCI ?, ?, ?", TipoGrafico, IdMCI, Anio ).Scan(&result)

   return result, nil
}