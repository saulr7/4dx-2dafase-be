package models

type GraficoPorMCI struct {
	IdRelacion    int `gorm:"column:IdRelacion"`
	IdMCI         int `gorm:"column:IdMCI"`
	IdTipoGrafico int `gorm:"column:IdTipoGrafico"`
}
