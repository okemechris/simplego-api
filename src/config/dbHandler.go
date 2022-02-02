package config

import (
	"simplegoapi/src/domains"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DbConnect() {
	// TODO: connect to postgres
	dsn := "root:@tcp(127.0.0.1:3306)/simple_go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Database = db
	if err != nil {
		panic("failed to connect database")
	}

	runMigrations()

}

func runMigrations() {
	Database.AutoMigrate(&domains.Event{})
}
