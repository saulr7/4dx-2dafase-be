package models

type Mediciones struct {
	IdPeriodo 	int  `gorm:"column:idMedicion"`	
    Descripcion string  `gorm:"column:Medicion"`	
}
