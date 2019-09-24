package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
)

func GetMediciones(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	var usuarioModel, erro = services.GetMediciones()

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
