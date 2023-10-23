package db

import (
	"github.com/McFalljb/engineering-assessment/golang/dao"
	"github.com/McFalljb/engineering-assessment/golang/models"
)

func Migrate() error {
	if err := dao.DB.AutoMigrate(&models.FoodTruck{}); err != nil {
		return err
	}
	return nil
}
