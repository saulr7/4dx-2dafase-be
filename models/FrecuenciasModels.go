package models

type Frecuencia struct {
	IdPeriodo 	int  `gorm:"column:idPeriodo"`	
    Descripcion string  `gorm:"column:Descripcion"`	
}
