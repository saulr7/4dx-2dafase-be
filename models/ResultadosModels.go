package models

import (
	"time"
)

type Resultados struct {
	IdResultado       int       `gorm:"column:idResultado"`
	IdMP              int       `gorm:"column:idMP"`
	IdFrecuencia              int       `gorm:"column:idFrecuencia"`
	IdMedicion              int       `gorm:"column:idMedicion"`
	Anio              int       `gorm:"column:Anio"`
	Mes               int       `gorm:"column:Mes"`
	Semana            int       `gorm:"column:Semana"`
	Dia               int       `gorm:"column:Dia"`
	Valor             float32   `gorm:"column:Valor"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
}

func (Resultados) TableName() string {
	return "dbResultados"
}
