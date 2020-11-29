package routes

import (
	"api-golang/controllers"
	"api-golang/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

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
}
