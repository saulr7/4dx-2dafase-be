package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)

func RegistrarEventoDelSistema(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var newEvento models.EventoDelSistema
	err := json.NewDecoder(r.Body).Decode(&newEvento)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var evento, err2 = services.RegistarEventoDelSistema(newEvento)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(evento)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
