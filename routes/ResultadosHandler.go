package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func GetResultados(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	idMP := vars["idMP"]

	var Resultados, erro = services.GetResultados(idMP)

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

func ResultadosUpdate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var updatedResultados models.Resultados

	err := json.NewDecoder(r.Body).Decode(&updatedResultados)
	fmt.Println(updatedResultados)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectos")
		return
	}

	var Resultado, err2 = services.ResultadosUpdate(updatedResultados)

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
