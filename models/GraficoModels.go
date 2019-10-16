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

type GraficoGlobal struct{
	IdMCI int `gorm:"column:idMCI"`
	MCI string `gorm:"column:MCI"`
	IdMP int `gorm:"column:idMP"`
	MP string `gorm:"column:MP"`
	MetaMCI float32 `gorm:"column:MetaMCI"`
	ResultadoMCI float32 `gorm:"column:ResultadoMCI"`
	MetaMP float32 `gorm:"column:MetaMP"`
	ResultadoMP float32 `gorm:"column:ResultadoMP"`
	MideMeta bool `gorm:"column:MideMeta"`
}