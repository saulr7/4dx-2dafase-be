package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func GetResultadosMCI(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	fmt.Println("ResultadosMCIHandler")

	fmt.Println(r)

	vars := mux.Vars(r)
	idMCI := vars["idMCI"]

	var Resultados, erro = services.GetResultadosMCI(idMCI)

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

func ResultadosMCIUpdate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var updatedResultados models.ResultadosMCI

	err := json.NewDecoder(r.Body).Decode(&updatedResultados)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectos")
		return
	}

	var Resultado, err2 = services.ResultadosMCIUpdate(updatedResultados)

	if err2 != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Resultado)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetResultadosGraficaMCI(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	TipoGrafico := vars["TipoGrafico"]
	IdMCI := vars["IdMCI"]
	Anio := vars["Anio"]

	var Resultados, erro = services.GetGraficaMCI(TipoGrafico, IdMCI, Anio)

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

func GetPeriodicidadMCI(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	fmt.Println("PeriodicidadMCIHandler")

	fmt.Println(r)

	var Resultados, erro = services.GetPeriodicidadMCI()

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

func ResultadosMCICreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	fmt.Println("ResultadosMCICreate")

	fmt.Println(r)

	vars := mux.Vars(r)
	idMCI := vars["idMCI"]
	fmt.Println(idMCI)

	var Resultados, erro = services.ResultadosMCICreate(idMCI)

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

func MetaMCIAdd(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var updatedResultados models.ResultadosMCI

	err := json.NewDecoder(r.Body).Decode(&updatedResultados)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectos")
		return
	}

	var Resultado, err2 = services.MetaMCIAdd(updatedResultados)

	if err2 != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Resultado)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}