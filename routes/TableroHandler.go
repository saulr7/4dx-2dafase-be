package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func TableroColaborador(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	codigoEmpleado := vars["codigoEmpleado"]
	mesId := vars["mesId"]

	var tablero, err = services.TableroColaborador(codigoEmpleado, mesId)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&tablero)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
