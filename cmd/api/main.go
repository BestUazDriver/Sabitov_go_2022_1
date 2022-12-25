package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"web1/internal/controllers"
	"web1/internal/repositories/mongo"
	"web1/internal/services"

	"time"
	"web1/internal/config"
)

// @title Fiber Swagger Example Api
// version 2.0
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	viper.AddConfigPath("config")
	viper.SetConfigName("conf")
	errConfig := viper.ReadInConfig()
	if errConfig != nil {
		log.Panic(errConfig)
	}
	port := viper.GetString("http.port")

	app := fiber.New()
	config.SwaggerSetUp(app)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoDB, err := config.MongoSetUp(ctx, cancel)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err != nil {
		log.Panic(err)
	}

	productRepository := mongo.NewProductRepository(mongoDB.Collection("products"))
	productService := services.NewProductService(productRepository)
	productHandler := controllers.NewProductHandler(productService)
	productHandler.InitRoutes(app)

	err = app.Listen(":" + port)
}
