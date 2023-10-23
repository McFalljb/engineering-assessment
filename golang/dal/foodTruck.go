package dal

import (
	"github.com/McFalljb/engineering-assessment/golang/dao"
	"github.com/McFalljb/engineering-assessment/golang/models"
)

type FoodTruck struct {
	*models.FoodTruck
}

func (ft *FoodTruck) Create() error {
	if err := dao.DB.Create(ft).Error; err != nil {
		return err
	}
	return nil
}

func (ft *FoodTruck) Update(updateData map[string]interface{}) error {
	if err := dao.DB.Model(&FoodTruck{}).Where("id = ?", updateData["id"]).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func (ft *FoodTruck) Delete() error {
	if err := dao.DB.Delete(ft).Error; err != nil {
		return err
	}
	return nil
}

func (ft *FoodTruck) GetByID() error {
	if err := dao.DB.Where("id = ?", ft.ID).First(ft).Error; err != nil {
		return err
	}
	return nil
}

func (ft *FoodTruck) SearchFoodItem() ([]models.FoodTruck, error) {
	var foodTrucks []models.FoodTruck
	if err := dao.DB.Where("fooditems ILIKE ?", "%"+ft.FoodItems+"%").Find(&foodTrucks).Error; err != nil {
		return nil, err
	}

	return foodTrucks, nil
}

func (ft *FoodTruck) GetRandom() error {

	if err := dao.DB.Where(ft).Order("RANDOM()").First(ft).Error; err != nil {
		return err
	}
	return nil
}
