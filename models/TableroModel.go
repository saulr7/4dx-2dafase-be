package models

type Tablero struct {
	IdMCI            int    `gorm:"column:idMCI"`
	MCI              string `gorm:"column:MCI"`
	ResultadosMCI    []ResultadoMCI
	MedidaPredictiva []MedidaPredictiva
}

type ResultadoMCI struct {
	Meta         float32 `gorm:"column:Meta"`
	ResultadoMCI float32 `gorm:"column:Valor"`
	Mes          int     `gorm:"column:Mes"`
}

type MedidaPredictiva struct {
	IdMP             int     `gorm:"column:idMP"`
	MedidaPredictiva string  `gorm:"column:MedidaPredictiva"`
	MetaMP           float32 `gorm:"column:Meta"`
	ResultadosMP     []ResultadoMP
}

type ResultadoMP struct {
	Meta   float32 `gorm:"column:Meta"`
	Valor  float32 `gorm:"column:Valor"`
	Unidad string  `gorm:"column:Unidad"`
}
