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
	Unidad             string   `gorm:"column:Unidad"`
	Autorizado bool   `gorm:"column:Autorizado"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
	Meta float32 `gorm:"column:Meta"`	
}

func (ResultadosMCI) TableName() string {
	return "dbResultadosMCI"
}

type GraficaMCI struct {
	Label 	string  `gorm:"column:Label"`	
    Valor 	float32  `gorm:"column:Valor"`	
}

type PeriodicidadMCI struct{
	IdPeriodicidad    int       `gorm:"column:idPeriodicidad"`
	Descripcion       string       `gorm:"column:Descripcion"`
	MesesPorAnio      string       `gorm:"column:MesesPorAnio"`
	Autorizado      bool       `gorm:"column:Autorizado"`
}
