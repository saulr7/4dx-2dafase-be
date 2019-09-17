package routes

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/login", LoginHandler)
	myRouter.HandleFunc("/home", HomeHandler).Methods("GET")
	myRouter.HandleFunc("/areas", AreaGetAllHandler).Methods("GET")
	myRouter.HandleFunc("/area/{key}", AreaOneHandler).Methods("GET")
	myRouter.HandleFunc("/areaNew", AreaCreateHandler).Methods("POST")
	myRouter.HandleFunc("/areaUpdate", AreaUpdateHandler).Methods("POST")
	myRouter.HandleFunc("/areaDelete/{key}", AreaDeleteHandler).Methods("GET")
	myRouter.HandleFunc("/GetMedidasPredictivas/{codigoEmpleado}", GetMedidasPredictivas).Methods("GET")
	myRouter.HandleFunc("/GetFrecuencias", GetFrecuencias)
	myRouter.HandleFunc("/GetMediciones", GetMediciones)
	myRouter.HandleFunc("/GetResultado/{idMP}", GetResultados)

	return myRouter
}
