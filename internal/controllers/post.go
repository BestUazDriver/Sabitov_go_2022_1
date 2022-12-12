package controllers

import (
	"github.com/gofiber/fiber/v2"
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
}

func (handler *PostHandler) GetPosts(ctx *fiber.Ctx) error {
	return ctx.Status(200).
		JSON(
			map[string][]*core.Post{"posts": handler.postService.GetPosts()})
}
