package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"web1/internal/core"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]*core.Product, error)
	GetById(ctx context.Context, id string) (*core.Product, error)
	AddProduct(ctx context.Context, product *core.Product) (*core.Product, error)
}

type ProductHandler struct {
	productService ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{productService: service}
}

func (handler *ProductHandler) InitRoutes(app *fiber.App) {
	app.Get("/products", handler.GetAll)
	app.Get("/products/:productId", handler.GetById)
	app.Post("/products", handler.AddProduct)
}

func (handler *ProductHandler) GetAll(ctx *fiber.Ctx) error {
	products, err := handler.productService.GetAll(ctx.UserContext())

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	productMap := map[string][]*core.Product{
		"products": products,
	}

	return ctx.Status(http.StatusOK).JSON(productMap)
}

func (handler *ProductHandler) GetById(ctx *fiber.Ctx) error {
	product, err := handler.productService.GetById(ctx.UserContext(), ctx.Params("productId"))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	productMap := map[string]*core.Product{
		"product": product,
	}

	return ctx.Status(http.StatusOK).JSON(productMap)
}

func (handler *ProductHandler) AddProduct(ctx *fiber.Ctx) error {
	product := &core.Product{}
	err := ctx.BodyParser(product)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	savedProduct, err := handler.productService.AddProduct(ctx.UserContext(), product)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	productMap := map[string]*core.Product{
		"product": savedProduct,
	}

	return ctx.Status(http.StatusCreated).JSON(productMap)
}
