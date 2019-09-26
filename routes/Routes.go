package routes

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/login", LoginHandler)
	myRouter.HandleFunc("/GetMedidasPredictivas/{codigoEmpleado}", GetMedidasPredictivas).Methods("GET")
	myRouter.HandleFunc("/GetFrecuencias", GetFrecuencias)
	myRouter.HandleFunc("/GetMediciones", GetMediciones)
	myRouter.HandleFunc("/GetResultados/{idMP}", GetResultados).Methods("GET")
	myRouter.HandleFunc("/GetResultadosMCI/{idMCI}", GetResultadosMCI).Methods("GET")
	myRouter.HandleFunc("/PeriodosPorMPNew", DBPeriodosPorMPInsert).Methods("POST")
	myRouter.HandleFunc("/PeriodosPorMPUpdate", DBPeriodosPorMPUpdateHandler).Methods("POST")
	myRouter.HandleFunc("/ResultadosUpdate", ResultadosUpdate).Methods("POST")
	myRouter.HandleFunc("/ValorMCIAdd", ValorMCIAdd).Methods("POST")
	myRouter.HandleFunc("/TipoGraficos", TipoGraficosHandler).Methods("GET")
	myRouter.HandleFunc("/GetPeriodicidadMCI", GetPeriodicidadMCI)
	myRouter.HandleFunc("/GetResultadosGraficaMCI/{IdMCI}/{Anio}", GetResultadosGraficaMCI).Methods("GET")
	myRouter.HandleFunc("/GraficoPorMCINew", GraficoPorMCINewHandler).Methods("POST")
	myRouter.HandleFunc("/SubAreas", SubAreas).Methods("GET")
	myRouter.HandleFunc("/ColaboradoresPorArea/{AreaId}", ColaboradoresPorArea).Methods("GET")
	myRouter.HandleFunc("/PeriodosPorMCIAdd", PeriodosPorMCIAdd).Methods("POST")
	myRouter.HandleFunc("/ResultadosMCICreate/{idMCI}", ResultadosMCICreate).Methods("GET")
	myRouter.HandleFunc("/GetGraficoColaborador/{codigoEmpleado}", GetGraficoColaborador).Methods("GET")

	return myRouter
}
