package location

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

//SetUpRoutes build handler routes
func SetUpRoutes(app *fiber.App) {
	app.Get("/via_cep_wrapper/:zipCode", getCep)
}

//getCep GET endpoint to search a given zipCode
func getCep(ctx *fiber.Ctx) error {
	zipCode := ctx.Params("zipCode")
	result, err := SearchLocation(zipCode)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(ErrorMessage{err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
