package models

import (
	"time"
)

type ResultadosMCI struct {
	IdResultadoMCI    int       `gorm:"column:idResultadoMCI"`
	IdMCI             int       `gorm:"column:idMCI"`
	MCI             string       `gorm:"column:MCI"`
	Anio              int       `gorm:"column:Anio"`
	Mes               int       `gorm:"column:Mes"`
	Valor             float32   `gorm:"column:Valor"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
}

func (ResultadosMCI) TableName() string {
	return "dbResultadosMCI"
}
