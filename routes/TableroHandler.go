package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"

	"../services"
	"github.com/gorilla/mux"
)

func TableroColaborador(w http.ResponseWriter, r *http.Request) {

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

func EstiloTableroHanlder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	codigoEmpleado := vars["idColaborador"]

	var tablero, err = services.GetEstiloTablero(codigoEmpleado)

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

func UpdateEstiloTableroHanlder(w http.ResponseWriter, r *http.Request) {

	token, _ := services.GetToken(r)

	var newEstilo models.ActualizarEstiloTablero
	err2 := json.NewDecoder(r.Body).Decode(&newEstilo)

	if err2 != nil {
		fmt.Println(err2)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrecta")
		return
	}

	newEstilo.Colaborador = token.Usuario.Empleado

	var tablero, err = services.UpdateEstiloTablero(newEstilo)

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
