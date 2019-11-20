package models

import "time"

type Brujula struct {
	IdBrujula          int       `gorm:"AUTO_INCREMENT;column:IdBrujula"`
	IdColaborador      int       `gorm:"column:IdColaborador"`
	Actividad          string    `gorm:"column:Actividad"`
	IdEstado           int       `gorm:"column:IdEstado"`
	FechaCreada        time.Time `gorm:"column:FechaCreada"`
	FechaModificada    time.Time `gorm:"column:FechaModificada"`
	Desde              time.Time `gorm:"column:Desde"`
	Hasta              time.Time `gorm:"column:Hasta"`
	ActividadComoLider bool      `gorm:"column:ActividadComoLider"`
	CreatedBy          int       `gorm:"column:CreatedBy"`
}

type BrujulaLista struct {
	Brujula
	Actividades []string
}

func (Brujula) TableName() string {
	return "dbBrujulaPorColaborador"
}

type BrujulaConEstado struct {
	Brujula
	Descripcion string `gorm:"column:Descripcion"`
}

func (BrujulaConEstado) TableName() string {
	return "dbBrujulaPorColaborador"
}

type BrujulaEstado struct {
	IdEstado    int    `gorm:"column:idEstado;"`
	Descripcion string `gorm:"column:Descripcion"`
}

type BrujulaCantidad struct {
	IdColaborador     int `gorm:"column:idColaborador;"`
	Cantidad          int `gorm:"column:Cantidad"`
	CantidadComoLider int `gorm:"column:CantidadComoLider"`
}

func (BrujulaEstado) TableName() string {
	return "dbEstadosBrujula"
}
