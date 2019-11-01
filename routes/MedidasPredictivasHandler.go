package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func GetMedidasPredictivas(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Application/json")

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

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetMetasColaborador(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Application/json")

	vars := mux.Vars(r)
	idEmpleado := vars["codigoEmpleado"]

	var usuarioModel, erro = services.GetMetasColaborador(idEmpleado)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetMCI(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Application/json")

	vars := mux.Vars(r)
	idEmpleado := vars["codigoEmpleado"]

	var usuarioModel, erro = services.GetMCIColaborador(idEmpleado)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
