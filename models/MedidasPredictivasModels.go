package models

import (
	_ "github.com/jinzhu/gorm"
)

type GetMPbyMCI struct {
	// gorm.Model
	 IdMCI	        	int     `gorm:"column:idMCI"`
	 IdMP	        	int 	`gorm:"column:idMP"`
	 MedidaPredictiva   string 	`gorm:"column:MedidaPredictiva"`
	 MCI            	string 	`gorm:"column:MCI"`	 	
	 Frecuencia         string  `gorm:"column:Frecuencia"` 	 	 
	 Medicion      		string  `gorm:"column:Medicion"`	
}