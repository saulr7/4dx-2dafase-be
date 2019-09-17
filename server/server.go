package server

import (
	"net/http"

	"../routes"
	"../services"
	"github.com/gorilla/handlers"
	// "github.com/gorilla/handlers"
)

func Serve() {

	myRouter := routes.Routes()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.Handle("/", services.IsLogginMiddleWare(myRouter))

	// http.Handle("/", myRouter)
	http.ListenAndServe(":8087", handlers.CORS(originsOk, headersOk, methodsOk)(services.IsLogginMiddleWare(myRouter)))
	// http.ListenAndServe(":8087", nil)
}
