package models

import (
	_ "github.com/jinzhu/gorm"
)

type DbPeriodoPorMP struct {
	// gorm.Model
	IdPeriodo    int `gorm:"AUTO_INCREMENT;column:idPeriodo"`
	IdMP         int `gorm:"column:idMP"`
	IdFrecuencia int `gorm:"column:idFrecuencia"`
	IdMedicion   int `gorm:"column:idMedicion"`
}

func (DbPeriodoPorMP) TableName() string {
	return "dbPeriodoPorMP"
}

type PeriodoPorMCI struct {
	// gorm.Model
	IdPeriodo    int `gorm:"AUTO_INCREMENT;column:idPeriodo"`
<<<<<<< HEAD
	IdMCI         int `gorm:"column:idMCI"`
	IdFrecuencia int `gorm:"column:idFrecuencia"`	
	Meta float32 `gorm:"column:Meta"`	
}
=======
	IdMCI        int `gorm:"column:idMCI"`
	IdFrecuencia int `gorm:"column:idFrecuencia"`
}

func (PeriodoPorMCI) TableName() string {
	return "dbPeriodoPorMCI"
}
>>>>>>> 9576035d11bf29e4489fb298b999018294bd25ea
