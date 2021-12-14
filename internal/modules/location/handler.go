package location

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetCEP(ctx *fiber.Ctx) error {
	zipCode := ctx.Params("cep")
	if zipCode == "" {
		return ctx.Status(http.StatusBadRequest).JSON(ErrorMessage{"Cep esta vazio"})
	}
	result, err := SearchLocation(zipCode)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(ErrorMessage{err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
