package auth

import ( 
	"net/smtp"
	"os" 
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/config"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/helpers"
	"github.com/kgermando/lobi/models"
	"github.com/kgermando/lobi/utils"
)

func ForgotPage(c *fiber.Ctx) error { 
	user, isAuth := helpers.UserJWT(c)

	 
	route := c.Path()
	return c.Render("pages/auth/forgot", fiber.Map{
		"nav":         route,
		"Title":       "Mot de passe oublié",
		"user":        user,
		"isAuth":      isAuth, 
	})
}

func Forgot(c *fiber.Ctx) error {
	db := database.DB

	email := c.FormValue("email")

	u := models.PasswordReset{
		Email: email,
	}

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	token := utils.GenerateRandomString(12)

	pr := &models.PasswordReset{
		Email: u.Email,
		Token: token,
	}

	// search for the email in the database, if the user exist
	um := &models.User{}
	db.Where("email = ?", u.Email).First(um)
	if um.ID == 0 {
		c.Status(400)
		return c.Render("pages/notify/infos", fiber.Map{
			"Statut":     "Erreur",
			"Title":      "Utilisateur non trouvé 😰",
			"route":      "/web/auth/forgot",
			"navigation": "Reprendre",
		})
	}

	// token expiration time is 3hr
	pr.ExpirationTime = time.Now().Add(time.Hour * time.Duration(3))
	pr.CreatedAt = time.Now()

	db.Create(pr)

	from := config.ConfigEnv("SMTP_MAIL") // os.Getenv("EMAIL_FROM")

	to := []string{
		u.Email,
	}

	auth := smtp.PlainAuth("", config.ConfigEnv("SMTP_USERNAME"), config.ConfigEnv("SMTP_PASSWORD"), config.ConfigEnv("SMTP_HOST"))

	url := os.Getenv("RESET_URL") + token

	msg := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")

	err := smtp.SendMail(config.ConfigEnv("SMTP_HOST")+":"+config.ConfigEnv("SMTP_PORT"), auth, from, to, msg)
	if err != nil {
		c.Status(400)
		return c.Render("pages/notify/infos", fiber.Map{
			"Statut":     "Erreur",
			"Title":      "Votre adresse mail n'a pas été envoyé 😰",
			"route":      "/web/auth/forgot",
			"navigation": "Reprendre",
		})
	}

	return c.Redirect("/web/auth/login")

}
