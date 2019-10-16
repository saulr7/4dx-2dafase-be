package models

type Brujula struct {
	IdBrujula 		int  `gorm:"column:idBrujula"`	
	IdMP 			int  `gorm:"column:idMP"`	
	IdColaborador 	int  `gorm:"column:idColaborador"`	
	Actividad 		string  `gorm:"column:Actividad"`	
	Estado 			string  `gorm:"column:idEstado"`	
}
