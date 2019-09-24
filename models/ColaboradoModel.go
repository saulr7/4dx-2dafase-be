package models

type Colaborador struct {
	IdColaborador int    `gorm:"column:idColaborador"`
	Nombre        string `gorm:"column:NombreColaborador"`
}

func (Colaborador) TableName() string {
	return "Colaboradores"
}
