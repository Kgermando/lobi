package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kgermando/lobi/handlers/about"
	"github.com/kgermando/lobi/handlers/admin"
	"github.com/kgermando/lobi/handlers/auth"
	"github.com/kgermando/lobi/handlers/contact"
	"github.com/kgermando/lobi/handlers/home"
	"github.com/kgermando/lobi/handlers/moncompte"
)

func SetupRoutes(app *fiber.App) {

	redirectPath := "/web"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(redirectPath, fiber.StatusMovedPermanently)
	})

	homeRoute := app.Group("/web", logger.New())
	homeRoute.Get("/", home.HomePage)
	homeRoute.Get("/about", about.AboutPage)
	homeRoute.Get("/contact", contact.ContactPage)
 

	// Authentification routes
	authRoute := homeRoute.Group("/auth")
	authRoute.Get("/login", auth.LoginPage)
	authRoute.Get("/register", auth.RegisterPage)
	authRoute.Get("/forgot", auth.ForgotPage)
	authRoute.Get("/reset", auth.ResetPage)

	authRoute.Post("/new-user", auth.NewUser)
	authRoute.Post("/login-form", auth.Login)
	authRoute.Get("/logout", auth.Logout)
	authRoute.Post("/forgot-password", auth.Forgot)
	authRoute.Post("/reset/:token", auth.ResetPassword)

	// My compte
	compteRoute := homeRoute.Group("/compte")
	compteRoute.Get("/profil", moncompte.ProfilPage) 
	compteRoute.Get("/change-password", moncompte.ChangePasswordPage)
	compteRoute.Post("/change-password", moncompte.ChangePasswordPage)
	compteRoute.Get("/edit-profil", moncompte.EditProfilPage)
	compteRoute.Post("/edit-profil-form", moncompte.EditProfilForm)
	

	// Admin 
	adminRoute := homeRoute.Group("/admin")
	adminRoute.Get("/", admin.AdminPage)

	// Users
	adminRoute.Get("/user-list", admin.UserListPage)
	adminRoute.Get("/user-view/:id", admin.UserViewPage)
	adminRoute.Post("/user-update", admin.UserUpdatePage)
	adminRoute.Post("/user-delete", admin.UserDelete)


	// NewsLetter
	adminRoute.Get("/newsletter-list", admin.NewsLetterPage)
	adminRoute.Post("/newsletter", admin.NewsLetterAdd)
	adminRoute.Post("/newsletter-delete", admin.NewsLetterDelete)

	// Contact admin
	adminRoute.Get("/contact-list", admin.ContactListPage)
	adminRoute.Get("/contact-view/:id", admin.ContactView)
	adminRoute.Post("/contact-update", admin.ContactUpdate)
}
