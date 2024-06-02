package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/models"
)

func ContactListPage(c *fiber.Ctx) error {
	db := database.DB

	var contacts []models.Contact
	db.Find(&contacts)

	route := c.Path()
	return c.Render("pages/admin/contact-list", fiber.Map{
		"nav":         route,
		"Title":       "Admin contact",
		"contactList": contacts,
	})
}

func ContactView(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var contact models.Contact
	db.First(&contact, id)
	if contact.Email == "" {
		c.Redirect("/web/admin/contact-list")
	}

	route := c.Path()
	return c.Render("pages/admin/contact-view", fiber.Map{
		"nav":     route,
		"Title":   "Admin Contact",
		"contact": contact,
	})
}

func ContactUpdate(c *fiber.Ctx) error {
	db := database.DB

	id := c.FormValue("id")
	is_read := c.FormValue("is_read")

	isRead, err := strconv.ParseBool(is_read)
	if err != nil {
		return err
	}

	input := models.Contact{
		IsRead: isRead,
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Redirect("/web/admin/contact-list")
	}

	var contact models.Contact
	db.First(&contact, id)
	contact.IsRead = input.IsRead
	db.Updates(&contact)

	return c.Redirect("/web/admin/contact-list")
}
