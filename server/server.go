package server

import (
	"fmt"
	"net/http"

	"../routes"
	"../services"
)

func Serve() {

	myRouter := routes.Routes()

	http.Handle("/", CORS(services.IsLogginMiddleWare(myRouter)))

	http.ListenAndServe(":8087", nil)
}

func CORS(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			fmt.Println("If")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization, X-Requested-With, Content-Length, Accept-Encoding,Screen")
			w.WriteHeader(http.StatusOK)
			return
		} else {
			fmt.Println("Else")
			h.ServeHTTP(w, r)
		}
	})
}
