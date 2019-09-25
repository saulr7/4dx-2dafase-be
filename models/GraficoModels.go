package models

type GraficoPorMCI struct {
	IdRelacion    int `gorm:"column:IdRelacion"`
	IdMCI         int `gorm:"column:IdMCI"`
	IdTipoGrafico int `gorm:"column:IdTipoGrafico"`
}

type TipoGrafico struct {
	Id          int    `gorm:"column:ID"`
	Nombre      string `gorm:"column:Nombre"`
	TIMESTAMP      string `gorm:"column:TIMESTAMP"`
	Descripcion string `gorm:"column:Descripcion"`
	Activo      bool   `gorm:"column:Activo"`
}

func (TipoGrafico) TableName() string {
	return "dbTipoGraficos"
}
