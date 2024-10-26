package Config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todo-app/Model"
)

var DB *gorm.DB

func DatabaseConnection() {
	env := loadEnv()
	sqlInfo := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		env.DBUser, env.DBPass, env.DBPort, env.DBName)

	var err error
	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&Model.Todo{})
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database")
}
