package main

import (
	"log"
	"os"

	"vote-app-api/infrastructure/api/router"
	"vote-app-api/infrastructure/datastore"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	AppEnv := os.Getenv("APP_ENV")

	log.Printf("Environment is %v\n", AppEnv)

	conn := datastore.NewMysqlDB()
	defer conn.Close()

	c := gin.Default()
	if AppEnv == "stg" || AppEnv == "prod" {
		c.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"https://production.com", "https://api.production.com", "https://stg.production.com"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
		}))
	} else {
		c.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*", "http://localhost:3000"},
			AllowMethods: []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
			AllowHeaders: []string{"*"},
		}))
	}

	router.NewApiV1Router(c, conn)
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("db cannot connect %+v\n", err)
		}
	}()

	c.Run()
}
