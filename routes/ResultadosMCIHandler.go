package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

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