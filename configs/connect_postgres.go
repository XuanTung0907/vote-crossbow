package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = ConnectDB()
	return Db
}

func ConnectDB() *gorm.DB {
	//Connect PostgreSQL
	dsn := "host=localhost user=postgres password=1234 dbname=vote_crossbow port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to DB : error=%v", err)
		return nil
	}
	return db
}
