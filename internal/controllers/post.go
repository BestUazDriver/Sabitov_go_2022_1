package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"web1/internal/core"
	"web1/internal/services"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{postService: service}
}

func (postHandler *PostHandler) InitPostRoutes(app *fiber.App) {
	app.Get("/posts", postHandler.GetPosts)
	app.Post("/addPost", postHandler.AddPost)
}

func (handler *PostHandler) GetPosts(ctx *fiber.Ctx) error {
	return ctx.Status(200).
		JSON(
			map[string][]*core.Post{"posts": handler.postService.GetPosts()})
}

func (handler *PostHandler) AddPost(ctx *fiber.Ctx) error {
	post := &core.Post{}
	err := ctx.BodyParser(post)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
			"Error": err.Error(),
		})
	}
	return ctx.Status(200).
		JSON(
			map[string]*core.Post{"post": handler.postService.AddPost(post)})
}
