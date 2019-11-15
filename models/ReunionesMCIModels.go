package models

import "time"

type ReunionesMCI struct {
	IdReunion           int       `gorm:"AUTO_INCREMENT;column:idReunion"`
	IdLider             int       `gorm:"column:idLider"`
	FechaInicio         time.Time `gorm:"column:FechaInicio"`
	UltimaActualizacion time.Time `gorm:"column:UltimaActualizacion"`
	TiempoSegundos      int       `gorm:"column:tiempoSegundos"`
}

type ReunionesMCIHora struct {
	ReunionesMCI
	TiempoTotal string `gorm:"column:TiempoTotal"`
}

func (ReunionesMCI) TableName() string {
	return "dbReunionesMCI"
}

type DetalleReunionesMCI struct {
	IdDetalleReunion int       `gorm:"AUTO_INCREMENT;column:IdDetalleReunion"`
	IdReunion        int       `gorm:"column:IdReunion"`
	Colaborador      int       `gorm:"column:Colaborador"`
	Fecha            time.Time `gorm:"column:Fecha"`
	TiempoSegundos   int       `gorm:"column:TiempoSegundos"`
}

func (DetalleReunionesMCI) TableName() string {
	return "dbDetalleReunionesMCI"
}
