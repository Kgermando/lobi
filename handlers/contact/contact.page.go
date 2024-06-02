package contact

import ( 

	"github.com/gofiber/fiber/v2" 
	"github.com/kgermando/lobi/helpers" 
)

func ContactPage(c *fiber.Ctx) error { 
	user, isAuth := helpers.UserJWT(c)
 

	route := c.Path()
	return c.Render("pages/contact/contact", fiber.Map{
		"nav":         route,
		"Title":       "Contactez-nous",
		"user":        user,
		"isAuth":      isAuth, 
	})
}
