package models

type ResultadosMCI struct{	
	IdResultadoMCI 	int  `gorm:"column:idResultadoMCI"`	
	IdMCI 			int  `gorm:"column:idMCI"`	
	Anio 			int  `gorm:"column:Anio"`	
	Mes 			int  `gorm:"column:Mes"`		
	Valor 			int  `gorm:"column:Valor"`	
	}