package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/helpers"
)

func AdminPage(c *fiber.Ctx) error {
	user, isAuth := helpers.UserJWT(c)

	route := c.Path()
	return c.Render("pages/admin/dashboard", fiber.Map{
		"nav":    route,
		"Title":  "Admin",
		"user":   user,
		"isAuth": isAuth,
	})
}
