package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/helpers"
	"github.com/kgermando/lobi/models"
	"github.com/kgermando/lobi/utils"
)

func RegisterPage(c *fiber.Ctx) error {
	user, isAuth := helpers.UserJWT(c)


	route := c.Path()
	return c.Render("pages/auth/register", fiber.Map{
		"nav":         route,
		"Title":       "register page",
		"user":        user,
		"isAuth":      isAuth,
	})
}

func NewUser(c *fiber.Ctx) error {
	db := database.DB

	email := c.FormValue("email")
	password := c.FormValue("password")
	fullname := c.FormValue("fullname")
	address := c.FormValue("address")
	telephone := c.FormValue("telephone")

	newUser := models.User{
		Email:         email,
		Password:      "1234",
		Fullname:      fullname,
		Address:       address,
		Telephone:     telephone,
		EmailVerified: false,
		Role:          "User",
		IsActive:      true,
	}

	if err := c.BodyParser(&newUser); err != nil {
		fmt.Println("Une erreur s'est produite de votre contenu", err.Error())
		c.Redirect("/web")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		fmt.Printf("Impossible de hacher le mot de passe %v \n", err.Error())
		c.Redirect("/web/auth/register")
	}

	newUser.Password = hash

	if err := db.Model(&models.User{}).Create(&newUser).Error; err != nil {
		fmt.Println("Une erreur s'est produite", err.Error())
		c.Redirect("/web/auth/register")
	}

	return c.Redirect("/web/auth/login")
}
