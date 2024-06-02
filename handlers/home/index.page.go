package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/helpers"
)

func HomePage(c *fiber.Ctx) error {
	// db := database.DB
	user, isAuth := helpers.UserJWT(c)

	// var productList []models.Product
	// db.Order("random()").Limit(16).Find(&productList)

	// var categories []models.Category
	// db.Order("random()").Find(&categories)

	 

	route := c.Path()
	return c.Render("pages/home/index", fiber.Map{
		"nav":           route,
		"Title":         "Acceuil",
		"user":          user,
		"isAuth":        isAuth, 
	})
}
