package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connectionString := "root:masukaja@tcp(0.0.0.0:3306)/kampusmerdeka?parseTime=True"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

// func MigrateDB() {
// 	DB.AutoMigrate(&article{})
// }
