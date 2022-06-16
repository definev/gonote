package database

import (
	"fmt"
	"strconv"

	"github.com/definev/gonote/config"
	"github.com/definev/gonote/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Idiot, you need to set the DB_PORT environment variable to a valid port number")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_NAME"), config.Config("DB_PASSWORD"))

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("Error connecting to the database")
	} else {
		fmt.Println("Connected to the database")
	}

	err = DB.AutoMigrate(&model.Note{})
	if err != nil {
		fmt.Println("Error auto migrating the database")
	} else {
		fmt.Println("Database Migration complete")
	}
}
