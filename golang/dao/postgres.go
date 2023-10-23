package dao

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() error {
	dsn := viper.GetString("POSTGRES_CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Debug()
	DB = db
	return nil
}
