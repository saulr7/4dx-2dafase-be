package models

import "time"

type Brujula struct {
	IdBrujula       int       `gorm:"AUTO_INCREMENT;column:IdBrujula"`
	IdColaborador   int       `gorm:"column:IdColaborador"`
	Actividad       string    `gorm:"column:Actividad"`
	IdEstado        int       `gorm:"column:IdEstado"`
	FechaCreada     time.Time `gorm:"column:FechaCreada"`
	FechaModificada time.Time `gorm:"column:FechaModificada"`
	Desde           time.Time `gorm:"column:Desde"`
	Hasta           time.Time `gorm:"column:Hasta"`
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