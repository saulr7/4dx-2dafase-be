package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)
func BrujulaPorMPCreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var newBrujula models.Brujula
	err := json.NewDecoder(r.Body).Decode(&newBrujula)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var Brujula, err2 = services.BrujulaPorMPCreate(newBrujula)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujula)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}