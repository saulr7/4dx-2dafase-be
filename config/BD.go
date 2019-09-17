package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  _ "github.com/jinzhu/gorm/dialects/mssql"
)


func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mssql", "sqlserver://app_4dxtablero:app_4dxtablero@des-cobbe01:1433?database=4DX")
	
	if err != nil {
		fmt.Println("Algo salió mal")
		panic(err)
	}

	// defer db.Close()
	db.LogMode(true)
	db.SingularTable(true)
	  
	return db
}