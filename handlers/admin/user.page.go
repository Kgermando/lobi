package admin

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/lobi/database"
	"github.com/kgermando/lobi/helpers"
	"github.com/kgermando/lobi/models"
)

func UserListPage(c *fiber.Ctx) error {
	db := database.DB
	user, isAuth := helpers.UserJWT(c)

	var users []models.User
	db.Find(&users)

	route := c.Path()
	return c.Render("pages/admin/user-list", fiber.Map{
		"nav":      route,
		"Title":    "Admin",
		"user":     user,
		"isAuth":   isAuth,
		"userList": users,
	})
}

func UserViewPage(c *fiber.Ctx) error {
	db := database.DB
	user, isAuth := helpers.UserJWT(c)

	id := c.Params("id")

	var u models.User
	db.Where("id = ?", id).First(&u)

	route := c.Path()
	return c.Render("pages/admin/user-view", fiber.Map{
		"nav":    route,
		"Title":  "Admin",
		"user":   user,
		"isAuth": isAuth,
		"u":      u,
	})
}

func UserUpdatePage(c *fiber.Ctx) error {
	db := database.DB

	id := c.FormValue("id")
	email := c.FormValue("email")
	fullname := c.FormValue("fullname")
	telephone := c.FormValue("telephone")
	role := c.FormValue("role")
	is_active := c.FormValue("is_active")
	address := c.FormValue("address")

	isActive, err := strconv.ParseBool(is_active)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
		return err
	}

	input := models.User{
		Email:     email,
		Fullname:  fullname,
		Telephone: telephone,
		Role:      role,
		IsActive:  isActive,
		Address:   address,
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Redirect("/web/admin/user-list")
	}

	var user models.User
	db.Where("id = ?", id).First(&user)

	db.First(&user, id)
	user.Email = input.Email
	user.Fullname = input.Fullname
	user.Telephone = input.Telephone
	user.Role = input.Role
	user.IsActive = input.IsActive
	user.Address = input.Address
	db.Updates(&user)

	return c.Redirect("/web/admin/user-list")
}

func UserDelete(c *fiber.Ctx) error {
	id := c.FormValue("id")
	db := database.DB

	var user models.User
	db.Find(&user, id)
	if user.Email == "" {
		return c.Redirect("/web/admin/user-list")

	}
	db.Delete(&user)
	return c.Redirect("/web/admin/user-list")
}
