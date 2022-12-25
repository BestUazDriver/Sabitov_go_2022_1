package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
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
	viper.SetConfigFile(".env")
	errConfig := viper.ReadInConfig()
	if errConfig != nil {
		log.Panic(errConfig)
	}
	port := viper.Get("PORT").(string)
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	abs, errPath := filepath.Abs("internal\\data")

	if errPath != nil {
		log.Panic(errPath)
	}

	userRepository := memory.NewUserRepository(abs + "\\users.txt")
	userService := services.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(userService)

	postRepository := memory.NewPostRepository(abs+"\\posts.txt", userRepository)
	postService := services.NewPostService(postRepository)
	postHandler := controllers.NewPostHandler(postService)

	//post := postService.AddPost(&core.Post{
	//	Likes:   234,
	//	Owner:   nil,
	//	Content: "Why not working",
	//})

	postHandler.InitPostRoutes(app)
	userHandler.InitRoutes(app)

	err := app.Listen(":" + port)

	if err != nil {
		log.Panic(err)
	}

}
