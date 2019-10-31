package models

type Colaborador struct {
	IdColaborador int    `gorm:"column:idColaborador"`
	Nombre        string `gorm:"column:NombreColaborador"`
	Correo        string `gorm:"column:Correo"`
	IdPerfil      string `gorm:"column:idPerfil1"`
}

func (Colaborador) TableName() string {
	return "Colaboradores"
}
