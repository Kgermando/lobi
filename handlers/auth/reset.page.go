package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/helpers"
	"github.com/kgermando/lobi/models"
	"golang.org/x/crypto/bcrypt"
)

func ResetPage(c *fiber.Ctx) error {
	user, isAuth := helpers.UserJWT(c)

	
	route := c.Path()
	return c.Render("pages/auth/reset", fiber.Map{
		"nav":         route,
		"Title":       "Reinitialisez votre mot de passe",
		"user":        user,
		"isAuth":      isAuth,
	})
}

func ResetPassword(c *fiber.Ctx) error {
	db := database.DB

	token := c.Params("token")

	passwordForm := c.FormValue("password")
	passwordConfirm := c.FormValue("password_confirm")

	rp := &models.PasswordReset{}

	if err := db.Where("token = ?", token).Last(rp); err.Error != nil {
		c.Status(400)
		fmt.Println("message: invalid token")
		return c.Redirect("/web/auth/login")
	}

	if rp.Id == 0 {
		c.Status(400)
		fmt.Println("message: invalid token")
		return c.Redirect("/web/auth/login")
	}

	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if now.After(rp.ExpirationTime) {
		c.Status(400)
		fmt.Println("message: token has expired")
		return c.Redirect("/web/auth/login")
	}

	// r := new(models.Reset)
	r := models.Reset{
		Password:        passwordForm,
		PasswordConfirm: passwordConfirm,
	}

	if r.Password != r.PasswordConfirm {
		c.Status(400)
		fmt.Println("message: password does not match")
		return c.Redirect("/web/auth/login")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(r.Password), 14)
	db.Model(&models.User{}).Where("email = ?", rp.Email).Update("password", password)

	fmt.Println("message: success")
	return c.Redirect("/web/auth/login")
}
