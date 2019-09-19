package models

import("time")

type ResultadosMCI struct{	
	IdResultadoMCI 	int  `gorm:"column:idResultadoMCI"`	
	IdMCI 			int  `gorm:"column:idMCI"`	
	Anio 			int  `gorm:"column:Anio"`	
	Mes 			int  `gorm:"column:Mes"`		
	Valor 			int  `gorm:"column:Valor"`	
	FechaModificacion time.Time   `gorm:"column:FechaModificacion"`
	}