package models

type TipoGrafico struct {
	Id          int    `gorm:"column:ID"`
	Nombre      string `gorm:"column:Nombre"`
	Descripcion string `gorm:"column:Descripcion"`
	Activo      bool   `gorm:"column:Activo"`
}

func (TipoGrafico) TableName() string {
	return "TipoGraficos"
}
