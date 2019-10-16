package models

import "time"

type Brujula struct {
	IdBrujula       int       `gorm:"AUTO_INCREMENT;column:IdBrujula"`
	IdResultadoMP   int       `gorm:"AUTO_INCREMENT;column:IdResultadoMP"`
	IdColaborador   int       `gorm:"column:IdColaborador"`
	Actividad       string    `gorm:"column:Actividad"`
	IdEstado        int       `gorm:"column:IdEstado"`
	FechaCreada     time.Time `gorm:"column:FechaCreada"`
	FechaModificada time.Time `gorm:"column:FechaModificada"`
}

func (Brujula) TableName() string {
	return "dbBrujulaPorMP"
}

type BrujulaConEstado struct {
	Brujula
	Descripcion string `gorm:"column:Descripcion"`
}

func (BrujulaConEstado) TableName() string {
	return "dbBrujulaPorMP"
}
