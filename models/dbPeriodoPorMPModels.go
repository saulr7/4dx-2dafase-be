package models

import (
	_ "github.com/jinzhu/gorm"
)

type DbPeriodoPorMP struct {
	// gorm.Model
	IdPeriodo    int    `gorm:"AUTO_INCREMENT;column:idPeriodo"`
	IdMP         string `gorm:"column:idMP"`
	IdFrecuencia string `gorm:"column:idFrecuencia"`
	IdMedicion   bool   `gorm:"column:idMedicion"`
}
