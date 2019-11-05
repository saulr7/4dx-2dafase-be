package models

import "time"

type Brujula struct {
<<<<<<< HEAD
	IdBrujula       int       `gorm:"AUTO_INCREMENT;column:IdBrujula"`
	IdColaborador   int       `gorm:"column:IdColaborador"`
	Actividad       string    `gorm:"column:Actividad"`
	IdEstado        int       `gorm:"column:IdEstado"`
	FechaCreada     time.Time `gorm:"column:FechaCreada"`
	FechaModificada time.Time `gorm:"column:FechaModificada"`
	Desde           time.Time `gorm:"column:Desde"`
	Hasta           time.Time `gorm:"column:Hasta"`
	ActividadComoLider bool `gorm:"column:ActividadComoLider"`
	CreatedBy int `gorm:"column:CreatedBy"`
=======
	IdBrujula          int       `gorm:"AUTO_INCREMENT;column:IdBrujula"`
	IdColaborador      int       `gorm:"column:IdColaborador"`
	Actividad          string    `gorm:"column:Actividad"`
	IdEstado           int       `gorm:"column:IdEstado"`
	FechaCreada        time.Time `gorm:"column:FechaCreada"`
	FechaModificada    time.Time `gorm:"column:FechaModificada"`
	Desde              time.Time `gorm:"column:Desde"`
	Hasta              time.Time `gorm:"column:Hasta"`
	ActividadComoLider bool      `gorm:"column:ActividadComoLider"`
>>>>>>> 83bc26b4b204f60a5faee4f98ceca48fa1e2cb70
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

func (BrujulaEstado) TableName() string {
	return "dbEstadosBrujula"
}
