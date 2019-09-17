package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)

func DBPeriodosPorMPInsert(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	var newdbPeriodos models.DbPeriodoPorMP
	err := json.NewDecoder(r.Body).Decode(&newdbPeriodos)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrecta")
		return
	}

	var dbPeriodo, err2 = services.DbPeriodoPorMPAdd(newdbPeriodos)

	if err2 != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(dbPeriodo)

	fmt.Fprintf(w, string(response))
}
