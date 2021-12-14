package location

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type Handler struct {
	locationService Service
}

func NewHandler(locationService Service) *Handler {
	return &Handler{locationService: locationService}
}

//SetUpRoutes build handler routes
func (handler *Handler) SetUpRoutes(app *fiber.App) {
	app.Get("/via_cep_wrapper/:zipCode", handler.getCep)
}

//getCep GET endpoint to search a given zipCode
func (handler *Handler) getCep(ctx *fiber.Ctx) error {
	zipCode := ctx.Params("zipCode")
	result, err := handler.locationService.SearchLocation(zipCode)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(ErrorMessage{err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
