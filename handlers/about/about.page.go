package about

import ( 

	"github.com/gofiber/fiber/v2" 
	"github.com/kgermando/lobi/helpers" 
)

func AboutPage(c *fiber.Ctx) error { 
	user, isAuth := helpers.UserJWT(c)
 

	route := c.Path()
	return c.Render("pages/about/about", fiber.Map{
		"nav":         route,
		"Title":       "A propos de nous",
		"user":        user,
		"isAuth":      isAuth, 
	})
}
