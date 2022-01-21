package database

import (
	"adminProject/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Product{})
}
