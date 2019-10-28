package models

import "time"

//	_"github.com/jinzhu/gorm"

type DbPeriodoPorMP struct {
	// gorm.Model
	IdPeriodo     int       `gorm:"AUTO_INCREMENT;column:idPeriodo"`
	IdMP          int       `gorm:"column:idMP"`
	IdFrecuencia  int       `gorm:"column:idFrecuencia"`
	IdMedicion    int       `gorm:"column:idMedicion"`
	Anio          int       `gorm:"column:Anio"`
	FechaCreacion time.Time `gorm:"column:FechaCreacion"`
}

func (DbPeriodoPorMP) TableName() string {
	return "dbPeriodoPorMP"
}

type PeriodoPorMCI struct {
	// gorm.Model
	IdPeriodo    int `gorm:"AUTO_INCREMENT;column:idPeriodo"`
	IdMCI        int `gorm:"column:idMCI"`
	IdFrecuencia int `gorm:"column:idFrecuencia"`
	Anio         int `gorm:"column:Anio"`
}

func (PeriodoPorMCI) TableName() string {
	return "dbPeriodoPorMCI"
}
