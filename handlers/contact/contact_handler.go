package contact

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/models"
)

func ContactMe(c *fiber.Ctx) error {
	db := database.DB

	input := new(models.Contact)
	if err := c.BodyParser(input); err != nil {
		fmt.Println("Une erreur s'est produite de votre contenu", err.Error())
		c.Redirect("/web/contact")
	}

	if err := db.Create(&input).Error; err != nil {
		fmt.Println("Une erreur s'est produite", err.Error())
		c.Redirect("/web/contact")
	}

	return c.Redirect("/web/contact")
}
