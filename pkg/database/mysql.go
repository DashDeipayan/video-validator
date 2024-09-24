package database

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/video-validator/constants"
	"github.com/video-validator/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLCreds() string {
	mysqlUsername := viper.GetString("DATABASE_USERNAME")
	mysqlPassword := viper.GetString("DATABASE_PASSWORD")
	mysqlHost := viper.GetString("DATABASE_HOST")
	mysqlPort := viper.GetString("DATABASE_PORT")
	mysqlDbname := viper.GetString("DATABASE_DBNAME")

	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUsername,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDbname,
	)

	return mysqlDsn
}

func MySQLInit() {
	mysqlDns := MySQLCreds()
	db, err := gorm.Open(mysql.Open(mysqlDns), &gorm.Config{})
	if err != nil {
		panic(constants.DB_PANIC)
	}
	db.AutoMigrate(&models.Job{})
}
