package models

type EstiloTableroModel struct {
	Elemento       string `gorm:"column:Elemento"`
	PropiedadCSSId string `gorm:"column:PropiedadCSSId"`
	Valor          string `gorm:"column:Valor"`
}

type ActualizarEstiloTablero struct {
	Colaborador int
	Estilo      string
}
