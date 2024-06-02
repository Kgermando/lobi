package moncompte

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/helpers"
)

func ProfilPage(c *fiber.Ctx) error {
	user, isAuth := helpers.UserJWT(c)

	

	route := c.Path()
	return c.Render("pages/mycompte/profil", fiber.Map{
		"nav":         route,
		"Title":       "Mon profil",
		"user":        user,
		"isAuth":      isAuth,
	})
}
