package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
)

func GetFrecuencias(w http.ResponseWriter, r *http.Request) {

	var usuarioModel, erro = services.GetFrecuencias()

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
