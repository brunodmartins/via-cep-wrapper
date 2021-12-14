package main

import (
	"BrunoDM2943/via-cep-wrapper/internal/modules/location"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/via_cep_wrapper/:cep", location.GetCEP)
	log.Fatal(app.Listen(":8080"))
}
