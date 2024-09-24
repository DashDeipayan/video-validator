package database

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/video-validator/constants"
	"github.com/video-validator/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresCreds() string {
	postgresUsername := viper.GetString("DATABASE_USERNAME")
	postgresPassword := viper.GetString("DATABASE_PASSWORD")
	postgresHost := viper.GetString("DATABASE_HOST")
	postgresPort := viper.GetString("DATABASE_PORT")
	postgresDbname := viper.GetString("DATABASE_DBNAME")

	postgresDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		postgresHost,
		postgresUsername,
		postgresPassword,
		postgresDbname,
		postgresPort,
	)

	return postgresDsn
}

func PostgresInit() {
	postgresDsn := postgresCreds()
	db, err := gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})
	if err != nil {
		panic(constants.DB_PANIC)
	}
	db.AutoMigrate(&models.Job{})
}
