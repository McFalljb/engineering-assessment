package controllers

import (
	"log"
	"net/http"

	"github.com/McFalljb/engineering-assessment/golang/dal"
	"github.com/gin-gonic/gin"
)

func CreateFoodTruck(c *gin.Context) {
	var foodTruck dal.FoodTruck
	if err := c.ShouldBind(&foodTruck); err != nil {
		log.Printf("error binding foodTruck: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := foodTruck.Create(); err != nil {
		log.Printf("error creating foodTruck: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, foodTruck)
}

func UpdateFoodTruck(c *gin.Context) {
	// Bind the request to a map to get all fields
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		log.Printf("error binding request data: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := updateData["id"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be present in the request"})
		return
	}

	var foodTruck dal.FoodTruck
	if err := foodTruck.Update(updateData); err != nil {
		log.Printf("error updating foodTruck: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foodTruck)
}

func DeleteFoodTruck(c *gin.Context) {
	var foodTruck dal.FoodTruck
	if err := c.ShouldBind(&foodTruck); err != nil {
		log.Printf("error binding foodTruck: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := foodTruck.Delete(); err != nil {
		log.Printf("error deleting foodTruck: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, foodTruck)
}

func GetFoodTruckByID(c *gin.Context) {
	var foodTruck dal.FoodTruck
	if err := c.ShouldBind(&foodTruck); err != nil {
		log.Printf("error binding foodTruck: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := foodTruck.GetByID(); err != nil {
		log.Printf("error getting foodTruck: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, foodTruck)
}

func SearchFoodTrucks(c *gin.Context) {
	var foodTruck dal.FoodTruck
	if err := c.ShouldBind(&foodTruck); err != nil {
		log.Printf("error binding foodTruck: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foodTrucks, err := foodTruck.SearchFoodItem()
	if err != nil {
		log.Printf("error getting foodTrucks: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, foodTrucks)
}

func GetRandomFoodTruck(c *gin.Context) {
	foodTruck := dal.FoodTruck{}

	if err := foodTruck.GetRandom(); err != nil {
		log.Printf("error getting foodTruck: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foodTruck)
}
