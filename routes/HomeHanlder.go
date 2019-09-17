package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	type Data struct {
		Name string `json:"name"`
	}

	var data Data

	data.Name = "Saul"
	response, _ := json.Marshal(&data)

	fmt.Fprintf(w, string(response))
}
