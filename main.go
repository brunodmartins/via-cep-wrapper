package main

import (
	"BrunoDM2943/via-cep-wrapper/internal/gateway/viacep"
	"BrunoDM2943/via-cep-wrapper/internal/modules/location"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := viacep.NewGateway(http.DefaultClient, "https://viacep.com.br")
	service := location.NewLocationService(api)
	handler := location.NewHandler(service)
	handler.SetUpRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
