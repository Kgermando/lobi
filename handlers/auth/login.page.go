package auth

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kgermando/lobi/config"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/helpers"
	"github.com/kgermando/lobi/models"
	"github.com/kgermando/lobi/utils"
)

func LoginPage(c *fiber.Ctx) error {
	user, isAuth := helpers.UserJWT(c)


	route := c.Path()
	return c.Render("pages/auth/login", fiber.Map{
		"nav":         route,
		"Title":       "Login page",
		"user":        user,
		"isAuth":      isAuth,
	})
}

func Login(c *fiber.Ctx) error {
	db := database.DB

	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	input := LoginInput{
		Email:    email,
		Password: password,
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Render("pages/notify/infos", fiber.Map{
			"Statut":     "Erreur",
			"Title":      "Champs non valides 😰",
			"route":      "/web/auth/login",
			"navigation": "Reprendre",
		})
	}

	var user models.User

	db.Where("email = ?", input.Email).First(&user) //Check the email is present in the DB

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		c.Status(fiber.StatusNotFound)
		return c.Render("pages/notify/infos", fiber.Map{
			"Statut":     "Erreur",
			"Title":      "Utilisateur non trouvé 😰",
			"route":      "/web/auth/login",
			"navigation": "Reprendre",
		})
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.Render("pages/notify/infos", fiber.Map{
			"Statut":     "Erreur",
			"Title":      "Mot de passe incorrect 😰",
			"route":      "/web/auth/login",
			"navigation": "Reprendre",
		})
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(expireTime),
	})
	token, err := t.SignedString([]byte(config.ConfigEnv("SECRET")))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Render("pages/notify/infos", fiber.Map{
			"Statut":     "Erreur",
			"Title":      "Vous ne pouvez pas vous conneter 😰",
			"route":      "/web/auth/login",
			"navigation": "Reprendre",
		})
	}

	expires := time.Now().Add(time.Hour * 72)
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expires,
		HTTPOnly: true,
		Secure:   true,
	} // Creates the cookie to be passed.

	c.Cookie(&cookie)

	return c.Redirect("/web")
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // 1 day ,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.Redirect("/web/auth/login")
}
