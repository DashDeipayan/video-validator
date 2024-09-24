package database

import (
	"log/slog"

	"github.com/spf13/viper"
	"github.com/video-validator/constants"
	"github.com/video-validator/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteInit() {
	dbName := viper.GetString("DATABASE_DBNAME")
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		slog.Error("error", slog.Any("err", err))
		panic(constants.DB_PANIC)
	}
	db.AutoMigrate(&models.Job{})
}
