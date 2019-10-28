package models

type Colaborador struct {
	IdColaborador int    `gorm:"column:idColaborador"`
	Nombre        string `gorm:"column:NombreColaborador"`
	Correo        string `gorm:"column:Correo"`
}

func (Colaborador) TableName() string {
	return "Colaboradores"
}
