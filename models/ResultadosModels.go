package models

import (
	"time"
)

type Resultados struct {
	IdResultado       int       `gorm:"column:idResultado"`
	IdMP              int       `gorm:"column:idMP"`
	MedidaPredictiva  string    `gorm:"column:MedidaPredictiva"`
	IdFrecuencia      int       `gorm:"column:idFrecuencia"`
	IdMedicion        int       `gorm:"column:idMedicion"`
	IdColaborador     int       `gorm:"column:idColaborador"`
	Anio              int       `gorm:"column:Anio"`
	Mes               int       `gorm:"column:Mes"`
	Semana            int       `gorm:"column:Semana"`
	Dia               int       `gorm:"column:Dia"`
	Valor             float32   `gorm:"column:Valor"`
	Unidad            string    `gorm:"column:Unidad"`
	LlegoAMeta        bool      `gorm:"column:LlegoAMeta"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
	Autorizado        bool      `gorm:"column:Autorizado"`
	Meta              string    `gorm:"column:Meta"`
	Periodo           string    `gorm:"column:Periodo"`
}

func (Resultados) TableName() string {
	return "dbResultados"
}
