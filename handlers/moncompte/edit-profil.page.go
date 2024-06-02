package moncompte

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/helpers"
	"github.com/kgermando/lobi/models"
)

func EditProfilPage(c *fiber.Ctx) error {
	user, isAuth := helpers.UserJWT(c)



	route := c.Path()
	return c.Render("pages/mycompte/edit-profil", fiber.Map{
		"nav":         route,
		"Title":       "Modification du profil",
		"user":        user,
		"isAuth":      isAuth,
	})
}

func EditProfilForm(c *fiber.Ctx) error {
	db := database.DB
	user, isAuth := helpers.UserJWT(c)
	fmt.Println(isAuth)

	fullname := c.FormValue("fullname")
	address := c.FormValue("address")
	telephone := c.FormValue("telephone")

	input := models.User{
		Fullname:  fullname,
		Address:   address,
		Telephone: telephone,
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Render("pages/notify/infos", fiber.Map{
			"Statut":     "Erreur",
			"Title":      "Champs non valides",
			"route":      "/web/compte/edit-profil",
			"navigation": "Reprendre",
		})
	}

	var u models.User
	db.First(&u, user.ID)
	u.Fullname = input.Fullname
	u.Address = input.Address
	u.Telephone = input.Telephone

	db.Updates(&u)

	return c.Redirect("/web/compte/profil")
}
