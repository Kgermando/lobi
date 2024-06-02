package admin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/models"
)

func NewsLetterPage(c *fiber.Ctx) error {
	db := database.DB

	var newsletters []models.NewsLetter
	db.Find(&newsletters)

	route := c.Path()
	return c.Render("pages/admin/newsletter-list", fiber.Map{
		"nav":            route,
		"Title":          "Admin newsletter",
		"newsletterList": newsletters,
	})
}

func NewsLetterAdd(c *fiber.Ctx) error {
	db := database.DB

	email := c.FormValue("email")

	newsletter := models.NewsLetter{
		Email: email,
	}

	if err := c.BodyParser(&newsletter); err != nil {
		fmt.Println("Une erreur s'est produite de votre contenu", err.Error())
		c.Redirect("/web")
	}

	if err := db.Model(&models.NewsLetter{}).Create(&newsletter).Error; err != nil {
		fmt.Println("Une erreur s'est produite", err.Error())
		c.Redirect("/web")
	}

	return c.Redirect("/web")
}

func NewsLetterDelete(c *fiber.Ctx) error {
	id := c.FormValue("id")
	db := database.DB

	var newsletter models.NewsLetter
	db.Find(&newsletter, id)
	if newsletter.Email == "" {
		return c.Redirect("/web/admin/newsletter-list")
	}
	db.Delete(&newsletter)
	return c.Redirect("/web/admin/newsletter-list")
}
