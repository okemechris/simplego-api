package config

import (
	"fmt"
	"os"
	"simplegoapi/src/domains"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

var (
	user         = os.Getenv("KEYCLOAK_DATABASE_USER")
	password     = os.Getenv("KEYCLOAK_DATABASE_NAME")
	databaseName = os.Getenv("KEYCLOAK_DATABASE_PASSWORD")
	hostname     = os.Getenv("KEYCLOAK_DATABASE_HOST")
	port         = 5433
	// TODO: accept db port as .env
	// TODO: allow multiple db vendor setup
	// vendor       = os.Getenv("KEYCLOAK_DATABASE_VENDOR")
)

func DbConnect() {
	dbURL := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s`, user, password, hostname, port, databaseName)
	fmt.Println("connecting to: ", dbURL)
	// dsn := "root:@tcp(127.0.0.1:3306)/simple_go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	Database = db
	if err != nil {
		panic("failed to connect database")
	}

	runMigrations()

}

func runMigrations() {
	fmt.Println("running migrations...")
	Database.AutoMigrate(&domains.Event{})
}
