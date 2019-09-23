package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"../models"
)

func TipoGraficosHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	var graficos, err = services.TipoGraficos()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&graficos)

	fmt.Fprintf(w, string(response))
}

func GraficoPorMCIHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var newGrafico models.GraficoPorMCI
	err := json.NewDecoder(r.Body).Decode(&newGrafico)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var area, err2 = services.GraficoPorMCICreate(newGrafico)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(area)

	fmt.Fprintf(w, string(response))
}
