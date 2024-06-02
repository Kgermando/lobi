package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
	// "github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/router"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3050"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	// database.ConnectDB()

	engine := django.New("./templates", ".html")

	engine.Reload(true) // Optional. Default: false
	engine.PreRenderCheck()

	app := fiber.New(fiber.Config{
		Views:        engine,
		AppName:      "LOBI",
		ServerHeader: "lobi-es.com",
	})

	// Initialize default config
	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Static("/uploads", "./public")

	app.Static("/", "./public")
	app.Static("/web", "./public")
	app.Static("/web/auth", "./public")
	app.Static("/web/boutique", "./public")
	app.Static("/web/admin", "./public")
	app.Static("/web/admin/categorie-view", "./public")
	app.Static("/web/admin/user-view", "./public")
	app.Static("/web/admin/product-view", "./public")
	app.Static("/web/admin/contact-view", "./public")
	app.Static("/web/compte", "./public")

	log.Fatal(app.Listen(getPort()))

}
