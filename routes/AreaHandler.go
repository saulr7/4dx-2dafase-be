package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func AreaGetAllHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	var areas, err = services.AreasGetAll()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&areas)

	fmt.Fprintf(w, string(response))
}

func AreaOneHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	IdArea := vars["key"]

	fmt.Println(IdArea)

	var area, err = services.AreaGetOne(IdArea)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&area)

	fmt.Fprintf(w, string(response))
}

func AreaCreateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	var newArea models.Area
	err := json.NewDecoder(r.Body).Decode(&newArea)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Credenciales incorrectas")
		return
	}

	var area, err2 = services.AreaCreate(newArea)

	if err2 != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(area)

	fmt.Fprintf(w, string(response))
}

func AreaUpdateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	var updatedArea models.Area
	err := json.NewDecoder(r.Body).Decode(&updatedArea)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectos")
		return
	}

	var area, err2 = services.AreaUpdate(updatedArea)

	if err2 != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(area)

	fmt.Fprintf(w, string(response))
}

func AreaDeleteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	IdArea := vars["key"]

	fmt.Println(IdArea)

	var seElemino, err = services.AreaDeleted(IdArea)

	if err != nil || !seElemino {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido eliminar la data")
		return
	}

	fmt.Fprintf(w, "Se ha eliminado")
}
