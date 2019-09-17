package services

import (
	"fmt"

	"../config"
	"../models"
)

func AreasGetAll() ([]models.Area, error) {

	var areas []models.Area

	db := config.ConnectDB()
	defer db.Close()

	db.Find(&areas)

	return areas, nil
}

func AreaGetOne(IdArea string) (models.Area, error) {

	var area models.Area

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idArea = ?", IdArea).Find(&area)

	return area, nil
}

func AreaCreate(area models.Area) (models.Area, error) {

	area.ID = 0

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&area)

	return area, nil
}

func AreaUpdate(area models.Area) (models.Area, error) {

	var updatedArea models.Area

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idArea = ?", area.ID).Find(&updatedArea)

	db.Model(&updatedArea).Where("idArea= ?", area.ID).Update(models.Area{Area: area.Area, IdEmpresa: area.IdEmpresa, Activo: area.Activo})

	return updatedArea, nil
}

func AreaDeleted(IdArea string) (bool, error) {

	var area models.Area

	db := config.ConnectDB()
	defer db.Close()

	db.Where("idArea = ? ", IdArea).Delete(&area)

	fmt.Println(area)

	return true, nil
}
