package models

type Resultados struct{	
IdResultado 	int  `gorm:"column:idResultado"`	
IdMP 			int  `gorm:"column:idMP"`	
Anio 			int  `gorm:"column:Anio"`	
Mes 			int  `gorm:"column:Mes"`	
Semana 			int  `gorm:"column:Semana"`	
Dia 			int  `gorm:"column:Dia"`	
Valor 			int  `gorm:"column:Valor"`	
}