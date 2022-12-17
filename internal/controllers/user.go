package controllers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	core "web1/internal/core"
	"web1/internal/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (userHandler *UserHandler) InitRoutes(app *fiber.App) {
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:id", userHandler.GetById)
}

// GetAll
// @Summary users
// @Tags users
// @Produce JSON
// @Status 200
// @Router /users [GET]
func (userHandler *UserHandler) GetUsers(fiber *fiber.Ctx) error {

	users := userHandler.userService.GetUsers()

	usersMap := map[string][]*core.User{
		"users": users,
	}

	return fiber.Status(200).JSON(
		usersMap,
	)
}

// GetById
// @Summary users/:id
// @Tags users
// @Produce JSON
// @Status 200
// @Router /users/:id [GET]
func (userHandler *UserHandler) GetById(fiber *fiber.Ctx) error {

	id, err := strconv.Atoi(fiber.Params("id"))

	if err != nil {

		errMsg := "Incorrect \"id\" parameter"

		return fiber.Status(404).JSON(
			map[string]string{
				"Error": errMsg,
			},
		)
	}

	user := userHandler.userService.GetById(id)
	usersMap := map[string]*core.User{
		"users": user,
	}
	return fiber.Status(200).JSON(
		usersMap,
	)
}
