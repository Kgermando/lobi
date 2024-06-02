package moncompte

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/helpers"
	"github.com/kgermando/lobi/models"
	"github.com/kgermando/lobi/utils"
	"golang.org/x/crypto/bcrypt"
)

func ChangePasswordPage(c *fiber.Ctx) error {
	user, isAuth := helpers.UserJWT(c)

	

	route := c.Path()
	return c.Render("pages/mycompte/change-password", fiber.Map{
		"nav":         route,
		"Title":       "Changez votre mot de passe.",
		"user":        user,
		"isAuth":      isAuth,
	})
}

func ChangePassword(c *fiber.Ctx) error {
	db := database.DB
	user, isAuth := helpers.UserJWT(c)
	fmt.Println(isAuth)

	type LoginInput struct {
		OldPassword     string `json:"password"`
		NewPassword     string `json:"newpassword"`
		ConfirmPassword string `json:"confirmpassword"`
	}

	oldpassword := c.FormValue("oldpassword")
	newpassword := c.FormValue("newpassword")
	confirmpassword := c.FormValue("confirmpassword")

	input := LoginInput{
		OldPassword:     oldpassword,
		NewPassword:     newpassword,
		ConfirmPassword: confirmpassword,
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"errors":  err.Error(),
		})
	}

	if input.NewPassword != input.ConfirmPassword {
		c.Status(400)
		fmt.Println("message: password does not match")
		return c.Redirect("/web/compte/profil")
	}

	if !utils.CheckPasswordHash(input.OldPassword, user.Password) {
		return c.Redirect("/web/compte/profil")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 14)
	db.Model(&models.User{}).Where("email = ?", user.Email).Update("password", password)

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // 1 day ,
		HTTPOnly: true,
		// SameSite: "lax",
	}
	c.Cookie(&cookie)

	return c.Redirect("/web/auth/login")
}
