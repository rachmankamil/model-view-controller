package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rachmankamil/kampus-merdeka-b/models"
)

var DB *gorm.DB

const JWT_SECRET string = "testmvc"

const JWT_EXP int = 1

func InitDB(status string) {
	db := "kampusmerdeka"
	if status == "testing" {
		db = "kampusmerdeka-test"
	}
	connectionString := fmt.Sprintf("root:masukaja@tcp(0.0.0.0:3306)/%s?parseTime=True", db)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.Article{})
}
