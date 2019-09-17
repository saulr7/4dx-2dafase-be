package models

import (
	_ "github.com/jinzhu/gorm"
)

type Area struct {
	// gorm.Model
	ID        int    `gorm:"AUTO_INCREMENT;column:idArea"`
	Area      string `gorm:"column:Area"`
	IdEmpresa string `gorm:"column:idEmpresa"`
	Activo    bool   `gorm:"column:Activo"`
}

func (Area) TableName() string {
	return "Areas"
}
