package models

type BrujulaEstado struct {
	IdEstado    int    `gorm:"column:idEstado;"`
	Descripcion string `gorm:"column:Descripcion"`
}

func (BrujulaEstado) TableName() string {
	return "dbEstadosBrujula"
}
