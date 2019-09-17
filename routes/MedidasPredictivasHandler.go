package routes

import (	
	"encoding/json"
	"fmt"
	"net/http"

	//"../models"
	"../services"
	"github.com/gorilla/mux"
)

func GetMedidasPredictivas(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")	
	
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	fmt.Println("GetMedidasPredictivas")

	fmt.Println(r)

	vars := mux.Vars(r)
	idEmpleado := vars["codigoEmpleado"]

	var usuarioModel, erro = services.GetMedidasPredictivas(idEmpleado)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	fmt.Fprintf(w, string(response))
}
