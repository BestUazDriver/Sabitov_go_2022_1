package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"time"
	"web1/internal/config"
	"web1/internal/controllers"
	"web1/internal/repositories/memory"
	"web1/internal/services"
)

// @title Fiber Swagger Example Api
// version 2.0
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
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

	abs, errPath := filepath.Abs("internal\\data")

	if errPath != nil {
		log.Panic(errPath)
	}

	userRepository := memory.NewUserRepository(abs + "\\users.txt")
	userService := services.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(*userService)

	postRepository := memory.NewPostRepository(abs+"\\posts.txt", userRepository)
	postService := services.NewPostService(postRepository)
	postHandler := controllers.NewPostHandler(postService)

	postHandler.InitPostRoutes(app)
	userHandler.InitRoutes(app)

	err = app.Listen(":" + port)

	if err != nil {
		log.Panic(err)
	}

}
