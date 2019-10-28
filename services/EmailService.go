package services

import (
	"../config"
	"../models"
)

func SendEmail(model models.Email) (bool, error) {

	type Result struct {
		Send bool `gorm:"column:Send"`
	}

	var result Result

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_SendEmail ?, ?, ?, ? ", model.ProfileName, model.Recipients, model.Subject, model.Body).Scan(&result)

	return result.Send, nil

}
