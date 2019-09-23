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

	fmt.Fprintf(w, string(response))
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

	fmt.Fprintf(w, string(response))
}



func GetResultadosGraficaMCI(w http.ResponseWriter, r *http.Request) {
	//fmt.println("GetResultadosGraficaMCI handler")
	//fmt.Println("Prueba");
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	TipoGrafico := vars["TipoGrafico"]
	IdMCI := vars["IdMCI"]
	Anio := vars["Anio"]

	fmt.Println("variables");
	//fmt.Println(TipoGrafico);
	//fmt.Println(IdMCI);
	fmt.Println(Anio);

	var Resultados, erro = services.GetGraficaMCI(TipoGrafico, IdMCI, Anio)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprintf(w, string(response))
}
