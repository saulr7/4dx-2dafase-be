package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	var emailModel models.Email

	err := json.NewDecoder(r.Body).Decode(&emailModel)
	emailModel.ProfileName = "4DX"

	colaboradores, _ := services.ColaboradoresAdmins()

	var remitentes = ""

	for _, colaborador := range colaboradores {
		remitentes = remitentes + colaborador.Correo + ";"
	}

	emailModel.Recipients = "saulr@banpais.hn"
	emailModel.Subject = emailModel.Subject + " " + remitentes
	// emailModel.Recipients = remitentes

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	var sended, err2 = services.SendEmail(emailModel)

	if err2 != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error enviando email")
		return
	}

	response, _ := json.Marshal(&sended)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
