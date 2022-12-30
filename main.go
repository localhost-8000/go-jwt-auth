package main

import (
	"jwt-auth/database"
	"jwt-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {
	database.ConnectDB()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

    routes.Setup(app)

    app.Listen(":3000")
}


