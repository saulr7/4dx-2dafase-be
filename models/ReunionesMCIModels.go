package models

import "time"

type ReunionesMCI struct {
	IdReunion           int       `gorm:"AUTO_INCREMENT;column:IdReunion"`
	IdLider             int       `gorm:"column:IdLider"`
	FechaInicio         time.Time `gorm:"column:FechaInicio"`
	UltimaActualizacion time.Time `gorm:"column:UltimaActualizacion"`
}

type DetalleReunionesMCI struct {
	IdDetalleReunion int    `gorm:"AUTO_INCREMENT;column:IdDetalleReunion"`
	IdReunion        int    `gorm:"column:IdReunion"`
	Colaborador      int    `gorm:"column:Colaborador"`
	HoraInicio       string `gorm:"column:HoraInicio"`
	Horafin          string `gorm:"column:Horafin"`
}
