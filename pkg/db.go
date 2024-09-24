package pkg

import (
	"github.com/spf13/viper"
	"github.com/video-validator/pkg/database"
)

func DbInit() {
	dbType := viper.GetString("DATABASE_DBTYPE")
	if dbType == "sqlite" { // I have used SQLite3 for local testing
		database.SqliteInit()
	}

	if dbType == "postgres" {
		database.PostgresInit()
	}

	if dbType == "mysql" {
		database.MySQLInit()
	}

	if dbType == "mongodb" {
		database.MongoDBInit()
	}
}
