package routes

import (
	"api-golang/controllers"
	"api-golang/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// RegisterProducts : Register all routes related to products CRUD
func RegisterProducts(app *fiber.App) {
	app.Post("/products", func(ctx *fiber.Ctx) error {

		var product models.Product
		json.Unmarshal([]byte(ctx.Body()), &product)

		result, err := controllers.CreateProduct(product)

		if err != nil {
			ctx.Status(500).JSON(err)
			return err
		}

		ctx.JSON(&result)
		return nil
	})

	app.Get("/products", func(ctx *fiber.Ctx) error {
		products, err := controllers.GetProducts()

		if err != nil {
			ctx.Status(500).JSON(err)
			return err
		}

		ctx.Send(products)
		return nil
	})

	app.Delete("/products/:id", func(ctx *fiber.Ctx) error {
		result, err := controllers.DeleteProductByID(ctx.Params("id"))

		if err != nil {
			ctx.JSON(err)
			return err
		}

		ctx.JSON(result)

		return nil
	})
}
