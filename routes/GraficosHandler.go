package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
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

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GraficoPorMCINewHandler(w http.ResponseWriter, r *http.Request) {

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

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetGraficoColaborador(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	codigoEmpleado := vars["codigoEmpleado"]

	var Resultados, erro = services.GetGraficoColaborador(codigoEmpleado)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
