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

func SendEmailForAutorization(model models.Email) (bool, error) {

	model.ProfileName = "4DX"

	colaboradores, _ := ColaboradoresAdmins()

	var remitentes = ""

	for _, colaborador := range colaboradores {
		remitentes = remitentes + colaborador.Correo + ";"
	}

	colaborador, _ := ColaboradorPorCodigo(model.ColaboradorId)

	remitentes = remitentes + colaborador.Correo

	// emailModel.Recipients = remitentes
	model.Recipients = "saulr@banpais.hn"
	model.Subject = model.Subject + " " + remitentes

	var sended, err2 = SendEmail(model)

	type Result struct {
		Send bool `gorm:"column:Send"`
	}
	var result Result

	if err2 != nil {
		return result.Send, err2
	}

	return sended, nil

}
