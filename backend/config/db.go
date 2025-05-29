package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:k@l3nd@riumdev@tcp(172.17.0.4:3306)/mariadb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal()
	}
	DB = db
	//defer sqlDB.Close() // defer function wykona się po zakończeniu main()
}
