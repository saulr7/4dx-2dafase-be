package models

type Brujula struct {
	IdBrujula     int    `gorm:"column:idBrujula"`
	IdMP          int    `gorm:"column:idMP"`
	IdColaborador int    `gorm:"column:idColaborador"`
	Actividad     string `gorm:"column:Actividad"`
	IdEstado      int    `gorm:"column:idEstado"`
}

func (Brujula) TableName() string {
	return "dbBrujulaPorMP"
}
