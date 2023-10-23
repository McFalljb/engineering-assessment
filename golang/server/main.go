package main

import (
	"log"
	"net/http"
	"os"

	"github.com/McFalljb/engineering-assessment/golang/dao"
	"github.com/McFalljb/engineering-assessment/golang/db"
	"github.com/McFalljb/engineering-assessment/golang/server/controllers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	gin.DisableConsoleColor()

	// Enable automatic environment variable binding
	viper.AutomaticEnv()

	if err := dao.ConnectPostgres(); err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}
	log.Println("Connected to database")
	if err := db.Migrate(); err != nil {
		log.Fatalf("error migrating database: %s", err)
	}
	log.Println("Migrated database")

	r := gin.Default()
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// So any panics you want to stop the server from running should be ran before this.
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// ROUTES
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	r.POST("/killSwitch", killSwitch)
	r.POST("/api/v1/foodTruck/create", controllers.CreateFoodTruck)
	r.PUT("/api/v1/foodTruck/update", controllers.UpdateFoodTruck)
	r.POST("/api/v1/foodTruck/delete", controllers.DeleteFoodTruck)
	r.POST("/api/v1/foodTruck/getByID", controllers.GetFoodTruckByID)
	r.POST("/api/v1/foodTruck/searchFoodItem", controllers.SearchFoodTrucks)
	r.POST("/api/v1/foodTruck/getRandom", controllers.GetRandomFoodTruck)

	server := &http.Server{
		Addr:    ":" + viper.GetString("BACKEND_PORT"),
		Handler: r,
	}

	server.ListenAndServe()
}

func killSwitch(c *gin.Context) {
	log.Println("Immediate termination requested via /killSwitch endpoint!")
	c.JSON(http.StatusOK, gin.H{"message": "Terminating immediately..."})
	os.Exit(0)
}
