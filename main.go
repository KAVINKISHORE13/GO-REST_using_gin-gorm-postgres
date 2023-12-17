package main

import (
	"github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/config"
	"github.com/kavinkishore13/go-rest_using_gin-gorm-postgres/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	_ "github.com/go-redis/redis/v8"
	_"gorm.io/gorm"
)

func main() {
	
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}


	redisClient := config.SetupRedis()

	recentDelClient := config.SetupRecentDeletedRedis()

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, db, redisClient, recentDelClient)
	router.Run(":8080")
}